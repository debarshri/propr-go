package main

import (
	"testing"
	"./cluster"
)

func TestGetHost(*testing.T){

	if testing.Short() {

		cluster.AddHosts("test")
		cluster.HostsAsList()
		cluster.Hosts()

	}
}
