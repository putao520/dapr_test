```
http://192.168.75.80:8888/putao520
putao520
123321aaa
=========================
mysql:
sudo docker run -d --restart=always -p 0.0.0.0:3306:3306 --name root -e MYSQL_ROOT_PASSWORD=123456 -d mysql
=========================
minio:
docker run -p 0.0.0.0:9000:9000 -p 0.0.0.0:9001:9001 --name minio \
-d --restart=always \
-e "MINIO_ACCESS_KEY=minio" \
-e "MINIO_SECRET_KEY=putao520" \
-v /home/data:/data \
-v /home/config:/root/.minio \
minio/minio server /data --console-address "0.0.0.0:9001"
http://192.168.75.80:9001
=========================
redis:
docker run --name redis -d -p 0.0.0.0:6379:6379  redis redis-server --requirepass 123456
=========================
registries:
docker run -d --restart=always -di --name=registry -p 0.0.0.0:5000:5000 registry
http://192.168.75.80:5000/v2/_catalog
```
