package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

type Win32_Product64 struct {
	Name string
}

func main() {
	var products []Win32_Product64
	q1 := "SELECT * FROM Win32_Product64 WHERE Name LIKE '%Avast%'"
	q2 := "SELECT * FROM Win32_Product64 WHERE Name LIKE '%Windows Defender%'"
	err := wmi.Query(q1, &products)
	if err == nil && len(products) > 0 {
		fmt.Println("Avast is installed")
	}
	err = wmi.Query(q2, &products)
	if err == nil && len(products) > 0 {
		fmt.Println("Windows Defender is installed")
	} else {
		fmt.Println("Neither installed")
	}
}
