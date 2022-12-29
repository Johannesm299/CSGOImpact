package main

import (
	"os"
	"path/filepath"

	calc "github.com/Johannesm299/CSGOImpact/pkg/calc"
)

func main() {
	pwd, _ := os.Getwd()
	filepath := filepath.Join(pwd, "..", "..", "cologne_2022", "m5-nuke.dem")
	calc.CalculateImpact(filepath)
}
