package main

import (
	"fmt"
	"net/http"

	"github.com/mauFade/revo/web"
)

const PORT = "8081"

func main() {
	httpHanlder := web.NewHttpHandler()

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: httpHanlder,
	}

	fmt.Println("\nREST API running at port: ", PORT)

	err := server.ListenAndServe()

	if err != nil {
		panic("Error starting application: " + err.Error())
	}
}
