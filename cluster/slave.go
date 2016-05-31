package cluster

import (
	"net"
	"log"
)

type SlaveServer string

type Args struct{
	Bucket, Key, Value string
}

type Response struct{
	Value string
}

func (s *SlaveServer) Get(args *Args, response *Response) error {
	response.Value = b.Get(args.Bucket, args.Key)
	return nil
}

func (s *SlaveServer) Put(args *Args, response *Response) error {

	log.Println("It works")
	b.Put(args.Bucket, args.Key, args.Value)

	response = &Response{"It works"}
	return nil
}

func (s *SlaveServer) Create(args *Args, response *Response) error {
	b.Create(args.Bucket)
	return nil
}

func (s *SlaveServer) IpAddress(args *Args, response *Response) error {
	ip := new(net.IPAddr)
	response.Value = ip.Network()
	return nil;
}
