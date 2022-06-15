package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/regisrocha3/web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.BuscaProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produtos{}
		produto.Descricao = r.FormValue("descricao")
		produto.Nome = r.FormValue("nome")
		produto.Preco, _ = strconv.ParseFloat(r.FormValue("preco"), 64)
		produto.Quantidade, _ = strconv.Atoi(r.FormValue("quantidade"))

		models.Insert(produto)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.Delete(idProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	productId := r.URL.Query().Get("id")

	produto := models.BuscarProdutoPorId(productId)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func SaveEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produtos{}
		produto.Descricao = r.FormValue("descricao")
		produto.Nome = r.FormValue("nome")
		produto.Preco, _ = strconv.ParseFloat(r.FormValue("preco"), 64)
		produto.Quantidade, _ = strconv.Atoi(r.FormValue("quantidade"))
		produto.Id, _ = strconv.Atoi(r.FormValue("id"))

		models.Update(produto)
	}

	http.Redirect(w, r, "/", 301)
}
