package dao

import (
	"log"

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

func (m *WishesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *WishesDAO) FindAll() ([]Wish, error) {
	var wishes []Wish
	err := db.C(COLLECTION).Find(bson.M{}).All(&wishes)
	return wishes, err
}

func (m *WishesDAO) FindById(id string) (Wish, error) {
	var wish Wish
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&wish)
	return wish, err
}

func (m *WishesDAO) Insert(wish Wish) error {
	err := db.C(COLLECTION).Insert(&wish)
	return err
}

func (m *WishesDAO) Delete(wish Wish) error {
	err := db.C(COLLECTION).Remove(&wish)
	return err
}

func (m *WishesDAO) Update(wish Wish) error {
	err := db.C(COLLECTION).UpdateId(wish.ID, &wish)
	return err
}
