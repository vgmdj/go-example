# Examples for rpcx 3.0

A lot of examples for [rpcx](https://github.com/smallnest/rpcx/tree/v3.0) 3.0


## How to run
you should build rpcx with necessary tags, otherwise only need to install rpcx:

```sh
go get -u -v github.com/smallnest/rpcx/...
```

if you want to use "zookeeper" registry, you need to add tag `zookeeper`,

```sh
go get -u -v -tags "zookeeper" github.com/smallnest/rpcx/...
```

Similarly， if you want to use `etcd` registry and `quic` network, you need to :

```sh
go get -u -v -tags "etcd quic" github.com/smallnest/rpcx/...
```

You can install all features of rpcx with those below tags:

```sh
go get -u -v -tags "zookeeper etcd consul ping quic kcp reuseport" github.com/smallnest/rpcx/...
```

If you install succeefullly, you can run examples in this repository.

Enter one sub directory in this repository,  `go run server.go` in one terminal and `cd client; go run client.go` in another ternimal, and you can watch the run result.

For example,

```sh
cd 101basic
go run server.go
```

And

```sh
cd 101basic/client
go run client.go
```
