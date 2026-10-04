package ui

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintJSON(v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func PrintSectionHeader(title string) {
	fmt.Println(Bold(Cyan("=== " + title + " ===")))
	fmt.Println()
}

func Fatal(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, Red("Error: ")+format+"\n", a...)
	os.Exit(1)
}
