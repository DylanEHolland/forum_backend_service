package main

import (
	"context"
	"fmt"
)

type post struct {
	id      int
	message string
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

func getPost(id int) *post {
	conn := db_connect()

	var node *post
	err := conn.QueryRow(context.Background(), "select id, message from posts where id = $1", id).Scan(&node.id, &node.message)
	conn.Close(context.Background())
	if err != nil {
		return node
	}

	return nil
}
