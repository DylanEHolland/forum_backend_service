package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//routes()
	//createPost("This is a test")
	node := getPost(2)
	b, err := json.Marshal(node)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
