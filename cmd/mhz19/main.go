package main

import (
	"fmt"
	"os"

	"github.com/kebhr/mhz19"
)

func main() {
	m := mhz19.MHZ19{}
	if err := m.Connect(); err != nil {
		fmt.Printf("{\"error\": \"%s\"}", err)
		os.Exit(0)
	}
	v, err := m.ReadCO2()
	if err != nil {
		fmt.Printf("{\"error\": \"%s\"}", err)
		os.Exit(0)
	}
	fmt.Printf("{\"co2\": \"%d\"}", v)
}
