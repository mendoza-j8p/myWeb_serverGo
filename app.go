package main

import (
	"net/http"
	"fmt"
	"text/template"
)

type Users struct {
	Name string
	Skills string
	Age	int
}

func Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	
	user := Users{"Jenifer Mendoza", "Desarrollador web", 34}

	if err != nil {
		panic(err)
	} else {
		template.Execute(w, user)
	}
}


func main()  {
	http.HandleFunc("/", Index)


	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}

