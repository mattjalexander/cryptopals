package main

import (
  "fmt"
  "encoding/hex"
//  "crypto/cipher"
)

// Challenge Two
func fixed_xor(x, y string) string{

  x_bytes, x_err := hex.DecodeString(x);

  if x_err != nil {
    fmt.Println(x_err);
  }
  fmt.Println(string(x_bytes[:]));

  y_bytes, y_err := hex.DecodeString(y);
  if y_err != nil {
    fmt.Println(y_err);
  }
  fmt.Println(string(y_bytes[:]));

  if len(x_bytes) != len(y_bytes) {
    fmt.Println("[FATAL] Cannot XOR two differently length strings.");
  }

  dst := x_bytes;
  for i := 0; i < len(x_bytes); i++ {
    dst[i] = x_bytes[i] ^ y_bytes[i];
  }

  fmt.Println(string(dst[:]));
  return hex.EncodeToString(dst);
}

func main() {
  input    := "1c0111001f010100061a024b53535009181c";
  XOR_ME   := "686974207468652062756c6c277320657965";
  expected := "746865206b696420646f6e277420706c6179";

  output := fixed_xor(input, XOR_ME);
  fmt.Println("Given   : " + input + " XOR " + XOR_ME);
  fmt.Println("Expected: " + expected);
  fmt.Println("Got     : " + output);

  if (output == expected) {
    fmt.Println("Hooray!");
  } else {
    fmt.Println("Boooo");
  }
}
