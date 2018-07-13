package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "api_sacolao/config"
	. "api_sacolao/dao"
	. "api_sacolao/models"

	"github.com/gorilla/mux"
)

var config = Config{}
var dao = FrutasDAO{}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Handler)
	router.HandleFunc("/sacolao/frutas", selecionaFrutas).Methods("GET")
	//router.HandleFunc("/sacolao/funcionarios", selecionaFuncionario).Methods("GET")
	router.HandleFunc("/sacolao/{id}", selecionaFruta).Methods("GET")
	router.HandleFunc("/sacolao/frutas", adicionaFruta).Methods("POST")
	router.HandleFunc("/sacolao/funcionario", adicionaFuncionario).Methods("POST")
	router.HandleFunc("/sacolao/fornecedor", adicionaFornecedor).Methods("POST")
	//router.HandleFunc("/sacolao/fornecedores", selecionaFornecedores).Methods("GET")
	//router.HandleFunc("/sacolao", alteraFruta).Methods("PUT")
	//router.HandleFunc("/sacolao", deletaFruta).Methods("DELETE")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("views/index.html")
	data := map[string]string{
		"Title": "Go Store",
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

//Frutas
func selecionaFrutas(w http.ResponseWriter, r *http.Request) {
	frutas, err := dao.FindAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, frutas)
}

func selecionaFruta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	frutas, err := dao.FindById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID da fruta não encontrado!")
		return
	}
	respondWithJson(w, http.StatusOK, frutas)
}

func adicionaFruta(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var fruta Fruta

	err := json.NewDecoder(r.Body).Decode(&fruta)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Requisição inválida!")
		return
	}

	fruta.ID = bson.NewObjectId()
	if err := dao.Insert(fruta); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, fruta)
}

/*func alteraFruta(w http.ResponseWriter, r *http.Request) {

}

func deletaFruta(w http.ResponseWriter, r *http.Request) {

}*/

//Funcionarios
/*func selecionaFuncionario(w http.ResponseWriter, r *http.Request) {
	funcionarios, err := dao.FindAllF()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, funcionarios)
}*/

func adicionaFuncionario(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var funcionarios Funcionario

	err := json.NewDecoder(r.Body).Decode(&funcionarios)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Requisição inválida!")
		return
	}

	funcionarios.ID = bson.NewObjectId()
	if err := dao.InsertF(funcionarios); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, funcionarios)
}

//Fornecedores
func adicionaFornecedor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var fornecedores Fornecedor

	err := json.NewDecoder(r.Body).Decode(&fornecedores)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Requisição inválida!")
		return
	}

	fornecedores.ID = bson.NewObjectId()
	if err := dao.InsertFornecedor(fornecedores); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, fornecedores)
}

/*func selecionaFornecedores(w http.ResponseWriter, r *http.Request) {
	fornecedores, err := dao.GetFornecedor()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, fornecedores)
}*/

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
