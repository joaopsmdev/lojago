package controllers

import (
	"Projetos_GO/Loja/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // traz template para visualizar
// template.Must encapsula template e retorna mensagem de erro
// template.ParseGlob traz tempolate da basta

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProdutos()         //traz todos os produtos e armazena na váriavel
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) //exibe todos a variavem dentro do template
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	idDoProdutoConv, err := strconv.Atoi(idDoProduto)
	if err != nil {
		log.Println("erro na conversão", err)
	}
	models.DeleteProduto(idDoProdutoConv)
	http.Redirect(w, r, "/", 301)
}
func Update(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.AlteraProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

// idDoProdutoConv, err := strconv.Atoi(idDoProduto)
// if err != nil {
// 	log.Println("erro na conversão", err)
// }
