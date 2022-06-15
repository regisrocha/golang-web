package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/regisrocha3/web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
