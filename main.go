package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
)

var opt string

func init() {
	flag.StringVar(&opt, "o", "ieee", "output format")
}

func main() {
	flag.Parse()

	result := []interface{}{}

	// convert input to human readable string
	for _, arg := range flag.Args() {
		arg := strings.ToUpper(arg)

		if !strings.HasPrefix(arg, "0X") {
			result = append(result, arg)
			continue
		}

		a, err := hex.DecodeString(arg[2:])
		if err != nil {
			log.Fatal(err)
		}

		switch opt {
		case "ieee":
			r := math.Float32frombits(binary.BigEndian.Uint32(a))
			result = append(result, r)
		case "int":
			switch len(a) {
			case 2:
				r := binary.BigEndian.Uint16(a)
				result = append(result, r)
			case 4:
				r := binary.BigEndian.Uint32(a)
				result = append(result, r)
			case 8:
				r := binary.BigEndian.Uint64(a)
				result = append(result, r)
			default:
				log.Fatal("value too long")
			}
		default:
			fmt.Printf("cannot parse %v, invalid option", a)
		}
	}

	fmt.Println(result)
}
