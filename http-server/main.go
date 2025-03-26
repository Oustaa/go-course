package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/oustaa/http-server-go/internal/app"
	"github.com/oustaa/http-server-go/internal/routers"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend out app")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

    //defer app.DB.Close()

	app.Logger.Printf("We are runing on port %d\n", port)

	r := routers.SetupRouters(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
