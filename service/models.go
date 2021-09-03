package main

import (
	"context"
	"fmt"
)

type post struct {
	Id      int
	Message string
}

type user struct {
	Id       int    `json:"id"`
	First    string `json:"first_name"`
	Last     string `json:"last_name"`
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
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

func isUser(user_name string, pass_word string) int {
	con := db_connect()

	var id int
	err := con.QueryRow(context.Background(), "select id from users where user_name = $1 and pass_word = $2;", user_name, pass_word).Scan(&id)
	if err != nil {
		fmt.Println("Error in User: ", err, id)
	}

	return id
}

func createUser(u user) int {
	con := db_connect()
	statement := fmt.Sprintf("insert into users (user_name, pass_word) values ('%s', '%s') returning id;", u.Username, u.Password)
	// _, err := con.Exec(context.Background(), statement)
	id := 0
	err := con.QueryRow(context.Background(), statement).Scan(&id)
	con.Close(context.Background())
	if err != nil {
		panic(err)
	}

	return id
}
