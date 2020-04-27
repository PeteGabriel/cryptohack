package enc

import (
	"github.com/matryer/is"
	"math/big"
	"strings"
	"testing"

)

func TestAsciToStr(t *testing.T){
	is := is.New(t)
	exptd := "crypto"
	input := []byte { 99, 114, 121, 112, 116, 111}
	is.Equal(AsciToStr(input), exptd)
}

func TestHexToStr(t *testing.T){
	is := is.New(t)
	exptd := "crypto"
	input := "63727970746f"
	is.Equal(HexToStr(input), exptd)
}

func TestBytesToBase64(t *testing.T){
	is := is.New(t)
	exptd := "crypto/B"
	input := "72bca9b68fc1"
	is.Equal(HexToBase64(input), exptd)
}

func TestNumToBytes(t *testing.T){
	is := is.New(t)
	exptd := "HELLO"
	input, ok := new(big.Int).SetString("48454c4c4f", 16)
	is.True(ok == true)
	is.Equal(NumToString(input), exptd)
}

func TestRot13OverString(t *testing.T){
	is := is.New(t)
	s := "refngm_nhgubevgnevnavfzf_rivaprq"
	r := strings.Map(Rot13, s)
	is.Equal(r, "ersat`_authoritarianisms_evinced")
}