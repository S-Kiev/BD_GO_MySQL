package main

import (
	"fmt"
	"log"

	"github.com/S-Kiev/BD_GO_MySQL/pkg/product"
	"github.com/S-Kiev/BD_GO_MySQL/storage"
)

//Documentacion de MySQL:
//https://pkg.go.dev/github.com/go-sql-driver/mysql#section-documentation

func main() {
	driver := storage.MySQL
	storage.New(driver)
	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)

	m, err := serviceProduct.GetByID(4)
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(m)
}
