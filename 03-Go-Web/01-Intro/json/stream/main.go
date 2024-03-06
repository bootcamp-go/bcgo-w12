package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	// Simulando una fuente de datos de flujo continuo con un pipe
	r, w := io.Pipe()

	// Simulación de escritura de datos JSON en un goroutine
	// para imitar datos que llegan en forma de stream
	go func() {
		for i := 0; i < 5; i++ {
			// Escribe un objeto JSON en el Pipe
			obj := map[string]int{"number": i}
			if err := json.NewEncoder(w).Encode(obj); err != nil {
				log.Fatal(err)
			}
			// Espera un poco antes de enviar el siguiente objeto
			time.Sleep(1 * time.Second)
		}
		w.Close() // Cierra el Pipe cuando termines de enviar los datos
	}()

	// Crear un json.Decoder para leer del Pipe
	decoder := json.NewDecoder(r)

	for {
		var obj map[string]int
		if err := decoder.Decode(&obj); err == io.EOF {
			break // Sal del bucle si no hay más datos que leer
		} else if err != nil {
			log.Fatal(err) // Manejar otros posibles errores
		}

		fmt.Println(obj) // Procesar el objeto JSON leído
	}

	fmt.Println("Fin del flujo de datos JSON.")
}

