package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func main() {
	// encode
	hello :=`{"Name":"junbin", "Age":21, "Gender":true}`
	debyte := base64Encode([]byte(hello))
	fmt.Println(string(debyte))
	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}
