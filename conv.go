package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Pech99/conv/address"
	"github.com/Pech99/conv/pool"
)

// complile: go build -ldflags "-H windowsgui"

func main() {

	var out string

	c, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	cnt := strings.Split(string(c), "\n")

	add, err := address.New(cnt[0])
	if err != nil {
		out += fmt.Sprint(err, "\n\n")
		return
	}
	pool := pool.New(add)

	for _, e := range cnt[1:] {
		mask, err := address.GetMask(e)
		if err != nil {
			out += fmt.Sprint(err, "\n\n")
		}

		if pool.IsEmpty() {
			out += "fine dello spazio\n\n"
			break
		}

		add, pool = pool.Get(mask)
		inf := add.Info()
		out += e + inf + "\n\n"
	}

	f, err := os.Create("a.txt")
	if err != nil {
		return
	}
	f.Write([]byte(out))

}
