package main

import (
	"fmt"
	"sync"
)

var basicInstance *BasicSingleton
var once sync.Once

type BasicSingleton struct {
}

func GetBasicInstance() *BasicSingleton {

	once.Do(func() {
		fmt.Println("Creating basic singleton Instance")
		basicInstance = &BasicSingleton{}
	})

	fmt.Println("Returning basic singleton Instance")
	return basicInstance
}
