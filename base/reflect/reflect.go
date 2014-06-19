package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{"Rinc", 25}
	
	v1 := reflect.ValueOf(person.Name)
	fmt.Printf("value of name: %v\n", v1)
	fmt.Printf("settability of name: %v\n", v1.CanSet())
	
	// Use pointer and Elem to modify value:
	v2 := reflect.ValueOf(&person.Name).Elem()
	fmt.Printf("settability of name: %v\n", v2.CanSet())
	v2.SetString("Rinc Liu")
	fmt.Printf("value of name: %v\n", person.Name)

	for aName, aType := range attributes(&person) {
		fmt.Printf("attrName: %s, attrType: %s\n", aName, aType.Name())
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
	
	// Travel all fields:
	for i := 0; i < typ.NumField(); i++ {
		if f := typ.Field(i); !f.Anonymous {
			attrs[f.Name] = f.Type
		}
	}
	
	return attrs
}
