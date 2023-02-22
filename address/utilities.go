package address

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// converte un indirizzo ip x.x.x.x in un indirizzo uint32
func getIP(IP string) (uint32, error) {

	IP = strings.TrimSpace(IP)

	addr := strings.Split(IP, ".")
	if len(addr) != 4 {
		return 0, errors.New("IP non valido")
	}

	var bin uint32
	for _, o := range addr {
		n, err := strconv.Atoi(o)
		if err != nil {
			return 0, errors.New("IP non valido")
		}

		if n > 255 || n < 0 {
			return 0, errors.New("IP non valido")
		}

		bin <<= 8
		bin += uint32(n)
	}

	return bin, nil
}

// converte una maschera in una maschera CIDR uint32
//
//	x  - maschera CIDR della rete
//	#y - numro massimo di Ip nella rete
//	-x - taglia della rete
func GetMask(Mask string) (uint32, error) {
	var mask uint32

	Mask = strings.TrimSpace(Mask)

	if Mask[0] == '#' {
		dim, err := strconv.Atoi(Mask[1:])
		if err != nil || dim < 0 {
			return 0, errors.New("dimensione non valida")
		}

		mask = 32 - getMinDim(uint32(dim))

	} else {
		dim, err := strconv.Atoi(Mask)

		if dim < 0 && dim > -32 {
			dim = 32 + dim
		}

		if err != nil || dim < 0 || dim > 30 {
			return 0, errors.New("maschera non valida")
		}

		mask = uint32(dim)
	}
	return mask, nil
}

// coverte un indirizzo uint32 in norazione puntata x.x.x.x
func toString(addr uint32) string {
	var add [4]uint8
	add[0] = uint8(addr)
	add[1] = uint8(addr >> 8)
	add[2] = uint8(addr >> 16)
	add[3] = uint8(addr >> 24)

	return fmt.Sprintf("%d.%d.%d.%d", add[3], add[2], add[1], add[0])
}

// converte una maschera CIDR in una maschera di bit
func getSbnetMask(n uint32) uint32 {
	var mask uint32
	mask--
	mask <<= (32 - n)
	return mask
}

// calcola la dimansione minima della rete contenente il numero n di indirizzi
func getMinDim(n uint32) uint32 {
	var dim uint32 = 0
	n--

	for {
		n = n >> 1
		dim++

		if n <= 0 {
			break
		}
	}

	return dim
}
