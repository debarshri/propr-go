package main

import (
	"testing"
	"fmt"
	"./bolter"
)

func TestCreateProperty(*testing.T){
	if testing.Short() {
		fmt.Print("Works")
	}
}

func BenchmarkBoltPropertyGet(*testing.B){
	for i := 0; i < 1000; i++ {
		bolter.Get("test","asd")
	}
}

func BenchmarkBoltPropertyCreate(*testing.B){
	for i := 0; i < 1000; i++ {
		bolter.Create("test"+string(i))
	}
}

func BenchmarkBoltPropertyAdd(*testing.B){
	for i := 0; i < 1000; i++ {
		bolter.Put("test"+string(i), "Key"+string(i), "Value"+string(i))
	}
}

func BenchmarkBoltPropertyGetAll(*testing.B){
	for i := 0; i < 1000; i++ {
		bolter.GetAll("test"+string(i))
	}
}
