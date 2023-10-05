# 深入浅出Docker

## 概览
- 对容器发展影响大的技术
  - 内核命名空间(Kernel Namespace)
  - 控制组(Control Group)
  - 联合文件系统(Union File System)
- 实现容器所需的核心Windows内核技术被统称为Windows容器(Windows Container), 用户空间是通过Docker来完成与Windows容器之间交互的
- 运行中的容器共享主机的内核.
- kubernetes提供了一个可插拔的容器运行时接口CRI
- Docker用于创建, 管理和编排容器.
- Docker Engine是用于运行和编排容器的基础设施工具. Docker Engine是运行容器的核心容器运行时.
- Docker开源项目(Moby)
- 开放容器计划: The Open Container Initiative(OCI): 是一个旨在对容器基础架构中的基础组件进行标准化的管理委员会.
  - 镜像规范
  - 运行时规范
- 升级docker
  - 停止docker守护程序
  - 移除旧版本docker
  - 安装新版本docker
  - 配置新版本的docker为开机自启动
  - 确保容器重启成功
- `docker system info` 查看docker存储驱动, 包括client和server相关配置信息
  - Docker的`Device Mapper`存储驱动利用Lvm实现
  - `Device Mapper`: 存储驱动, 使用**direct-lvm**
  - LVM: logical volume manager
  - ```json
    {
      "storage-driver": "devicemapper",
      "storage-opts": [
        "dm.directlvm_device=/dev/xdf",
        "dm.thinp_percent=95", 
        "dm.thinp_metapercent=1",
        "dm.thinp_autoextend_threshold=80", 
        "dm.thinp_autoextend_percent=20",  
        "dm.directlvm_device_force=false"
      ]
    }  
    ```
- docker: docker client和docker daemon(有时也被成为服务端/server或引擎/engine)
  - daemon实现了docker engine的api
  - Linux: client和daemon之间的通信是通过本地IPC/UNIX socket完成的(/var/run/docker.sock)
  - Windows: 通过npipe://// ./pipe/docker_engine的管道来完成的.
## Docker基本命令
- ```shell
  docker image ls
  # it参数会将shell切换到容器终端, 开启容器的交互模式并且用户当前shell连接到容器终端
  docker container run -it
  # 退出容器的同时保持容器运行, 不会杀死容器进程
  Ctrl -PQ
  docker container ls
  docker container exec -it
  docker container stop
  docker container rm
  ```
- `docker exec`: execute a command in a running container
  - usage: docker exec [OPTIONS] CONTAINER COMMAND [ARG...]
  - aliases: docker container exec, docker exec
  - -d(--detach): 分离模式, 在后台运行命令;
  - --detach-keys [string]: Override the key sequence for detaching a container
  - -e(--env) [list]: 设置环境变量
  - --env-file [list]: 从文件中读取环境变量
  - -i(--interactive): 即使没有连接, 也要保持运行.
  - --previleged: 为命令授予权限
  - -t(--tty): 分配一个伪终端
  - -u(--user) [string]: 用户名或uid(format: "<name|uid>[:<group|gid>]")
  - -w(--workdir) [string]: 容器内工作/运行目录



## Docker Engine
> 用来运行和管理容器的核心软件
- ![DockerEngine](../images/DockerEngine1.png)
- 组成: Docker Client, Docker daemon, containerd, runc
- Docker daemon: API和其他特性, 不再包含任何容器运行时的代码, 所有的容器运行代码在一个单独的OCI兼容层(runc)来实现
  - daemon使用一种CRUD风格的api, 通过grpc与containerd进行通信
  - 镜像管理, 镜像构建, REST API, 身份验证, 安全, 核心网络以及编排
- containerd: 容器的生命周期管理, 镜像管理等
  - containerd将Docker镜像转换为OCI bundle, 并让runc基于此创建一个新的容器
  - 然后runc与操作系统内核接口进行通信, 基于所有必要的工具来创建容器
  - 容器进程作为runc的子进程启动, 启动完毕后, runc自动退出
  - **已成为kubernetes中默认的常见的容器运行时**
- ![DockerStart](../images/DockerStart.png)
- 将所有的用于启动, 管理容器的逻辑和代码从daemon中移除, 意味着容器运行时与Docker daemon是解耦的, 有时称之为"无守护进程的容器"
  - 旧模型中, 所有容器运行时的逻辑都在daemon中实现, 启动和停止daemon都会导致宿主机上所有运行中的容器被砍掉.
- shim: 实现无daemon容器
  - 每次创建一个容器时containerd会fork一个runc实例, 创建完毕, 对应的runc进程就会退出(因此即使运行上百个容器, 也无需保持上百个运行中的runc实例)
  - 一旦容器进程的父进程runc退出, 相关联的containerd-shim进程就会成为容器的父进程.
    - 保持所有的STDIN和STDOUT流是开启状态,从而当daemon重启的时候, 容器不会因为管道的关闭而终止
    - 将容器的退出状态反馈给daemon
- runc: 默认的容器运行时, 实质上是一个轻量的, 针对Libcontainer进行了包装的命令行交互工具, 作用是创建容器
  - runc所在那一层称为OCI层

## Docker Image
- 镜像有多个层组成, 每层叠加之后, 从外部看来如同一个独立的对象.
- 镜像内部是一个精简的操作系统, 同时还包含着应用运行所必须的文件和依赖包.
- Docker镜像追求快速和小巧, 构建镜像时会才减掉不必要的东西.
  - 通常Docker镜像中只有一个精简的shell, 甚至没有shell.
  - 镜像中不包含内核: **容器都是共享宿主机的内核**
  - 容器仅包含必要的操作系统(通常只有操作系统文件和文件系统对象)
- 本地镜像仓库: `/var/lib/docker/<storage-driver>`

- `docker image pull repository:tag`
- `docker image ls` 
  - `-a --all`: 显示所有的镜像(默认隐藏中间镜像)
     - `--digests`: 显示摘要
  - `-f --filter [filter]`: 基于给定的条件过滤输出
     `--format [string]`: 使用自定义模板格式化输出: `docker image ls --format "{{.Size}}"`
        table: 表格
        table TEMPLATE: 使用给定的Go模板
        json: 使用json格式
        TEMPLATE: 使用被给定的Go模板
     --no-trunc: 不要截断输出
  - `-q --quiet`: 只显示镜像id
  - filter
    - dangling [bool]: 是否返回悬虚镜像
    - before: 需要镜像名称或者ID作为参数, 返回在之前被创建的全部镜像
    - since: 与before类似, 不过返回的是指定镜像之后创建的全部镜像
    - label: 根据标注(label)的名称或者值, 对镜像进行过滤. docker image ls命令输出中不显示标注内容
    - reference: 其他过滤方式, 比如: `docker image ls --filter=reference="*:latest"`
- `docker image pull [OPTIONS] NAME[:TAG|@DIGEST]`
  - `alias: docker image ls, docker image list, docker images`
  - `-a --all-tags`: 下载仓库中所有打标签的镜像
    - `--disable-content-trust`: 跳过镜像验证
    - `--platform [string]`: 如果服务器支持多平台, 则设置平台
  - `-q --quiet`:    抑制详细输出
- `docker image prune`: 移除全部的悬虚镜像

- 镜像可以有多个标签, latest不一定是最新的镜像
- 没有标签的镜像被称为悬虚对象: `<none>:<none>`, 原因是构建了一个新镜像, 然后为该镜像打了一个已经存在的标签.
- docker search命令允许通过cli的方式搜索docker hub
  - `docker search alpine --filter "is-official=true"`
  - 官方: `docker search alpine --filter=is-official=true`
  - 自动创建的仓库: `docker search alpine --filter=is-automated=true`
  - 增加返回的行数: `--limit=xx`
- 镜像由一些松耦合的只读镜像层组成. docker负责堆叠这些景象层, 并且将他们表示为单个统一的对象.
- `docker image inspect container`
- 所有的docker镜像都起始于一个基础镜像层, 当进行修改或增加新的内容时, 就会在当前镜像层之上, 创建新的镜像层.
- Docker通过存储引擎(新版本通过快照机制)的方式来实现镜像层堆叠, 并保证多镜像层对外展示为统一的文件系统.
- Linux上可用的存储引擎
  - AUFS
  - Overlay2
  - Device Mapper
  - Btrfs
  - ZFS
- 多个镜像之间可以并且确实会共享镜像层, 节省空间+提升性能
- Docker在Linux上支持很多存储引擎(snapshotter), 每个存储引擎都有自己的镜像分层, 镜像层共享以及写时复制(COW)技术的具体实现
- 每个镜像都有单独的签名(digest)
- 镜像层之间是完全独立的, 是实际数据存储的地方
- 镜像的唯一标识是一个加密ID, 即配置对象本身的散列值. 每个镜像层也由一个加密ID区分, 值为镜像层本身内容的散列值
- 镜像内容或其中任意的镜像层发生改动, 都会导致散列值变化, 这就是内容散列(content hash)
- 为了避免压缩带来的散列改变, 每个镜像层还回包含一个分发散列值(压缩版镜像的散列值). 用于校验拉取的镜像是否被篡改过.
- Manifest: 某个镜像标签支持的架构列表
- 多架构镜像
- 删除操作会在当前主机上删除该镜像以及相关的镜像层.
- 如果某个镜像层被多个镜像共享, 那只有当全部依赖该镜像层的镜像都被删除后, 该镜像层才会被删除.