package main

import (
  "strconv"
  "strings"
)

var private_code string
var contador = 0

func numeroSecreto(code string){
  private_code = code
}

func validarCodigo(code string)string{
  intentos := contadorIntentos()
  code_array := strings.Split(code, "")
  private_code_array := strings.Split(private_code, "")
  respuesta := ""
  size := len(code)

  if private_code == code{
    respuesta += "xxxx, Ganaste"
    return respuesta
  }

  for i := 0; i < size; i++ {
    if private_code_array[i] == code_array[i] {
      respuesta += "x"
    } else if strings.Contains(private_code, code_array[i]) {
      respuesta += "-"
    }
  }

  return respuesta + intentos
}

func contadorIntentos() string{
  numeroIntentos := ""
  contador += 1
  if (contador > 10) {
    numeroIntentos = "  Excedsite el maximo nuḿero de intentos, el número secreto es: " + private_code + ". Vuelve a intentarlo"
    contador = 0
    return numeroIntentos
  }
  numeroIntentos = "  Llevas " + strconv.Itoa(contador) + " de 10 intetos"
  return numeroIntentos
}

func validarDigitos(numero string)bool{
  respuesta := False
  numero_array := strings.Split(code, "")
  for i := 0; i < size; i++ {
    for j:= i+1; j < size; j++{
      if numero_array[i] == numero_array[j] {
        respuesta = True
      }
    }
  }
  return respuesta
}

func validarTamaño(code string)bool{
  respuesta := False
  size := len(code)
  if size != 4 {
    respuesta = True
  }
  return respuesta
}
