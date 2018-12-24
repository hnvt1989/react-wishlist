package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/hnvt1989/react-wishlist/go-server/config"
	. "github.com/hnvt1989/react-wishlist/go-server/dao"
	. "github.com/hnvt1989/react-wishlist/go-server/models"
)

var config = Config{}
var dao = WishesDAO{}

// Enable CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
}

// GET list of wishes
func AllWishesEndPoint(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	wishes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, wishes)
}

// GET a wish by its ID
func FindWishEndpoint(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
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
	// enableCors(&w)
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
	// enableCors(&w)
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
	// enableCors(&w)
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

func PreflightAddResource(w http.ResponseWriter, r *http.Request) {
	var empty []Wish
	// enableCors(&w)
	respondWithJson(w, http.StatusOK, empty)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
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
	r.HandleFunc("/wishes", PreflightAddResource).Methods("OPTIONS") //prelfight
	if err := http.ListenAndServe(":3003", r); err != nil {
		log.Fatal(err)
	}
}
