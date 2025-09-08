package main

import (
	"encoding/json"
	"fmt"

	"github.com/ncostamagna/go-http-utils/meta"

	"github.com/ncostamagna/go-http-utils/response"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mail      string `json:"mail"`
	Age       int    `json:"age"`
}

func main() {
	u := &User{
		FirstName: "Nahuel",
		LastName:  "Costamagna",
		Mail:      "nlcostamagna@gmail.com",
		Age:       32,
	}

	print(response.OK("", u, nil))

	print(response.OK("sarasa", u, nil))

	// 15 es el limite de paginacion por defecto en caso de no asignarle valor
	m, err := meta.New(1, 30, 100, "15")
	if err != nil {
		fmt.Println(err)
	}

	print(response.OK("sarasa", u, m))

	print(response.Created("", u, nil))
	print(response.Accepted("", u, nil))

	print(response.BadRequest("my error"))
	print(response.NotFound("my error"))
	print(response.InternalServerError("my error"))

}

func print(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}
