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