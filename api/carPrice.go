package api

import (
	"carPrice/data/car"
	"fmt"
	"strconv"

	"github.com/gocraft/web"
)

func (c *Api) GetCarPrices(rw web.ResponseWriter, req *web.Request) {
	//?id=...time=...

	idStr := req.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err.Error())
	}
	time := req.FormValue("time")

	if time != "" {
		carPriceOneTime := car.CarPriceByIdOneTime(id, time)
		bs, err := carPriceOneTime.ParseToJson()
		if err != nil {
			panic(err)
		}
		fmt.Fprint(rw, string(bs))
	} else {
		carPrice := car.CarPriceById(id)
		bs, err := carPrice.ParseToJson()
		if err != nil {
			panic(err)
		}
		fmt.Fprint(rw, string(bs))
	}

}
