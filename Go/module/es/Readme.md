# Elasticsearch
## 部署
- 创建docker网络
`docker network create es`
- 创建`plugin`目录进行挂载
- 启动es
> 必须带有参数`--privileged=true`
```shell
docker run -it -d --privileged=true --name es --net es -p 9200:9200 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
-v ./es-data/plugins:/usr/share/elasticsearch/plugins \
-v ./es-data/data:/usr/share/elasticsearch/data \
elasticsearch:8.9.2
```
- 启动kibana
```shell
docker run -d --name kb --net es -p 5601:5601 kibana:8.9.2
```
- 通过`docker logs containerName 查看token和密码和验证码`
-------------------------------
> kibana版本和elasticsearch版本要一致

## 功能
- 支持水平扩展, `Restful Api`
- 海量数据的分布式存储
- 近实时搜索
- 海量数据的近实时分析(聚合)
- ![](../../../images/elastic_stack.png)
- logstash: 开源的服务器端数据数据处理管道
  - 实时解析和转换数据
  - 可扩展
  - 可靠性安全性
  - 监控
- kibana: 数据可视化
- beats: Go开发, 轻量的数据采集器
- 日志管理: 日志搜集, 格式化分析, 全文检索, 风险管理

## 介绍

### 文档元数据
- `_index`: 文档所属的索引名
- `_type`: 文档所属的类型名
- `_id`:文档唯一ID
- `_source`:文档的原始JSON数据
- `_version`: 文档的版本信息
- `_score`: 相关性打分

- 索引是文档的容器, 是一类文档的结合
- 索引的Mapping和Settings
  - Mapping定义文档字段的类型
  - Setting定义不同的数据分布
- 一个索引只能创建一个`_doc`

### 分布式系统的可用性与扩展性
- 高可用行
  - 服务可用性: 允许有节点停止服务
  - 数据可用性: 部分节点丢失, 不会丢失数据
- 可扩展性
  - 请求量提升/数据的不断增长(将数据分布到所有节点上)
- 每个节点都有一个节点名, 不同节点承担不同的角色
- 节点
  - master eligible, 只有master节点才能修改集群的状态信息(所有的节点信息, 所有的索引和其相关的Mapping和Setting信息,分片的路由信息), 任意节点都能修改数据会导致数据的不一致性 
  - Data Node: 可以保存数据的节点. 负责保存分片数据, 在数据扩展上起到了至关重要的作用
  - coordinating node: 负责接受client的请求, 将请求分发到合适的节点, 最终把结果汇聚到一起, 每个节点默认起到了coordinating node的职责
  - hot & warm node: 配置高/低, 降低集群部署的成本
  - machine learning node: 负责跑机器学习的job, 用来做异常检测
- 分片
  - 主分片: 用以解决数据视屏扩展的问题. 通过主分片, 可以将数据分布到集群内的所有节点上
  - 副本: 用以解决数据高可用的问题. 分片是主分片的拷贝,副本分片数, 可以动态调整
  - 分片设定
    - 分片数设置过小
      - 导致后续无法增加节点实现水平扩展
      - 单个分片的数据太大, 导致数据重新分配耗时
    - 分片数设置过大
      - 影响搜索结果的相关性打分, 影响统计结果的准确性
      - 单个节点上过多的分片, 会导致资源浪费, 同时也会影响性能
    - `GET _cluster/health`