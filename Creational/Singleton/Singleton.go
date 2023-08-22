package main

import (
	"fmt"
	"reflect"
	"sync"
)

var lock sync.Mutex
var instances = make(map[interface{}]interface{})

type Singleton struct {
	DataType reflect.Type
	Data     interface{}
}

func GetInstance(dataType interface{}, data interface{}) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if instance, ok := instances[dataType]; ok {
		fmt.Println("returning Instance of", instance.(*Singleton).DataType)
		return instance
	}

	typ := reflect.TypeOf(dataType)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	instance := &Singleton{Data: data, DataType: typ}
	instances[dataType] = instance
	fmt.Println("Instance of", instance.Data, " has created")

	return instance
}
