package main

import (
	"net"
	"net/rpc"
	"./cluster"
	"log"
)

func main(){


	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	slave := new(cluster.SlaveServer)
	rpc.Register(slave)
	rpc.Accept(inbound)
}
