package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kusubooru/ne/rest"
	"github.com/kusubooru/ne/shimmie2"
	"github.com/kusubooru/shimmie/store"
)

var (
	httpAddr       = flag.String("http", "localhost:8081", "HTTP listen address")
	driverName     = flag.String("driver", "mysql", "database driver")
	dataSourceName = flag.String("datasource", "", "database data source")
)

func main() {
	flag.Parse()
	if *dataSourceName == "" {
		log.Println("No database datasource specified, exiting...")
		return
	}
	s := store.Open(*driverName, *dataSourceName)
	userService := shimmie2.NewUserService(s)
	api := rest.New(userService)

	log.Fatal(http.ListenAndServe(*httpAddr, api.Handlers()))
}
