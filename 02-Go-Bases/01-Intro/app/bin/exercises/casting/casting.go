package main

import "strconv"

/*
	Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.

	Necesita ayuda para:
	Detectar cuáles de estas variables que declaró el alumno son correctas.
	Corregir las incorrectas.
		var 1firstName string
		var lastName string
		var int age
		1lastName := 6
		var driver_license = true
		var person height int
		childsNumber := 2
*/

func main() {
	// original
	// var student1 string = "23"

	// solution
	// var student1 string = "23"
	// var student1 int = 23
	// student1 := 23
	// var student1 string = "John"
	// var student1 string = string(23) // "↨"

	var student1 string = strconv.Itoa(23) // "23"

	// numberOfStudents := 23
	// numberOfStudents := 25 // already redeclared numberOfStudents

	println(student1)
}