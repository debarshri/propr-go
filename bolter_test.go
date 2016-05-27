package main

import (
	"testing"
	"fmt"
	"./bolter"
)

var b = bolter.New()
func TestCreateProperty(*testing.T){

	if testing.Short() {
		fmt.Print("Works")
	}
}

func BenchmarkBoltPropertyGet(*testing.B){


	for i := 0; i < 100; i++ {
		b.Get("test","asd")
	}

}

func BenchmarkBoltPropertyCreate(*testing.B){

	for i := 0; i < 1000; i++ {
		b.Create("test"+string(i))
	}
}

func BenchmarkBoltPropertyAdd(*testing.B){

	for i := 0; i < 1000; i++ {
		b.Put("test"+string(i), "Key"+string(i), "Value"+string(i))
	}
}

func BenchmarkBoltPropertyGetAll(*testing.B){

	for i := 0; i < 1000; i++ {
		b.GetAll("test"+string(i))
	}
}
