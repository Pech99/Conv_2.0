package pool

import (
	"sort"

	"github.com/Pech99/conv/address"
)

type AddrPool []address.InetAddress

var isSorted bool = true

// crea un nuovo pull di InetAddress tramite un indirizzo base
func New(BaseAddr address.InetAddress) AddrPool {
	var pool AddrPool = make([]address.InetAddress, 0)
	pool = append(pool, BaseAddr)
	isSorted = true
	return pool
}

// restituisce il primo pool libero grande quanto la CIDR mask richiesta
func (this AddrPool) Get(mask uint32) (address.InetAddress, AddrPool) {
	var i int
	var Addr address.InetAddress

	if this.IsEmpty() {
		return 0, this
	}

	this.sort()

	for i = 0; i < len(this) && !this[i].HasEnoughtSpace(mask); i++ {
	}

	if !(i < len(this)) {
		return 0, this
	}

	Addr = this[i]
	this = this.remove(i)

	if Addr.HasMinSpace(mask) {
		return Addr, this
	}

	return this.split(Addr, mask)
}

// rompe una rete più grande in sottoreti più piccole
func (this AddrPool) split(Addr address.InetAddress, mask uint32) (address.InetAddress, AddrPool) {

	if Addr.HasMinSpace(mask) {
		return Addr, this
	}

	Addr1, Addr2, _ := Addr.Split()
	this = this.add(Addr2)
	return this.split(Addr1, mask)
}

// true se non sono più disponibili IP
func (this AddrPool) IsEmpty() bool {
	return len(this) == 0
}

// aggiunge un indirizzo al pool
func (this AddrPool) add(Addr address.InetAddress) AddrPool {
	isSorted = false
	return append(this, Addr)
}

// rimuove un indirizzo dal pool dato l'indice
func (this AddrPool) remove(i int) AddrPool {
	return append(this[:i], this[i+1:]...)
}

// ordina il pool di indirizzi, dal più piccolo al più grande
func (this AddrPool) sort() {
	if !isSorted {
		less := func(a, b int) bool {
			return this[a].Min(this[b])
		}
		sort.Slice(this, less)
	}
	isSorted = true
}
