package main

import (
	"database/sql"
	_ "fmt"

	_ "github.com/fabiocaettano74/calculate-price/internal/order/entity"
	"github.com/fabiocaettano74/calculate-price/internal/order/infra/database"
	"github.com/fabiocaettano74/calculate-price/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDto{
		ID:    "124",
		Price: 100,
		Tax:   10,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output.FinalPrice)

}

/*func main() {
	order, err := entity.NewOrder("123", 50, 1.002)
	if err != nil {
		panic(err)
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The final price is %f", order.FinalPrice)

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	repository := database.NewOrderRepository(db)
	err = repository.Save(order)

	if err != nil {
		panic(err)
	}
}*/
