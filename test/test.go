package main

import "fmt"

func main() {
	var n uint32
	var i int
	for n, i = 1, 1; n != 0; i++ {
		n = n << 1
		fmt.Print(i, "\t-\t", n, "\n")
	}

	fmt.Print("1\t-\t", getMinDim(1), "\n")
	fmt.Print("2\t-\t", getMinDim(2), "\n")
	fmt.Print("4\t-\t", getMinDim(4), "\n")
	fmt.Print("8\t-\t", getMinDim(8), "\n")
	fmt.Print("16\t-\t", getMinDim(16), "\n")
	fmt.Print("32\t-\t", getMinDim(32), "\n")
	fmt.Print("64\t-\t", getMinDim(64), "\n")
	fmt.Print("128\t-\t", getMinDim(128), "\n")
	fmt.Print("256\t-\t", getMinDim(256), "\n")

	fmt.Print("3\t-\t", getMinDim(3), "\n")

	fmt.Print("10\t-\t", getMinDim(10), "\n")
	fmt.Print("20\t-\t", getMinDim(20), "\n")
	fmt.Print("30\t-\t", getMinDim(30), "\n")
	fmt.Print("100\t-\t", getMinDim(100), "\n")
	fmt.Print("500\t-\t", getMinDim(500), "\n")
	fmt.Print("1000\t-\t", getMinDim(1000), "\n")

	return

}

func getMinDim(n int) int {
	dim := 0
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

/*
1       -       2
2       -       4
3       -       8
4       -       16
5       -       32
6       -       64
7       -       128
8       -       256
9       -       512
10      -       1024
11      -       2048
12      -       4096
13      -       8192
14      -       16384
15      -       32768
16      -       65536
17      -       131072
18      -       262144
19      -       524288
20      -       1048576
21      -       2097152
22      -       4194304
23      -       8388608
24      -       16777216
25      -       33554432
26      -       67108864
27      -       134217728
28      -       268435456
29      -       536870912
30      -       1073741824
31      -       2147483648
32      -       0
*/
