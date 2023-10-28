package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s" validate:"gte=0"`
}

func main() {
	conta := Conta{Numero: 123, Saldo: 1000}

	res, err := json.Marshal(conta) // returning a byte array
	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) // returning a byte array
	if err != nil {
		println(err)
	}

	contaJson := []byte(`{"n":123,"s":1000}`)

	var conta2 Conta
	err = json.Unmarshal(contaJson, &conta2)
	if err != nil {
		println(err)
	}

	println(conta2.Numero)
	println(conta2.Saldo)

}
