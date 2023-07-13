package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
}

func main() {
	router := gin.Default()

	router.GET("/address", addressHandler)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("falha ao iniciar o servidor:", err)
	}
}

func addressHandler(c *gin.Context) {
	scrapedData, err := scrapeWebsite()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro na coleta de dados"})
		return
	}

	c.JSON(http.StatusOK, scrapedData)
}

func scrapeWebsite() (Address, error) {
	var data Address

	response, err := http.Get("https://viacep.com.br/ws/14570000/json/")
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
