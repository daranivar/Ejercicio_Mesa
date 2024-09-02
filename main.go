package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id              int
	Nombre          string
	Precio          float64
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func main() {

	router := gin.Default()

	router.GET("/producto", func(ctx *gin.Context) {
		datosBytes, err := ioutil.ReadFile("./productos.json")
		if err != nil {
			log.Fatal(err)
		}
		datosString := string(datosBytes)

		fmt.Println(datosString)

		var listProducts []product

		if err := json.Unmarshal([]byte(datosString), &listProducts); err != nil {
			log.Fatal(err)
		}
		ctx.JSON(200, listProducts)
	})

	router.Run(":8080")

}
