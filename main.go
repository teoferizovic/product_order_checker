package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jasonlvhit/gocron"
	"product_order_checker/model"
	"product_order_checker/processor"
	_ "github.com/go-sql-driver/mysql"
)

var conf model.Config

func updateOrders(){
	fmt.Println("orders")
}

func updateOrderProducts()  {
	fmt.Println("order_products")
}

func task() {

	db, err := sql.Open("mysql", conf.Username+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.DB)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	err = processor.UpdateOrder(db)

	if err != nil{
		panic(err)
	}

	err = processor.UpdateOrderProcessor(db)

	if err != nil{
		panic(err)
	}

}

func init(){
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v\n", conf)
}

func main() {
	s := gocron.NewScheduler()
	s.Every(10).Seconds().Do(task)
	//s.Every(1).Day().At("14:12").Do(task)
	<- s.Start()
}
