package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func contarLetra(filename string, letra string, wg *sync.WaitGroup, resultado chan string, parciales chan int) {
	defer wg.Done()

	cmd := exec.Command("grep", "-o", letra, filename)
	output, err := cmd.Output()
	if err != nil {
		resultado <- fmt.Sprintf("Error al ejecutar grep en el archivo %s: %v", filename, err)
		return
	}

	count := len(strings.Split(string(output), "\n")) - 1 // Restar 1 por la línea vacía

	resultado <- fmt.Sprintf("La letra '%s' aparece %d veces en el archivo %s.", letra, count, filename)
	parciales <- count
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <letra> <archivo1> <archivo2> ...")
		return
	}

	letra := os.Args[1]
	resultado := make(chan string)
	parciales := make(chan int)
	var wg sync.WaitGroup

	var wgParciales sync.WaitGroup

	wgParciales.Add(1)
	go func() {
		defer wgParciales.Done()
		totalApariciones := 0
		for parcial := range parciales {
			totalApariciones += parcial
		}
		fmt.Printf("Total de apariciones de la letra '%s' en todos los archivos: %d\n", letra, totalApariciones)
	}()

	go func() {
		for res := range resultado {
			fmt.Println(res)
		}
	}()

	for _, filename := range os.Args[2:] {
		wg.Add(1)
		go contarLetra(filename, letra, &wg, resultado, parciales)
	}

	wg.Wait()
	close(resultado)
	close(parciales)

	wgParciales.Wait()
}
