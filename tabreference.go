package main

import (
	"gopkg.in/mgo.v2/bson"
)

type TabReference struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Artist    string        `json:"artist" bson:"artist"`
	Url       string        `json:"url" bson:"url"`
	PageViews int           `json:"page_views" bson:"page_views"`
	Rating    float64       `json:"rating" bson:"rating"`
	IpfsHash  string        `json:"ipfs_hash" bson:"ipfs_hash"`
}

type TabReferences []TabReference
