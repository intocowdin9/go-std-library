package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("exo"))

	fmt.Println(regex.FindAllString("eko edi elo eku", 4))
}
