## 隔离管理
- namespace(资源隔离)

  | namespace | 系统调用参数  |          隔离内容          |
  | :-------: | :-----------: | :------------------------: |
  |    UTC    | CLONE_NEWUTS  |         主机和域名         |
  |    IPC    | CLONE_NEWIPC  | 信号量, 消息队列和共享内存 |
  |    PID    | CLONE_NEWPID  |          进程编号          |
  |  Network  | CLONE_NEWNET  |   网络设备, 网络栈, 端口   |
  |   MOunt   |  CLONE_NEWNS  |      挂载点(文件系统)      |
  |   User    | CLONE_NEWUSER |        用户和用户组        |


- cgroup(资源限制)
  - 资源限制: 限制任务使用的资源总额, 超过这个配额时发出提示
  - 优先级分配: 分配CPU时间片数量及磁盘IO带宽大小, 控制任务运行的优先级
  - 资源统计: 统计系统资源使用量, 如CPU使用时长, 内存用量等
  - 任务控制: 对任务挂起, 恢复等操作

## 命令

- docker create --name xx image: 只负责创建容器
- docker pause container: 暂停容器
- docker unpause container: 恢复运行
- docker kill container: 是强制
- docker stop container: 优雅停机
- 容器状态: created(新建), up(运行中), pause(暂停), exited(退出)
- **docker run**立即启动, 默认前台启动
- docker run -d(后台启动)
- docker run -d == docker create + docker start
- docker attach 绑定的是控制台, 可能导致容器停止
- docker exec -it --privileged mynginx /bin/bash
- docker exec -it -u 0:0 mynginx /bin/bash: 0用户以root用户进入
- i: 交互模式, t: 分配一个新的终端
- docker diff: 检查容器里文件结构的更改.A: 添加文件或目录. D: 文件或者目录删除. C: 文件或者目录更改.
- docker run -it busybox:latest 以交互模式启动.