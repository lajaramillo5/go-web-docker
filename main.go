package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Alumno struct {
	Universidad                     string `json:"universidad"`
	Curso                           string `json:"curso"`
	Alumno                          string `json:"alumno"`
	Período                         string `json:"perido"`
	Lenguajedeprogramacionpreferido string `json:"lenguaje"`
	AspiraciónPostGraduación        string `json:"aspiracion"`
}

// let's declare a global Alumnos array
// that we can then populate in our main function
// to simulate a database
type Alumnos []Alumno

func allAlumnos(w http.ResponseWriter, r *http.Request) {
	Alumnos := Alumnos{
		Alumno{Universidad: "UTPL",
			Curso:                           "Procesos de Ingeniería de Software",
			Alumno:                          "Luis jaramillo Uday",
			Período:                         "Abr/Ago 2021",
			Lenguajedeprogramacionpreferido: "Python",
			AspiraciónPostGraduación:        "Trabajar en la universidad"},
	}
	fmt.Println("mi primer endpoint")

	json.NewEncoder(w).Encode(Alumnos)
}

func handleRequests() {
	http.HandleFunc("/", allAlumnos)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func main() {
	handleRequests()
}
