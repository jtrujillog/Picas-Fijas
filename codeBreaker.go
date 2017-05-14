package main

import (
  "strconv"
  "strings"
)

var private_code string
var contador = 0

//para guardar el codigo secreto
func numeroSecreto(code string){
  private_code = code
}

//función para validar el número nuevo
func validarCodigo(code string)string{
  intentos := contadorIntentos()
  code_array := strings.Split(code, "")
  private_code_array := strings.Split(private_code, "")
  respuesta := ""
  size := len(code)

  if size != 4{
    respuesta += "Debes ingresar un número de 4 dígitos"
    return respuesta
  }

  digitos := validarDigitos(code)
  if digitos == 1{
    respuesta += "Los digitos del número no deben ser iguales entre si, vuelve a intetarlo"
    return respuesta
  }

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


//valida la cantidad de intentos
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

//función para validar los dígitos del número y que estos no sean repetidos
func validarDigitos(numero string) int{
  respuesta := 0
  numero_array := strings.Split(numero, "")
  size := len(numero)
  for i := 0; i < size; i++ {
    for j:= i+1; j < size; j++{
      if numero_array[i] == numero_array[j] {
        respuesta = 1
      }
    }
  }
  return respuesta
}
