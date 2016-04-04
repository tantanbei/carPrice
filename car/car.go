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

type CarPrice struct {
	Id     int      `json:"id"`
	Time   []string `json:"time"`
	Price  []string `json:"price"`
	Remark []string `json:"remark"`
}

func CarPriceById(id int) *CarPrice {
	carPrice := &CarPrice{}

	db, err := sql.Open("mysql", "root:tantan@tcp(127.0.0.1:3306)/chexiang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	carPrice.Id = id

	row, err := db.Query(fmt.Sprintf("SELECT * FROM id_%d", carPrice.Id))
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&time, &price, &remark)
		if err != nil {
			panic(err.Error())
		}
		carPrice.Price = append(carPrice.Price, price)
		carPrice.Time = append(carPrice.Time, time)
		carPrice.Remark = append(carPrice.Remark, remark)
	}
	return carPrice
}

func (self *CarPrice) ParseToJson() ([]byte, error) {
	return json.Marshal(&self)
}
