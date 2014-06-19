package main

import (
	"fmt"
	"reflect"
)

type Dish struct {
	Id int
	Name string
	Query func()
}

func main() {
	for name, mtype := range attributes(&Dish{}) {
		fmt.Printf("Name: %s, Type: %s\n", name, mtype.Name())
	}		
}

func attributes(x interface{}) (map[string]reflect.Type) {
	typ := reflect.TypeOf(x)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	
	attrs := make(map[string]reflect.Type)
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}
	
	for i := 0; i < typ.NumField(); i++ {
		if f := typ.Field(i); !f.Anonymous {
			attrs[f.Name] = f.Type
		}
	}
	
	return attrs
}
