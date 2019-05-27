package main

import (
	"fmt"

	"gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://10.39.6.202:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	fmt.Println("conn es succ")
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("insert succ")
}
