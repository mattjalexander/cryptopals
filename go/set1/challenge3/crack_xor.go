package main

import (
  "fmt"
  "strconv"
  "../../cryptopals"
)

func main() {
  input    := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";

  winner, top_score, distance, key := cryptopals.CrackSingleXor(input);
  fmt.Println("Given  : " + input);
  fmt.Println("Winner : " + winner);
  fmt.Println("Score  : " + strconv.FormatFloat(top_score, 'f', 2, 64));
  fmt.Println("Lead   : " + strconv.FormatFloat(distance, 'f', 2, 64));
  fmt.Println("Key    : " + string(key));
}
