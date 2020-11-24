# mhz19
Read CO2 concentration from MH-Z19 on Raspberry Pi.

## Installation
```shell script
$ go get github.com/kebhr/mhz19/cmd/mhz19
```

## Usage
```shell script
$ mhz19
{"co2": "554"}
```

## Use as a library
```go
package main

import (
	"log"
	"github.com/kebhr/mhz19"
)

func main() {
	m := mhz19.MHZ19{}
	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}
	v, err := m.ReadCO2()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CO2 Concentration: %dppm\n", v)
}
```