package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("Type:", v.Type())        
	fmt.Println("Kind:", v.Kind())       
	fmt.Println("Value:", v.Float())      
}

