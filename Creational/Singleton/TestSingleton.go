package main

import "fmt"

func main() {

	fmt.Println("Basic singleton result")
	for i := 0; i < 10; i++ {
		go GetBasicInstance()
	}

	fmt.Println("Dynamic singleton result")

	type ServerAddress struct {
		IP   string
		Port int
	}

	for i := 0; i < 10; i++ {
		go GetInstance(ServerAddress{}, ServerAddress{
			IP:   "com.example.org",
			Port: 8080,
		})
	}

	instance1 := GetInstance(ServerAddress{}, ServerAddress{
		IP:   "com.example.org",
		Port: 8080,
	}).(*Singleton)

	instance2 := GetInstance(ServerAddress{}, ServerAddress{
		IP:   "com.example.net",
		Port: 8181,
	}).(*Singleton)

	fmt.Println("Are Instance1 and Instance2 the same basicInstance?", instance1 == instance2)
	fmt.Println("Instance 1:", instance1.Data.(ServerAddress))
	fmt.Println("Instance 2:", instance2.Data.(ServerAddress))
	fmt.Scanln()
}
