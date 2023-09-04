# <span style="color: pink;">Nginx</span>

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
firewalld-cmd --zone=public --add-port=80/tcp --permanent
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