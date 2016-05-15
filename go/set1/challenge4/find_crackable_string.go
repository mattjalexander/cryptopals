package main

import (
	"../../cryptopals"
  "bufio"
  "fmt"
  "os"
)

func main() {
  inFile, _ := os.Open("4.txt")
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  var top_score float64 = 0
  var king string = ""
  var original string = ""
  for scanner.Scan() {
    winner, score, _, _ := cryptopals.CrackSingleXor(scanner.Text())
    if  score > top_score {
      top_score = score
      king = winner
      original = scanner.Text()
    }
  }
  fmt.Println("CT:  " + original)
  fmt.Println("PT:  " + king)
}
