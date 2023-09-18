# <span style="color: pink;">Nginx</span>

## 部署

```shell
wget https://nginx.org/download/nginx-1.24.0.tar.gz
tar zxvf nginx-1.24.0.tar.gz
mv nginx-1.24 nginx
cd nginx
yum install -y gcc
yum install -y pcre pcre-devel
yum install -y zlib zlib-devel
./configure --prefix=/usr/local/nginx
make
make install
cd /usr/local/nginx/sbin
./nginx
# 关闭防火墙, 浏览器访问
systemctl stop firewalld.service
# 禁用开机自启
systemctl disable firewalld.service
firewall-cmd --zone=public --add-port=80/tcp --permanent
```
```shell
# 启动
./nginx
# 快速停止
./nginx -s stop
# 优雅关闭,在退出前完成已经接受的连接请求
./nginx -s quit
# 重新加载配置
./nginx -s reload
```

```shell
vim /usr/lib/systemd/system/nginx.service
[Unit]
Description=nginx - web server
After=network.target remote-fs.target nss-lookup.target

[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/usr/local/nginx/sbin/nginx -s reload
ExecStop=/usr/local/nginx/sbin/nginx -s stop
ExecQuit=/usr/local/nginx/sbin/nginx -s quit
PrivateTmp=true
[Install]
WantedBy=multi-user.target
# 重新加载系统服务
systemctl daemon-reload
systemctl start nginx.service
systemctl enable nginx
```

## 配置

### 最小核心配置文件
```shell
user root;
# 工作进程数量, 推荐物理CPU内核数
worker_processes  1;
events {
	# 每个worker的连接数
    worker_connections  1024;
}

http {
	# 导入其他文件. mime.types: 定义解析类型
    include       mime.types;
    # 默认类型
    default_type  application/octet-stream;
    # 数据零拷贝
    sendfile        on;
    # 长连接超时时间
    keepalive_timeout  65;
    server {
      	# 监听的端口号
        listen       80;
        # 主机名,可以配置域名或者主机名
        server_name  localhost;
        # uri
        location / {
          	# 指明根路径, 该location下的所有文件和子目录都相对于该根目录进行访问
            root   html;
            # 匹配到index路由要展示的页面
            index  index.html index.htm;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
          	# 从根目录找50x.html
            root   html;
        }
    }
}
```

#### 域名解析
- 修改hosts文件, 192.168.31.23 kamier.com(本机虚拟)
- 公网域名解析
- 泛域名解析

#### 配置多hosts
> **如果没有匹配到,会从上至下匹配, 匹配是有顺序的**
1. server_name
2. port
- 匹配规则
  - 一个server可以配置多个server_name,使用空格分割
  - 完整匹配
  - 通配符匹配: `server_name: *.kamier.top;`
  - 通配符结束匹配: `server_name: www.kamier.*;`
  - 正则匹配
- 多用户二级域名
- 短网址
- HttpDns(基于http的dns,c/s架构或手机APP)

### 反向代理
- 网关, 代理与反向代理
- 反向代理在系统架构中的应用场景
- Nginx的反向代理配置
- 基于反向代理的负载均衡器
- 负载均衡策略

### 反向代理/负载均衡
```shell
# 反向代理配置
http {
	# 导入其他文件. mime.types: 定义解析类型
    include       mime.types;
    # 默认类型
    default_type  application/octet-stream;
    # 数据零拷贝
    sendfile        on;
    # 长连接超时时间
    keepalive_timeout  65;
    server {
        listen       80;
        server_name  localhost;
        location / {
            # 反向代理服务器, 如果没有配置https, 不能代理到https协议的服务器
            proxy_pass  http://www.baidu.com;
            root   html;
            index  index.html index.htm;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

- 负载均衡策略: 轮询, 权重
```shell
upstream server_list {
    #轮询(无法保持会话)  
    server 192.168.31.22:80;
    server 192.168.31.23:80;
    #权重  
    server 192.168.31.22:80 weight=2;
    server 192.168.31.23:80 weight=10;
    #使用down下线服务器   
    server 192.168.31.24:80 weight=10 down;
    #备用机器
    server 192.168.31.25:80 weight=10 backup;
    #ip_hash
    #least_conn:最少连接访问
    #fair:根据后端服务器响应时间转发请求
    #url_hash:根据用户访问的url请求转发请求
}
http {
	# 导入其他文件. mime.types: 定义解析类型
    include       mime.types;
    # 默认类型
    default_type  application/octet-stream;
    # 数据零拷贝
    sendfile        on;
    # 长连接超时时间
    keepalive_timeout  65;
    server {
        listen       80;
        server_name  localhost;
        location / {
            # 反向代理服务器, 如果没有配置https, 不能代理到https协议的服务器
            proxy_pass  http://server_list;
            root   html;
            index  index.html index.htm;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

## 动静分离
- 动静分离的场景
- 动静分离原理
- Nginx动静分离配置
```shell
http {
	# 导入其他文件. mime.types: 定义解析类型
    include       mime.types;
    # 默认类型
    default_type  application/octet-stream;
    # 数据零拷贝
    sendfile        on;
    # 长连接超时时间
    keepalive_timeout  65;
    server {
        listen       80;
        server_name  localhost;
        location / {
            # 反向代理服务器, 如果没有配置https, 不能代理到https协议的服务器
            proxy_pass  http://www.baidu.com;
            root   html;
            index  index.html;
        }
        location / {
            proxy_pass  http://www.baidu.com;
        }
        location /css {
            root html;
            index index.htlm index.htm;``
        }
        location ~*/(js|img|css) {
            root html;
            index index.html;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```