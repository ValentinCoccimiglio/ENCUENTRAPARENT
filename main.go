package main

import (
	"fmt"
	"path/filepath"
)

func findParent(filesystem map[string]string, file1 string, file2 string) string {
	var resolvePath func(path string) string
	resolvePath = func(path string) string {
		if alias, ok := filesystem[path]; ok {
			return resolvePath(alias)
		}
		return path
	}
	//Obtener las rutas completas de los archivos
	file1 = resolvePath(file1)
	file2 = resolvePath(file2)

	// Obtener todos los directorios padres de file1
	parent1 := filepath.Dir(file1)

	// Iterar sobre directorios padres de file2  hasta encontrar el comun con file1
	dir := filepath.Dir(file2)
	for dir != "/" {
		if dir == parent1 {
			return dir
		}
		dir = filepath.Dir(dir)
	}

	return "" // No se encontró un directorio común

}

// Ejemplo de un sistema de archivos con alias

func main() {
	filesystem := map[string]string{
		"/var": "/a/b",
	}

	// Rutas de los archivos
	file1 := "/var/example/file1.txt"
	file2 := "/var/example/file2.txt"

	// Encontrar el directorio común
	commonParent := findParent(filesystem, file1, file2)

	fmt.Println("Directorio común:", commonParent)
}
