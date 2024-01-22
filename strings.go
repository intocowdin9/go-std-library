package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("eko kurniawan", "eko"))
	fmt.Println(strings.Split("eko kurniawan", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan"))
	fmt.Println(strings.ToUpper("Eko Kurniawan"))
	fmt.Println(strings.Trim("       eko kurniawan      ", " "))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan", "Kurniawan", "Khannedy"))
}
