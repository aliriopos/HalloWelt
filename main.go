package main

import (
  "fmt"
  "github.com/aliriopos/halloweltMod/pkg1"
  "github.com/aliriopos/halloweltMod/pkg2"
)

func init () {
  fmt.Println("Inicializando main...")
}

func main() {
  saludo := "Hola, saludo desde main..."
  fmt.Println(saludo)
  pkg2.Turron()
}
