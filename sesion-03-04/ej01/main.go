package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// Verificar que se pasen los argumentos necesarios
	if len(os.Args) < 4 {
		fmt.Println("Uso: go run main.go <texto_a_buscar> <archivo_entrada> <archivo_salida>")
		return
	}

	texto := os.Args[1]
	archivoEntrada := os.Args[2]
	archivoSalida := os.Args[3]

	// Crear el comando "grep" con el texto a buscar y el archivo de entrada
	cmd := exec.Command("grep", texto, archivoEntrada)

	// Obtener el pipe de salida estándar (stdout) del comando
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creando el pipe stdout:", err)
		return
	}

	// Iniciar el comando
	if err := cmd.Start(); err != nil {
		fmt.Println("Error iniciando el comando:", err)
		return
	}

	// Leer la salida del comando (las líneas coincidentes)
	output, err := io.ReadAll(stdout)
	if err != nil {
		fmt.Println("Error leyendo la salida:", err)
		return
	}

	// Esperar a que el comando termine de ejecutarse
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error esperando a que termine el comando:", err)
		return
	}

	// Crear el archivo de salida
	fileOut, err := os.Create(archivoSalida)
	if err != nil {
		fmt.Println("Error creando el archivo de salida:", err)
		return
	}
	defer fileOut.Close()

	// Escribir la salida en el archivo de salida
	_, err = fileOut.Write(output)
	if err != nil {
		fmt.Println("Error escribiendo en el archivo de salida:", err)
		return
	}

	fmt.Println("Proceso completado con éxito.")
}