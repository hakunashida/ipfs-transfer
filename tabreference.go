package main

type TabReference struct {
	Name      string  `json:"name" bson:"name"`
	Artist    string  `json:"artist" bson:"artist"`
	Url       string  `json:"url" bson:"url"`
	PageViews int     `json:"page_views" bson:"page_views"`
	Rating    float64 `json:"rating" bson:"rating"`
	BucketId  string  `json:"bucket_id" bson:"bucket_id"`
	FildId    string  `json:"file_id" bson:"file_id"`
}

type TabReferences []TabReference
