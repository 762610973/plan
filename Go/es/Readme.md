`docker network create es`
```shell
docker run --name xl-es --net es -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-e CLI_JAVA_OPTS="-Xms64m -Xmx128m" \
-v /home/xl/docker-conf/es/plugins/:/usr/share/elasticsearch/plugins \
-it elasticsearch:8.8.1
```
```shell
docker run --name kb --net es -p 5601:5601 kibana:8.8.1
```
> kibana版本和elasticsearch版本要一致