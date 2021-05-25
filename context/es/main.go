package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v2"
)
type Tweet struct {
	User string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.1.66:9200/"))
	if err != nil {
		fmt.Println("Connect es err:", err)
		return
	}

	fmt.Println("Connect es success")

	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()
	if err != nil{
		panic(err)
		return
	}
	fmt.Println("Insert success")
}
