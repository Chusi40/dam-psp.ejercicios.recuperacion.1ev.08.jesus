package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtener las variables de entorno
	envVars := os.Environ()

	// Iterar sobre cada variable y mostrarla
	for _, envVar := range envVars {
		fmt.Println(envVar)
	}
}
