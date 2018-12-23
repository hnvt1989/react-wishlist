package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/mlabouardy/wishes-restapi/config"
	. "github.com/mlabouardy/wishes-restapi/dao"
	. "github.com/mlabouardy/wishes-restapi/models"
)

var config = Config{}
var dao = WishesDAO{}

// GET list of wishes
func AllWishesEndPoint(w http.ResponseWriter, r *http.Request) {
	wishes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, wishes)
}

// GET a wish by its ID
func FindWishEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wish, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Wish ID")
		return
	}
	respondWithJson(w, http.StatusOK, wish)
}

// POST a new wish
func CreateWishEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wish Wish
	if err := json.NewDecoder(r.Body).Decode(&wish); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	wish.ID = bson.NewObjectId()
	if err := dao.Insert(wish); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, wish)
}

// PUT update an existing wish
func UpdateWishEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wish Wish
	if err := json.NewDecoder(r.Body).Decode(&wish); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(wish); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing wish
func DeleteWishEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wish Wish
	if err := json.NewDecoder(r.Body).Decode(&wish); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(wish); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/wishes", AllWishesEndPoint).Methods("GET")
	r.HandleFunc("/wishes", CreateWishEndPoint).Methods("POST")
	r.HandleFunc("/wishes", UpdateWishEndPoint).Methods("PUT")
	r.HandleFunc("/wishes", DeleteWishEndPoint).Methods("DELETE")
	r.HandleFunc("/wishes/{id}", FindWishEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3003", r); err != nil {
		log.Fatal(err)
	}
}
