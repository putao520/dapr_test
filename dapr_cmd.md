### go-sdk

```
https://github.com/dapr/go-sdk
```

### 启动组件组

```
http模式:
dapr run --app-id myapp --app-port 3500 --dapr-http-port 3500 --components-path ./my-components

grpc模式:
dapr run --app-id myapp --app-port 50001 --app-protocol grpc --dapr-grpc-port 50001 --components-path ./my-components

app在主机，dapr在容器 
http模式:
docker run --net="host" --mount type=bind,source=d:/dapr_test/my-components,target=/components putao520/dapr-gs:2021110708 ./linux_amd64/release/daprd --app-id myapp --app-port 3500 --dapr-http-port 3500 --placement-host-address placement:50006
grpc模式:
docker run --net="host" --mount type=bind,source=d:/dapr_test/my-components,target=/components putao520/dapr-gs:2021110708 ./linux_amd64/release/daprd --app-id myapp --app-port 50001 --app-protocol grpc --dapr-grpc-port 50001 --placement-host-address placement:50006
```

### 存储组件组

```
https://docs.dapr.io/zh-hans/reference/components-reference/supported-bindings/mysql/
```

### 组件源码库

```
https://github.com/dapr/components-contrib/
```

### 编译输出自己的dapr
```
补充复制代码（可以不执行）
Dockerfile最后增加COPY $PWD .

cd dist
docker build -t putao520/dapr-gs:2021110708 -f ../docker/Dockerfile .
docker tag putao520/dapr-gs:2021110708 putao520/dapr-gs:2021110708
docker push putao520/dapr-gs:2021110708
```

### 目前的问题
```
time="2021-11-07T15:56:44.3820888Z" level=warning msg="failed to load components: open : no such file or directory" app_id=myapp instance=docker-desktop scope=dapr.runtime type=log ver=edge
time="2021-11-07T15:56:44.382923Z" level=warning msg="failed to init actors: actors: couldn't connect to placement service: address is empty" app_id=myapp instance=docker-desktop scope=dapr.runtime type=log ver=edge
```