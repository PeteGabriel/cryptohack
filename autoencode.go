package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/petegabriel/cryptohack/enc"
	"math/big"
	"os"
	"strings"
)

const (
	BigInt= "bigint"
	Rot13 = "rot13"
	UTF8 = "utf-8"
	Base64 = "base64"
	HEX= "hex"
)


/**
This program allows for decoding on the fly by providing the following kind
of argument:
	{"type": "base64", "encoded": "bm9uZXZlbnRzX2dyYXlfQ2xpZmY="}

Based upon the type property, it decodes and sends out the following data:
	{"type": "base64", "decoded":"Ym05dVpYWmxiblJ6WDJkeVlYbGZRMnhwWm1ZPQ=="}

It supports the following types:
	base64
	bigint
	rot13
	utf-8
	hex
 */
func main() {
	fmt.Println("...")
	for true {
		reader := bufio.NewReader(os.Stdin)
		args, _ := reader.ReadString('\n')
		i := Data{}
		err := json.Unmarshal([]byte(args), &i)
		if err != nil {
			fmt.Println("Bad input")
		}
		switch i.T {
		case Rot13:
			fmt.Println(buildRes(i.T, strings.Map(enc.Rot13, i.D.(string))))
		case BigInt:
			b := []byte(i.D.(string))
			//remove hexadecimal notation
			bi, _ := new(big.Int).SetString(string(b[2:]), 16)
			fmt.Println(buildRes(i.T, enc.NumToString(bi)))
		case UTF8:
			//the default type assigned to is float64, I had to workaround this
			a := i.D.([]interface{})
			b := make([]uint8, len(a))
			for i := range a {
				b[i] = uint8(a[i].(float64))
			}
			fmt.Println(buildRes(i.T, enc.AsciToStr(b)))
		case Base64:
			fmt.Println(buildRes(i.T, enc.Base64ToStr(i.D.(string))))
		case HEX:
			fmt.Println(buildRes(i.T, enc.HexToStr(i.D.(string))))
		default:
			fmt.Println("bad type")
		}
	}
}

func buildRes(t string, res string) string {
	if j, err := json.Marshal(Decoded{T: t, D: res}); err != nil {
		return "bad encoding"
	}else {
		return string(j)
	}
}

type Data struct {
	T string `json:"type"`
	D interface{} `json:"encoded"`
}

type Decoded struct {
	T string `json:"type"`
	D interface{} `json:"decoded"`
}
