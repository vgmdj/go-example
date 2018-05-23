
[![Build Status](https://travis-ci.org/wwcd/grpc-lb.svg?branch=master)](https://travis-ci.org/wwcd/grpc-lb)

# 说明

[gRPC服务发现&负载均衡](https://segmentfault.com/a/1190000008672912)中的例子, 经[wwcd](https://github.com/wwcd/grpc-lb)修订如下问题

- register中重复PUT, watch时没有释放导致的内存泄漏
- 退出时不能正常unregister

# 改动
- 异常退出判断由svr发起，改为注册后持续判断
- 添加服务状态，可根据服务状态调整负载策略


# 测试

## 启动ETCD

	# https://coreos.com/etcd/docs/latest/op-guide/container.html#docker

	export NODE1=192.168.1.21

	docker volume create --name etcd-data
	export DATA_DIR="etcd-data"

	REGISTRY=quay.io/coreos/etcd
	# available from v3.2.5
	# REGISTRY=gcr.io/etcd-development/etcd

	docker run -d\
	  -p 2379:2379 \
	  -p 2380:2380 \
	  --volume=${DATA_DIR}:/etcd-data \
	  --name etcd ${REGISTRY}:latest \
	  /usr/local/bin/etcd \
	  --data-dir=/etcd-data --name node1 \
	  --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
	  --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
	  --initial-cluster node1=http://${NODE1}:2380

## 启动测试程序
    # grpc目录
    $GOPATH/src/grpc

    # 分别启动服务端
    go run cmd/svr/svr.go - port 50001
    go run cmd/svr/svr.go - port 50002
    go run cmd/svr/svr.go - port 50003

    # 启动客户端
    go run cmd/cli/cli.go
