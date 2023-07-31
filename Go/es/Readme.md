- 创建docker网络
`docker network create es`
- 创建`plugin`目录进行挂载
- 启动es
```shell
docker run -d --name xl-es --net es -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
-v /home/xl/docker-conf/es/plugins/:/usr/share/elasticsearch/plugins \
-it elasticsearch:8.8.1
```
- 启动kibana
```shell
docker run -d --name kb --net es -p 5601:5601 kibana:8.8.1
```
- 通过`docker logs containerName 查看token和密码和验证码`
-------------------------------
> kibana版本和elasticsearch版本要一致