package models

import "gopkg.in/mgo.v2/bson"

// Wish Represents a wish, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Wish struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	UserID      string        `bson:"userId" json:"userId"`
	Items       []WishItem    `bson:"items" json:"items"`
}

// WishItem Represent a Wish item
type WishItem struct {
	ID   int    `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
	Note string `bson:"note" json:"note"`
	URL  string `bson:"url" json:"url"`
}
