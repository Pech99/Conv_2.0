package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Pech99/conv_2.0/address"
	"github.com/Pech99/conv_2.0/pool"
)

// complile: go build -ldflags "-H windowsgui"

func main() {

	if len(os.Args) < 2 {
		fmt.Println(help() + "\nInserire almeno l'indirizzo base nella forma A.B.C.D/M (es. 80.101.99.104/18)")
		return
	}

	if strings.Count(os.Args[1], "help") > 0 {
		fmt.Println(help())
		return
	}

	add, err := address.New(os.Args[1])
	if err != nil {
		fmt.Println(help(), "\n", err)
		return
	}
	fmt.Print("Indirizzo Base della Rete (", os.Args[1], ")", add.Info(), "\n\n")

	if len(os.Args) <= 2 {
		return
	}

	pool := pool.New(add)
	for _, e := range os.Args[2:] {

		mask, err := address.GetMask(e)
		if err != nil {
			fmt.Print(e, "\t--> A.B.C.D/M\n", err, "\n\n")
			continue
		} else if mask == 0 {
			continue
		}

		if pool.IsEmpty() {
			fmt.Print("Fine dello Spazio\n\n")
			break
		}

		add, pool = pool.Get(mask)
		if add == 0 {
			fmt.Print(e, "\t--> A.B.C.D/", mask, "\nSpazio insufficiente per questa rete\n\n")
			continue
		}

		fmt.Print(e, "\t--> A.B.C.D/", mask, add.Info(), "\n\n")
	}

	if !pool.IsEmpty() {
		fmt.Println("Spazzi Rimanenti:\n" + pool.ToString())
	}
}

func help() string {
	return "Restituisce informazioni riguardo ad un indirizzo di rete ed opera il subnetting.\n\n" +
		"conv A.B.C.D/M [M | -M | #M]...\n\n" +
		"A.B.C.D/M\tIndirizzo base della rete (o un indirizzo contenuto in essa) nel formato CIDR. (es. 80.101.99.104/18)\n" +
		"M\t\tNumero di bit dedicati all'netID.\n" +
		"-M\t\tNumero di bit dedicati all'hostID.\n" +
		"#M\t\tNumero di indirizzi minimo che la rete deve contenere (compreso il Getaway).\n"
}
