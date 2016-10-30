package main

import (
	"fmt"

	"github.com/black13/gofun/cassandra"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	cassandra.Add("dr", "Dart")
	fmt.Println(cassandra.Get("dr"))
	languages := cassandra.GetAll()
	for _, v := range languages {
		fmt.Println(v)
	}
	fmt.Println(cassandra.Average(xs))
}
