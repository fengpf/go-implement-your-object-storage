# go-implement-your-object-storage

### 在同一台服务器上绑定多个地址

```shell
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:1 10.0.2.15/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:2 10.0.2.16/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:3 10.0.2.17/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:4 10.0.2.18/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:5 10.0.2.19/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:6 10.0.2.20/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:7 10.0.3.1/16
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# ifconfig enp0s3:8 10.0.3.2/16

```

### 安装rabbitmq-server
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# apt-get install rabbitmq-server
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# service rabbitmq-server start
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# rabbitmq-plugins enable rabbitmq_management
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# service rabbitmq-server restart 
```

### 查看rabbitmq server 启动
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# lsof -i:15672
COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
beam.smp 11800 rabbitmq   55u  IPv4 180882      0t0  TCP *:15672 (LISTEN)
```
### 查看amqp server 启动
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# lsof -i:5672
COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
beam.smp 11800 rabbitmq   54u  IPv6 180440      0t0  TCP *:amqp (LISTEN)
```

### 添加用户test，密码test
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# rabbitmqctl add_user test test
Creating user "test"
```

### 给test用户添加访问所有exchange 的权限
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# rabbitmqctl set_permissions -p / test ".*" ".*" ".*"
Setting permissions for user "test" in vhost "/"
```

### 下载rabbitmqadmin 管理共工具
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# wget 10.0.2.15:15672/cli/rabbitmqadmin
2019-10-16 16:05:12 (92.4 MB/s) - 已保存 “rabbitmqadmin” [36192/36192])
```

### 创建dataServer/apiServer 两个exchange
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# python3 rabbitmqadmin declare exchange name=dataServer type=fanout
exchange declared
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# python3 rabbitmqadmin declare exchange name=apiServer type=fanout
exchange declared
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# python3 rabbitmqadmin --vhost=/ --username=test --password=test  declare exchange name=dataServer type=fanout
```

### 查看所有exchange
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# python3 rabbitmqadmin list exchanges
+--------------------+---------+
|        name        |  type   |
+--------------------+---------+
|                    | direct  |
| amq.direct         | direct  |
| amq.fanout         | fanout  |
| amq.headers        | headers |
| amq.match          | headers |
| amq.rabbitmq.log   | topic   |
| amq.rabbitmq.trace | topic   |
| amq.topic          | topic   |
| apiServer          | fanout  |
| dataServer         | fanout  |
+--------------------+---------+
```
### 查看所有用户
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# rabbitmqctl  list_users
Listing users
guest	[administrator]
test	[]
administrator	[administrator]
```
### 查看某个用户的权限
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# rabbitmqctl  list_user_permissions  test
Listing permissions for user "test"
/	.*	.*	.*
```

### 创建STORAGE_ROOT 目录及其子目录objects
`
root@fpf-VirtualBox:/data/app/go/src/go-object-storage# for i in `seq 1 6`; do mkdir -p /tmp/$i/objects; done
`
### 注册amqp server
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# export RABBITMQ_SERVER=amqp://test:test@10.0.2.15:5672
```

### 启动dataServer 
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.2.15:12345 STORAGE_ROOT=/tmp/1 go run dataServer/dataServer.go
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.2.16:12345 STORAGE_ROOT=/tmp/2 go run dataServer/dataServer.go
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.2.17:12345 STORAGE_ROOT=/tmp/3 go run dataServer/dataServer.go
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.2.18:12345 STORAGE_ROOT=/tmp/4 go run dataServer/dataServer.go
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.2.19:12345 STORAGE_ROOT=/tmp/5 go run dataServer/dataServer.go
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.2.20:12345 STORAGE_ROOT=/tmp/6 go run dataServer/dataServer.go
```

### 启动 apiServer
```
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.3.1:12345 go run apiServer/apiServer.go &
root@fpf-VirtualBox:/data/app/go/src/go-object-storage/chapter2# LISTEN_ADDRESS=10.0.3.2:12345 go run apiServer/apiServer.go &
```

### 访问服务节点 10.0.3.2:12345,put一个名为test的对象
```
curl -v 10.0.3.1:12345/objects/test -X PUT -d "this is object"
```

### 访问服务节点 10.0.3.1:12345，使用locate 命令查看test对象被保存在哪个服务节点上
```
curl -v 10.0.3.1:12345/locate/test
```

### 访问服务节点 10.0.3.2:12345,获取test对象内容
```
curl -v 10.0.3.2:12345/objects/test
```