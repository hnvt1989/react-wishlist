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
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, PATCH, OPTIONS")
}

// GET list of wishes
func AllWishesEndPoint(w http.ResponseWriter, r *http.Request) {
	wishes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, wishes)
}

func AllWishesByUserIdEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wishes, err := dao.FindAllByUserId(params["id"])
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
	params := mux.Vars(r)
	err := dao.Delete(params["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func PreflightHandler(w http.ResponseWriter, r *http.Request) {
	var empty []Wish
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
	r.HandleFunc("/wishes/{id}", UpdateWishEndPoint).Methods("PATCH")
	r.HandleFunc("/wishes/{id}", DeleteWishEndPoint).Methods("DELETE")
	r.HandleFunc("/wishes/{id}", FindWishEndpoint).Methods("GET")
	r.HandleFunc("/wishes/userId={id}", AllWishesByUserIdEndPoint).Methods("GET")
	r.HandleFunc("/wishes", PreflightHandler).Methods("OPTIONS")      //prelfight
	r.HandleFunc("/wishes/{id}", PreflightHandler).Methods("OPTIONS") //prelfight for DELETE, PATCH
	if err := http.ListenAndServe(":3003", r); err != nil {
		log.Fatal(err)
	}
}
