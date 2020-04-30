## 文件介绍

####  1.dokcer-compose.yml

三个节点的zookeeper集群

#### 2.kfk-manager.yml

包含三个节点的kafka集群和一个kafka集群管理工具kafkamanager

#### 3.kafka.yml

三个节点的kafka集群

#### 4.kfkmanager.yml

kafka集群管理工具

## 启动顺序

1. 启动zookeeper集群
 ```shell
 docker-compose up -d
 ```
2. 启动kafka集群
```shell
docker-compose -f kafka.yml -d
```
3. 此时可启动kafka-manager
```shell
docker-compose -f kfamanager.yml -d
```

或者2，3合并同时启动kafka集群和kafka-manager
```shell
docker-compose -f kfa-manager.yml -d
```
启动集群管理工具，进入localhost:9001使用。