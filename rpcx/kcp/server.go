//go run -tags kcp server.go
package main

import (
	"crypto/sha1"
	"flag"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
	kcp "github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

const cryptKey = "rpcx-key"
const cryptSalt = "rpcx-salt"

func main() {
	flag.Parse()

	pass := pbkdf2.Key([]byte(cryptKey), []byte(cryptSalt), 4096, 32, sha1.New)
	bc, err := kcp.NewAESBlockCrypt(pass)
	if err != nil {
		panic(err)
	}

	s := server.NewServer(server.WithBlockCrypt(bc))
	s.RegisterName("Arith", new(example.Arith), "")

	err = s.Serve("kcp", *addr)
	if err != nil {
		panic(err)
	}
}
