package enc

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"math/big"
)

/**
Given a byte array of ASCII values transform it
into a readable string.
 */
func AsciToStr(nums []byte) string {
	var res string = ""
	for _, n := range nums {
		res = res + string(n)
	}
	return res
}

/**
Given a string in HEX notation
decode it into a readable ASCII string.
 */
func HexToStr(src string) string {
	if src == "" {
		return ""
	}
	bs, err := hex.DecodeString(src)
	if err != nil {
		return ""
	}
	return AsciToStr(bs)
}

/**
Given a string encoded with HEX notation
encode it into Base64 notation.
 */
func HexToBase64(hexStr string) string {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		return ""
	}
	v := base64.StdEncoding.EncodeToString(b)
	return v
}

func Base64ToStr(b string) string {
	src := []byte(b)
	dst := make([]byte, len(src))
	count, err := base64.StdEncoding.Decode(dst, src)
	if err != nil || count == 0  {
		return ""
	}
	return string(bytes.Trim(dst, "\x00"))
}

/**
Given a big int (base 10 or 16) decode it
into a readable message.
 */
func NumToString(n *big.Int) string {
	b := n.Bytes() //ascii bytes
	return AsciToStr(b)
}

func strToInt(str string) string {
	return ""
}
