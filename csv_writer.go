package main

import (
	"encoding/csv"
	"os"
)

func main() {

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
	_ = writer.Write([]string{"muhammad", "rafli"})

	writer.Flush()
}
