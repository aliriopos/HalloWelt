// Provides the following sercices:
//   - Turrón: to show the "turrón" of the exercise
package pkg2

import (
 "fmt"
)

func init () {
  fmt.Println("Inicializando pkg2...")
}

// To show hello world in German
// NOTE:
//   - It does not work with Turrón
//     func Turrón () { DOES NOT WORK
func Turron () {
  fmt.Println(" ---------- HALLO WELT ----------")
}

// To show hello world in Greek and utf8 codification
func TurronInGreek () {
    stringEnGriego := " ---------- Γειά σου Κόσμε ----------"
    runaEnGriego :=[]rune(stingEnGriego)

    fmt.Println(stringEnGriego)
    fmt.Println(runaEnGriego)
  }
