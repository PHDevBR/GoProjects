package main

import (
    "fmt"
    "net/http"
)

type Block struct {
    Try     func()
    Catch   func(Exception)
    Finally func()
}
 
type Exception interface{}
 
func Throw(up Exception) {
    panic(up)
}
 
func (tcf Block) Do() {
    if tcf.Finally != nil {
 
        defer tcf.Finally()
    }
    if tcf.Catch != nil {
        defer func() {
            if r := recover(); r != nil {
                tcf.Catch(r)
            }
        }()
    }
    tcf.Try()
}
 
func main() {
    fmt.Println("Programa iniciado!")
    Block{
        Try: func() {
            
            var site string

            fmt.Println("Insira o Site : ")
            fmt.Scanln(&site)

            fmt.Println("Site :", site)

            resp, _:= http.Get("https://" + site)
    
        if resp.StatusCode == 200 {
            fmt.Println("o Site :", site, "Se Encontra Online")
    } },
        Catch: func(e Exception) {
            //fmt.Printf("Caught %v\n", e)
            fmt.Println("O site não existe ou não se encontra online!")
        },
        Finally: func() {
            fmt.Println("Programa terminado!")
        },
    }.Do()
}