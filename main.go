package main


import (
	"debug/elf"
	"encoding/binary"
	"fmt"
)

func main() {
	data, err := ReadELFNote("abort.core")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func ReadELFNote(filename string) ([]byte, error) {
	f, err := elf.Open(filename)
	if err != nil {
		return nil, err
	}
	
}
