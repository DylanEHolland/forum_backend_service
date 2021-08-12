package main

import (
	"context"
	"fmt"
)

type post struct {
	Id      int    `json: id`
	Message string `json: "message"`
}

func createPost(message string) {
	/*

	 */
	conn := db_connect()

	// var node *post
	statement := fmt.Sprintf("insert into posts (message) values ('%s');", message)
	_, err := conn.Exec(context.Background(), statement)
	conn.Close(context.Background())
	if err != nil {
		panic(err)
	}
}

func getPost(id int) post {
	conn := db_connect()

	var l_id int
	var l_message string
	err := conn.QueryRow(context.Background(), "select id, message from posts where id = $1", id).Scan(&l_id, &l_message)
	conn.Close(context.Background())
	if err != nil {
		panic(err)
	}

	node := post{Id: l_id, Message: l_message}
	return node
}
