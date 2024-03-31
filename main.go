package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/huy2272/FormulaOneLite/internal/config"
	"github.com/huy2272/FormulaOneLite/internal/controller"
	"github.com/huy2272/FormulaOneLite/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// create server
	server := core.NewServer(config)

	coll := server.Db.Database("2024").Collection("Bahrain")
	driver := "S.PEREZ"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "Driver", Value: driver}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", driver)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	http_server := gin.Default()

	controller.RegisterRoutes(http_server)
	http_server.Run("localhost:8080") // listening on localhost:8080
}
