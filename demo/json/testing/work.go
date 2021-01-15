package demo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Post struct {
	ID int
}

func StructToJson01() {
	p1 := Post{ID: 100}
	_, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Error %v ", err)
	}
}

func StructToJson02() {
	p1 := Post{ID: 100}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(&p1)
}
