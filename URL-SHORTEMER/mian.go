package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/oustaa/url-shortner/internal/routes"
	"github.com/oustaa/url-shortner/internal/store"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "the port where the app will be running on")
	flag.Parse()

	_, err := os.Stat("./logs")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./logs", 0755)
			if err != nil {
				fmt.Printf("os.Mkdir: %v", err)
			}
		} else {
			fmt.Printf("os.Stat(\"./logs\"): %v\n", err)
		}
	}

	logFile, err := os.OpenFile(fmt.Sprintf("./logs/%v-logs.txt", time.Now().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("os.OpenFile (log file): %v\n", err)
	}
	defer logFile.Close()

	logger := log.New(io.MultiWriter(logFile, os.Stdout), "", log.Ldate|log.Ltime)

	db := store.DBConnection(logger)
	defer db.Close()

	r := routes.NewRouter()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	logger.Printf("server is listning on port %d", port)
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatalf("server.ListenAndServe: %v", err)
	}

}
