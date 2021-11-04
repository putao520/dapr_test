package main

import (
	"context"
	"fmt"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func httpServer(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})
}

func wsServer(app *fiber.App) {
	// 响应升级协议
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
}

func wsTest(app *fiber.App) {
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))
}

/*
func daprTestMysqlBind() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	in := &dapr.InvokeBindingRequest{
		Name:      "my-mysql-store",
		Operation: "query",
		Metadata:  map[string]string{"sql": "SELECT * FROM test"},
	}
	out, err := client.InvokeBinding(ctx, in)
	if err == nil {
		fmt.Println(string(out.Data))
	} else {
		fmt.Println(err)
	}
}


func daprTestRedis() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	ctx := context.Background()

	store := "my-redis-store"
	if err := client.SaveState(ctx, store, "key1", []byte("test")); err != nil {
		panic(err)
	}

	if item, err := client.GetState(ctx, store, "key1"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(item.Value))
	}

	if err := client.DeleteState(ctx, store, "key1"); err != nil {
		panic(err)
	}
}
*/

func daprTestS3() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	ctx := context.Background()

	in := &dapr.InvokeBindingRequest{
		Name:      "my-s3-store",
		Operation: "create",
		Data:      []byte("hello tester"),
		Metadata:  map[string]string{"key": "test-file.txt"},
	}
	out, err := client.InvokeBinding(ctx, in)
	if err == nil {
		fmt.Println(string(out.Data))
	} else {
		fmt.Println(err)
	}
}

func main() {
	app := fiber.New()
	httpServer(app)
	wsServer(app)
	// wsTest(app)

	// daprTestMysqlBind()

	// daprTestRedis()

	daprTestS3()

	app.Listen(":3000")
}

/* init dapr client
client, err := dapr.NewClient()
if err != nil {
	panic(err)
}
defer client.Close()
*/
//TODO: use the client here, see below for examples
