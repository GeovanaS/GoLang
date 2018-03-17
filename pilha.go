package main 

import (
	"fmt"
)

type Pilha struct{
	valores []interface{}
}

func (pilha Pilha) TamanhoPilha() int {
	return len(pilha.valores)
}

func (pilha Pilha) Empty() bool{
	return pilha.TamanhoPilha() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}){
	pilha.valores = append(pilha.valores, valor)
}

func(pilha *Pilha) Desempilhar() (interface{}) {
    if pilha.Empty() {
    	return nil
    }
    valor := pilha.valores[pilha.TamanhoPilha()-1]
    pilha.valores = pilha.valores[:pilha.TamanhoPilha()-1]
    return valor
}


func main() {
    
    pilha := Pilha{}

    fmt.Println("Is Empty? ", pilha.Empty())

    pilha.Empilhar(150)
    pilha.Empilhar("Lola")
    pilha.Empilhar(23.03)

    fmt.Println("Tamanho da pilha ", pilha.TamanhoPilha())


    for !pilha.Empty() {
    	i := pilha.Desempilhar()
    	fmt.Println("Desempilhando ", i)
    	fmt.Println("Is Empty? ", pilha.Empty())
    }

}