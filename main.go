package main

import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "api_sacolao/config"
	. "api_sacolao/dao"
	. "api_sacolao/models"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

var page []Page
var config = Config{}
var dao = FrutasDAO{}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("views").HTTPBox()))

	router.HandleFunc("/sacolao", adicionaFruta).Methods("POST")
	router.HandleFunc("/sacolao", adicionaFruta).Methods("OPTIONS")
	router.HandleFunc("/sacolao/frutas", selecionaFrutas).Methods("GET")
	router.HandleFunc("/sacolao/frutas", deletaFruta).Methods("DELETE")
	//router.HandleFunc("/sacolao/funcionarios", selecionaFuncionario).Methods("GET")
	router.HandleFunc("/sacolao/frutas{id}", selecionaFruta).Methods("GET")
	router.HandleFunc("/sacolao/funcionario", adicionaFuncionario).Methods("POST")
	router.HandleFunc("/sacolao/fornecedor", adicionaFornecedor).Methods("POST")
	//router.HandleFunc("/sacolao/fornecedores", selecionaFornecedores).Methods("GET")
	//router.HandleFunc("/sacolao", alteraFruta).Methods("PUT")

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./views/")))

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
}

var db *mgo.Database

/*func Handler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("views/index.html")
	/*data := map[string]string{
		"Title":      "Go Fruit Store",
		"Author":     "Brayan Antony",
		"Profission": "Developer",
	}
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, page)
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
}*/

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//Frutas
func selecionaFrutas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)

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
		respondWithError(w, http.StatusBadRequest, "ID do produto não encontrado.")
		return
	}
	respondWithJson(w, http.StatusOK, frutas)
}

func adicionaFruta(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)

	//BUSCA PRODUTOS NO BANCO PARA VALIDAÇÂO
	frutas, err := dao.FindAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, frutas)

	if r.Method == "OPTIONS" {
		return
	}

	defer r.Body.Close()
	var fruta Fruta

	err = json.NewDecoder(r.Body).Decode(&fruta)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Requisição inválida!"+err.Error())
		return
	}

	for _, item := range frutas {
		if fruta.Nome == item.Nome {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
	}

	fruta.ID = bson.NewObjectId()

	err = dao.Insert(fruta)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, fruta)

	//BUSCA PRODUTOS NO BANCO DE DADOS
}

func deletaFruta(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	defer r.Body.Close()
	var fruta Fruta

	err := json.NewDecoder(r.Body).Decode(&fruta)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if err := dao.Delete(fruta); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})
}

/*func alteraFruta(w http.ResponseWriter, r *http.Request) {

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

func errorDataDuplicated(err string) string {
	err = "Arquivo já  está cadastrado no sistema"
	return err
}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
