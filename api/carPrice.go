package api

import (
	"carPrice/data/car"
	"fmt"
	"strconv"

	"github.com/gocraft/web"
)

func (c *Api) GetCarPrices(rw web.ResponseWriter, req *web.Request) {
	//?id=...time=...

	str := req.FormValue("id")
	fmt.Println(str)
	time := req.FormValue("time")
	fmt.Println(time)
	n, _ := strconv.Atoi(str)
	carPrice := car.CarPriceById(n)
	bs, err := carPrice.ParseToJson()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(rw, string(bs))

}
