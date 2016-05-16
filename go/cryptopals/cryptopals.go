package cryptopals

import (
  "encoding/hex"
)

// Displays a score of how likely a string is to be english based on
// the frequency rates of individual letters.
// Based on https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_the_English_language
func score(x string) float64 {
  var freqs = map[string]float64 {
    "a": 8.167,
    "b": 1.492,
    "c": 2.782,
    "d": 4.253,
    "e": 12.702,
    " ": 13,
    "f": 2.228,
    "g": 2.015,
    "h": 6.094,
    "i": 6.966,
    "j": 0.153,
    "k": 0.772,
    "l": 4.025,
    "m": 2.406,
    "n": 6.749,
    "o": 7.507,
    "p": 1.929,
    "q": 0.095,
    "r": 5.987,
    "s": 6.327,
    "t": 9.056,
    "u": 2.758,
    "v": 0.978,
    "w": 2.361,
    "x": 0.150,
    "y": 1.974,
    "z": 0.074,
    "A": 8.167,
    "B": 1.492,
    "C": 2.782,
    "D": 4.253,
    "E": 12.702,
    "F": 2.228,
    "G": 2.015,
    "H": 6.094,
    "I": 6.966,
    "J": 0.153,
    "K": 0.772,
    "L": 4.025,
    "M": 2.406,
    "N": 6.749,
    "O": 7.507,
    "P": 1.929,
    "Q": 0.095,
    "R": 5.987,
    "S": 6.327,
    "T": 9.056,
    "U": 2.758,
    "V": 0.978,
    "W": 2.361,
    "X": 0.150,
    "Y": 1.974,
    "Z": 0.074,
  }

  var score float64 = 0;
  for _, char := range x {
    score += freqs[string(char)]
  }

  return score;
}

// XORs a hex encoded String with all single characters (0-255), and returns
// the string most likely to be english, and as well it's total score
// and distance from best compititor.
// This is Challenge 3 & 4.
func CrackSingleXor(x string) (string, float64, float64, byte) {
  x_bytes, x_err := hex.DecodeString(x);

  if x_err != nil {
    panic(x_err);
  }

  dst := x_bytes;
  var top_score float64 = 0;
  var distance float64 = 0;
  winner := string(dst[:]);
  var winning_key byte;
  for i := 0; i < 255; i++ {
    key := byte(i);

    for i := 0; i < len(x_bytes); i++ {
      dst[i] = x_bytes[i] ^ key;
    }

    score := score(string(dst[:]));
    if score > top_score {
      distance = score - top_score;
      top_score = score;
      winning_key = key;
      winner = string(dst[:]);
    }
  }

  return winner, top_score, distance, winning_key;
}

// Encrypts a given plaintext (line by line) with a repeating XOR.
// This is Challenge 5
func EncryptRepeatingXor(plaintext string, key string) (string) {
  keybytes := []byte(key)
  plainbytes := []byte(plaintext)
  cipherbytes := plainbytes

  for i := 0; i < len(plainbytes); i++ {
    cipherbytes[i] = plainbytes[i] ^ keybytes[i%len(keybytes)]
  }
  return hex.EncodeToString(cipherbytes)
}
