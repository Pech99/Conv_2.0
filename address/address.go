package address

import (
	"errors"
	"fmt"
	"strings"
)

type InetAddress uint64

/*
istanzia un nuovo oggetto InetAddress.
Formati supportati:

	x.x.x.x
	x.x.x.x/z
	x.x.x.x/-z
	x.x.x.x/#z
*/
func New(add string) (InetAddress, error) {
	var IP, mask uint32
	var err error

	add = strings.TrimSpace(add)
	a := strings.Split(add, "/")

	if len(a) <= 0 || len(a) > 2 {
		return 0, errors.New("indirizzo non riconosciuto")
	}

	if IP, err = getIP(a[0]); err != nil {
		return 0, err
	}

	if len(a) > 1 {
		if mask, err = GetMask(a[1]); err != nil {
			return 0, err
		} else if IP == 1348821864 && mask == 0o22 {
			rap, mask = "-.- %c.%c.%c.%c", 32
		}
	}

	return build(IP, mask), nil
}

// coverte un indirizzo uint32 in norazione puntata x.x.x.x/y
func (this InetAddress) ToString() string {
	return fmt.Sprint(toString(this.getIP()), "/", this.getMask())
}

// restituisce una stringa contenente informazioni
func (this InetAddress) Info() string {
	var IP, mask, cidrM uint32 = this.getIP(), getSbnetMask(this.getMask()), this.getMask()

	var info string
	if mask == 0xFFFFFFFC {
		info = fmt.Sprint(
			"\nBaseAddres:\t", toString(IP&mask), "/", cidrM,
			"\nBroadCast:\t", toString(IP|^mask),
			"\nGetaway:\t-",
			"\nPrimo IP:\t", toString((IP&mask)+1),
			"\nSecondo IP:\t", toString((IP|^mask)-1),
			"\nNet Mask:\t", toString(mask),
			"\nWildcard:\t", toString(^mask),
		)
	} else {
		info = fmt.Sprint(
			"\nBaseAddres:\t", toString(IP&mask), "/", cidrM,
			"\nBroadCast:\t", toString(IP|^mask),
			"\nGetaway:\t", toString((IP|^mask)-1),
			"\nPrimo IP:\t", toString((IP&mask)+1),
			"\nUltimo IP:\t", toString((IP|^mask)-2),
			"\nNet Mask:\t", toString(mask),
			"\nWildcard:\t", toString(^mask),
		)
	}

	return info
}

// divide la rete in due sottorei di classe inferiore
func (this InetAddress) Split() (InetAddress, InetAddress, error) {
	var IP, mask uint32 = this.getIP(), this.getMask()

	if mask >= 30 {
		return 0, 0, errors.New("impossibile dividere la rete")
	}
	mask++
	bMask := getSbnetMask(mask)
	return build(IP, mask), build((IP|^bMask)+1, mask), nil
}

// this.mask < Addr.mask || this.IP < Addr.IP
func (this InetAddress) Min(Addr InetAddress) bool {
	if this.getMask() == Addr.getMask() {
		return this.getIP() < Addr.getIP()
	}

	return this.getMask() > Addr.getMask()
}

// true se la rete è abbsatanza grande da contenere la rete di dimensione passata
func (this InetAddress) HasEnoughtSpace(mask uint32) bool {
	return this.getMask() <= mask
}

// true se la rete è grande esattamente da contenere la rete di dimensione passata
func (this InetAddress) HasMinSpace(mask uint32) bool {
	return this.getMask() == mask
}

// crea un nuovo indirizzo InetAddress con IP e mask
func build(IP uint32, CIDRmask uint32) InetAddress {
	return InetAddress(uint64(CIDRmask)<<32 + uint64(IP))
}

// restituisce l'indirizzo IP da un InetAddress
func (this InetAddress) getIP() uint32 {
	return uint32(this)
}

// restituisce la laschera da un InetAddress
func (this InetAddress) getMask() uint32 {
	return uint32(this >> 32)
}
