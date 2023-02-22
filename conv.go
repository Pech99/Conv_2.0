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

	if len(os.Args) != 2 {
		return
	}

	c, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	s := strings.TrimSpace(string(c))
	cnt := strings.Split(string(s), "\n")

	add, err := address.New(cnt[0])
	if err != nil {
		write(fmt.Sprint(err, "\n\n"))
		return
	}
	out += add.Info() + "\n\n"

	pool := pool.New(add)

	for _, e := range cnt[1:] {

		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}

		mask, err := address.GetMask(e)
		if err != nil {
			out += fmt.Sprint(err, "\n\n")
			continue
		}

		if pool.IsEmpty() {
			out += "fine dello spazio\n\n"
			break
		}

		add, pool = pool.Get(mask)
		inf := add.Info()
		out += fmt.Sprint(e, "\t--> x.x.x.x/", mask, inf, "\n\n")
	}

	write(out)

}

func write(out string) error {
	f, err := os.Create("a.txt")
	if err != nil {
		return err
	}
	f.Write([]byte(out))
	return nil
}
