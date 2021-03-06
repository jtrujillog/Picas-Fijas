package main

import (
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
)

func main(){
  port := os.Getenv("PORT")

  if port == "" {
    port = "8081"
  }

  router := gin.Default()

  router.GET("/codebreaker/setup/", func(c *gin.Context){
    c.String(http.StatusOK, "Ingresa un número de 4 dígitos despues de setup/")
  })

  router.GET("/codebreaker/setup/:numero", func(c *gin.Context){
    numero := c.Param("numero")
    respuesta := validarDigitos(numero)
    /*si el tamaño del número es diferente de cuatro*/
    if len(numero) == 5{//Se coloca 5 debido a los (:) del enlace
      c.String(http.StatusOK, "El número debe ser de 4 dígitos, vuelve a ingresarlo")
    }else{
      /*si los dígitos del numero son iguales*/
      if respuesta == 0{
        numeroSecreto(numero)
        c.String(http.StatusOK, "El numero secreto es: " + numero + ". Ahora ingresa a /codebreaker/guess/ un número de 4 dígitos para comenzar a adivinar el número secreto")
      }else{
        c.String(http.StatusOK, "Los digitos del número no deben ser iguales entre si")
      }
    }

  })

  router.GET("/codebreaker/guess/:numero", func(c *gin.Context){
    numero := c.Param("numero")
    resultado := validarCodigo(numero)
    c.String(http.StatusOK, "Pista: " + resultado)
  })

  router.Run(":" +port)

}
