package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	M := make(map[string]interface{})
	var I int
	var S string
	var D float64

	I = 10
	S = "this is a test"
	D = 3.141

	M["I"] = I
	M["S"] = S
	M["D"] = D

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(M)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(network.Bytes())

	M2 := make(map[string]interface{})
	err = dec.Decode(&M2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(M2)

	for key, val := range M2 {
		fmt.Print(key, ":")
		switch v := val.(type) {
		case int:
			fmt.Println("It is an int:", v)
		case string:
			fmt.Println("It is a string:", v)
		case float64:
			fmt.Println("It is a float64:", v)
		}
	}
}
