package main

import (
	"exercicios/golang_banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	// conta1 := contas.ContaCorrente{Titular: "Marcio", NumeroAgencia: 589, NumeroConta: 123456, saldo: 2000}
	// fmt.Println(conta1.saldo)

	// fmt.Println(conta1.Sacar(300))
	// fmt.Println(conta1.saldo)

	// status, valor := conta1.Depositar(300)
	// fmt.Println("status", status)
	// fmt.Println("valor", valor)
	// fmt.Println(conta1.saldo)

	// conta2 := ContaCorrente{"Alex", 581, 919191, 500}
	// fmt.Println(conta2)

	// var conta3 *ContaCorrente
	// conta3 = new(ContaCorrente)
	// conta3.titular = "Joao"
	// conta3.saldo = 500
	// fmt.Println(*conta3)

	// // TRANSFERENCIA ENTRE CONTAS:
	// contaDoJoao := contas.ContaCorrente{Titular: "Joao", saldo: 1000}
	// contaDaMaria := contas.ContaCorrente{Titular: "Maria", saldo: 2000}

	// statusTransf := contaDoJoao.Transferir(800, &contaDaMaria)
	// //statusTransf := contaDaMaria.Transferir(-2800, &contaDoJoao)
	// fmt.Println(statusTransf)
	// fmt.Println(contaDoJoao)
	// fmt.Println(contaDaMaria)

	// // COMPOSIÇÃO DE STRUCT
	// clienteMarcio := clientes.Titular{
	// 	Nome:      "Marcio",
	// 	Cpf:       "187261876728",
	// 	Profissao: "Analista de sistemas"}

	// contaDoMarcio := contas.ContaCorrente{
	// 	Titular:       clienteMarcio,
	// 	NumeroAgencia: 7711,
	// 	NumeroConta:   66236,
	// 	saldo:         100.0}

	// fmt.Println(contaDoMarcio)

	contaCorr := contas.ContaCorrente{}
	contaCorr.Depositar(100)

	// fmt.Println(contaTeste)

	// fmt.Println(contaTeste.ObterSaldo())

	contaPoup1 := contas.ContaPoupanca{}
	contaPoup1.Depositar(200)
	fmt.Println(contaPoup1)
	fmt.Println(contaPoup1.ObterSaldo())

	contaPoup1.Sacar(55)
	fmt.Println(contaPoup1.ObterSaldo())

	PagarBoleto(&contaCorr, 20)
	fmt.Println(contaCorr.ObterSaldo())

	PagarBoleto(&contaPoup1, 30)
	fmt.Println(contaPoup1.ObterSaldo())

}
