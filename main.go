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
  pkg2.Turron()
}
