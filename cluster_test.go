package main

import (
	"testing"
	"fmt"
	"./cluster"
)

func TestGetHost(*testing.T){

	if testing.Short() {

		cluster.AddHosts("test")
		cluster.HostsAsList()
		cluster.Hosts()

	}
}
