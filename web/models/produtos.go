package models

import (
	"github.com/regisrocha3/web/db"
)

type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produtos {
	db := db.ConectaBancoDeDados()
	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produtos{}
	produtos := []Produtos{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func BuscarProdutoPorId(id string) Produtos {
	db := db.ConectaBancoDeDados()
	produtoQuery, err := db.Query("select id, nome, quantidade, descricao, preco from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	var produtoParaAtualizar = Produtos{}

	for produtoQuery.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoQuery.Scan(&id, &nome, &quantidade, &descricao, &preco)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
	}
	defer db.Close()
	return produtoParaAtualizar
}

func Insert(produto Produtos) {
	db := db.ConectaBancoDeDados()

	insert, err := db.Prepare("insert into produtos (nome, preco, quantidade, descricao) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(produto.Nome, produto.Preco, produto.Quantidade, produto.Descricao)

	defer db.Close()
}

func Update(produto Produtos) {
	db := db.ConectaBancoDeDados()

	update, err := db.Prepare("update produtos set nome = $1, preco = $2, quantidade = $3, descricao = $4 where id = $5")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(produto.Nome, produto.Preco, produto.Quantidade, produto.Descricao, produto.Id)

	defer db.Close()
}

func Delete(id string) {
	db := db.ConectaBancoDeDados()

	delete, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}
