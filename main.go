// This is my first go program :-)
// It shows Hello World:
//      1.- in German
//      2.- in Greek
// exporting the services from internal packages
// experimenting with packages initializator messages and runes
package main

import (
  "fmt"
  _ "github.com/aliriopos/HalloWelt/pkg1"
  "github.com/aliriopos/HalloWelt/pkg2"
)

func init () {
  fmt.Println("Inicializando main...")
}

func main() {
  saludo := "Hola, saludo desde main..."
  fmt.Println(saludo)
  pkg2.Turrón()
}
