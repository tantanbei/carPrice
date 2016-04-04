package main

import (
	"carPrice/car"
	"fmt"
)

func main() {
	carPrice := car.CarPriceById(888)
	bs, err := carPrice.ParseToJson()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}
