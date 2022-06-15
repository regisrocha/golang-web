package routes

import (
	"net/http"

	"github.com/regisrocha3/web/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/editar", controllers.Edit)
	http.HandleFunc("/saveEdit", controllers.SaveEdit)

}
