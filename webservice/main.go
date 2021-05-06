// @title Simple Rest service
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"rest-service/api"
	_ "rest-service/docs"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", api.homePage)
	myRouter.HandleFunc("/songs", api.getAllSongs)
	myRouter.HandleFunc("/song", api.updateSong).Methods("PUT")
	myRouter.HandleFunc("/song", api.createNewSong).Methods("POST")
	myRouter.HandleFunc("/song/{id}", api.deleteSong).Methods("DELETE")
	myRouter.HandleFunc("/song/{id}", api.getSong)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	r := gin.New()
	url := ginSwagger.URL(goDotEnvVariable("GIN_URL")) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	go r.Run()
	fmt.Println("Rest API v2.0 - Mux Routers")
	go handleRequests()
	//starting grpc server
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fmt.Println("gRPC Server initialized")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
