package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"id", "name"})
	writer.Write([]string{"1", "yuki"})
	writer.Write([]string{"2", "nagae"})
	writer.Flush()
}
