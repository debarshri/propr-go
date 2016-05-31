package cluster

import (
	"../bolter"
	"container/list"
)

var(
	b = bolter.New()
)

func HostsAsList()(*list.List){
	return b.GetAllAsList("hosts")

	/**
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	*/
}

func CreateRemoteBucket(bucket string){
	list :=  b.GetAllAsList("hosts")

	for e := list.Front(); e != nil; e = e.Next() {
		//http dial
		//https://golang.org/pkg/net/rpc/
		//call a method to create remote buckets async
	}
}

func Put(bucket string, key string, value string){
	list :=  b.GetAllAsList("hosts")

	for e := list.Front(); e != nil; e = e.Next() {
		//http dial
		//https://golang.org/pkg/net/rpc/
		//call a method to create remote buckets async
	}
}

func Get(bucket string, key string, value string){
	list :=  b.GetAllAsList("hosts")

	for e := list.Front(); e != nil; e = e.Next() {
		//http dial
		//https://golang.org/pkg/net/rpc/
		//call a method to create remote buckets async
	}
}

func Hosts()(value string){
	return b.GetAll("hosts")
}

/**
string(time.Now().UnixNano() / 1000000)
timestamp
 */
func AddHosts(hostname string){
	 b.Put("hosts",hostname, "")
}