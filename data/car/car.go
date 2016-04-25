package car

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	time   string
	price  string
	remark string
)

type CarPrices struct {
	Id       int               `json:"id"`
	CarPrice []CarPriceOneTime `json:"carPrice"`
}

type CarPriceOneTime struct {
	Time   string `json:"time"`
	Price  string `json:"price"`
	Remark string `json:"remark"`
}

func CarPriceById(id int) *CarPrices {
	var carPrice []CarPriceOneTime
	carPrices := &CarPrices{
		Id:       id,
		CarPrice: carPrice,
	}

	carPriceOneTime := CarPriceOneTime{}

	db, err := sql.Open("mysql", "root:tantan@tcp(127.0.0.1:3306)/chexiang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row, err := db.Query(fmt.Sprintf("SELECT * FROM id_%d", carPrices.Id))
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&time, &price, &remark)
		if err != nil {
			panic(err.Error())
		}
		carPriceOneTime.Price = price
		carPriceOneTime.Time = time
		carPriceOneTime.Remark = remark
		carPrices.CarPrice = append(carPrices.CarPrice, carPriceOneTime)
	}
	return carPrices
}

func (self *CarPrices) ParseToJson() ([]byte, error) {
	return json.Marshal(&self)
}
