package main

import (
	"html/template"
	"os"
)

type Pet struct {
	Type string
	Name string
}

type User struct {
	Name   string
	Age    int
	Weight float64
	Roles  []string
	Kin    map[string]string
	Pet    Pet
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:   "John Smith",
		Age:    45,
		Weight: 34.5,
		Roles:  []string{"administrator", "manager"},
		Kin: map[string]string{
			"Sister":  "Megan",
			"Brother": "Hubert",
		},
		Pet: Pet{"Dog", "Sam"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
