package models

import (
	"Projetos_GO/Loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaComBanco()

	selectDeProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) // adiciona produtos no slice

	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()

	insertProdutos, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProdutos.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
	//db.ConectaComBanco().Query("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4) ")
}

func DeleteProduto(id int) {
	db := db.ConectaComBanco()

	deletaProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletaProduto.Exec(id)
	defer db.Close()
}

func AlteraProduto(id string) {
	db := db.ConectaComBanco()

	UpdateProduto, err := db.Prepare()
}
