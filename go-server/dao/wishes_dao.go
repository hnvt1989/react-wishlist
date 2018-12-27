package dao

import (
	"log"

	. "github.com/hnvt1989/react-wishlist/go-server/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WishesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "wishes"
)

// Establish a connection to database
func (m *WishesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *WishesDAO) FindAll() ([]Wish, error) {
	var movies []Wish
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Find a wish by its id
func (m *WishesDAO) FindById(id string) (Wish, error) {
	var wish Wish
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&wish)
	return wish, err
}

// Insert a wish into database
func (m *WishesDAO) Insert(wish Wish) error {
	err := db.C(COLLECTION).Insert(&wish)
	return err
}

// Delete an existing wish
func (m *WishesDAO) Delete(id string) error {
	var wish Wish
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&wish)

	if err != nil {
		return err
	}

	deleteErr := db.C(COLLECTION).Remove(&wish)
	return deleteErr
}

// Update an existing wish
func (m *WishesDAO) Update(wish Wish) error {
	err := db.C(COLLECTION).UpdateId(wish.ID, &wish)
	return err
}
