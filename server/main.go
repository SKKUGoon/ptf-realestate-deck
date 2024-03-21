package main

import (
	"fmt"
	"log"
	"melp-back/api"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create server
	ctrl := api.CreateController(os.Getenv("EXEC_STATE"))

	// Define route points
	common := ctrl.Conn.Group("/")
	src1 := ctrl.Conn.Group("/api/v1/")

	// Attach services
	ctrl.MountCommonService(common).MountV1Api(src1)

	// Serve
	log.Println(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
	serve := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")),
		Handler:        ctrl.Conn,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(serve.ListenAndServe())
}
