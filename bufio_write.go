package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("halo world!\n")
	writer.WriteString("selamat belajar")
	writer.Flush()
}
