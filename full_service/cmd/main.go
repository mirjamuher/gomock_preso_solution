package main

import (
	"fmt"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/product"
)

func main() {
	fmt.Println("Hello World")

	ps := product.ProductService{}
	if err := ps.CreateOrder(&product.Order{}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}