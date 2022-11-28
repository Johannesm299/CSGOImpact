package main

import (
	"os"
	"path/filepath"

	calc "github.com/Johannesm299/CSGOImpact/pkg/calc"
)

func main() {
	pwd, _ := os.Getwd()
	filepath := filepath.Join(pwd, "..", "..", "match730_003403354556619293004_0695644917_191.dem")
	calc.CalculateImpact(filepath)
}
