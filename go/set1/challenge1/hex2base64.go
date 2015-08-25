package main

import (
	"fmt"
  "encoding/hex"
  "encoding/base64"
)

func hex2base64(input string) string {
  bytes, err := hex.DecodeString(input);

  if err != nil {
    fmt.Println(err);
  }
  fmt.Println(string(bytes[:]));
  return base64.StdEncoding.EncodeToString(bytes);
}

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	output := hex2base64(input)
	fmt.Println("Given   : " + input)
	fmt.Println("Expected: " + expected)
	fmt.Println("Got     : " + output)

  if (output == expected) {
    fmt.Println("Hooray!");
  }
}
