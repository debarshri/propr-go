package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type args struct{
	Bucket, Key, Value string
}

type response struct{
	Value string
}

func main() {

	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply = new(response)

	args := &args{ Bucket:"buckk1", Key:"asdasd", Value:"Asda",}

	err = client.Call("SlaveServer.Get", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(reply.Value)

}
