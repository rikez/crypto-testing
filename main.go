package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Bitstream struct {
	arr    []int
	bitmap map[int]int
}

func generateBits() *Bitstream {
	bitstream := make([]int, 0)
	bitmap := make(map[int]int, 0)

	for i := 0; i < 20000; i++ {
		numb := rand.Intn(2)
		bitstream = append(bitstream, numb)

		if _, ok := bitmap[numb]; ok {
			bitmap[numb]++
		} else {
			bitmap[numb] = 1
		}
	}

	return &Bitstream{
		arr:    bitstream,
		bitmap: bitmap,
	}
}

func convertToBitstream(byteArr []byte) *Bitstream {
	bitstream := make([]int, 0)
	bitmap := make(map[int]int, 0)

	for _, b := range byteArr {
		str := hexToBin(b)

		for i := 0; i < len(str); i++ {
			integer, err := strconv.Atoi(string(str[i]))
			if err != nil {
				panic(err)
			}
			bitstream = append(bitstream, integer)

			if _, ok := bitmap[integer]; ok {
				bitmap[integer]++
			} else {
				bitmap[integer] = 1
			}
		}
	}

	return &Bitstream{
		arr:    bitstream,
		bitmap: bitmap,
	}
}

func monobit(bitstream *Bitstream) bool {
	if bitstream.bitmap[1] > 9654 && bitstream.bitmap[1] < 10346 {
		return true
	}
	return false
}

func poker(bitstream *Bitstream) bool {
	occurences := make(map[int]float64)
	for i := 0; i < len(bitstream.arr); i += 4 {
		nibble := bitstream.arr[i]*1000 + bitstream.arr[i+1]*100 + bitstream.arr[i+2]*10 + bitstream.arr[i+3]
		// nibble := fmt.Sprintf("%d%d%d%d", bitstream.arr[i], bitstream.arr[i+1], bitstream.arr[i+2], bitstream.arr[i+3])

		if _, ok := occurences[nibble]; ok {
			occurences[nibble]++
		} else {
			occurences[nibble] = 1
		}
	}

	sum := 0.0

	for _, v := range occurences {
		sum += v * v
	}

	result := 16.0/5000.0*sum - 5000.0

	if result > 1.03 && result < 57.4 {
		return true
	}

	return false
}

func runs(bitstream *Bitstream) bool {
	lengths := make(map[int]int)
	prev := -1
	counter := 1

	for i := 0; i < len(bitstream.arr); i++ {
		curr := bitstream.arr[i]

		if curr == prev {
			if counter < 6 {
				counter++
			}
		} else {
			if _, ok := lengths[counter]; ok {
				lengths[counter]++
			} else {
				lengths[counter] = 1
			}

			counter = 0
			prev = curr
		}
	}

	for k, v := range lengths {
		if !validateInterval(k, v) {
			return false
		}
	}

	return true
}

func validateInterval(k int, v int) bool {
	if k == 1 {
		return v >= 2267 && v <= 2733
	}

	if k == 2 {
		return v >= 1079 && v <= 1421
	}

	if k == 3 {
		return v >= 502 && v <= 748
	}

	if k == 4 {
		return v >= 223 && v <= 402
	}

	if k >= 5 {
		return v >= 90 && v <= 223
	}

	return false
}

func longruns(bitstream *Bitstream) bool {
	lengths := make(map[int]int)
	prev := -1
	counter := 0

	for i := 0; i < len(bitstream.arr); i++ {
		curr := bitstream.arr[i]

		if curr == prev {
			counter++
		} else {
			if _, ok := lengths[counter]; ok {
				lengths[counter]++
			} else {
				lengths[counter] = 1
			}

			counter = 0
			prev = curr
		}
	}

	if _, ok := lengths[34]; ok {
		return false
	}

	return true
}

func hexToBin(in byte) string {
	var out []byte
	for i := 3; i >= 0; i-- {
		b := (in >> uint(i))
		out = append(out, (b%2)+48)
	}

	return string(out)
}

func main() {
	// bitstream := generateBits()
	// fmt.Printf("Monobit: %v\n", monobit(bitstream))
	// fmt.Printf("Poker: %v\n", poker(bitstream))
	// fmt.Printf("Runs: %v\n", runs(bitstream))
	// fmt.Printf("Longrun: %v\n", longruns(bitstream))

	f, _ := os.Open("./keys")
	byteValue, _ := ioutil.ReadAll(f)

	keys := strings.Split(string(byteValue), "\n")

	for i := 0; i < len(keys); i++ {
		byteArr := []byte(keys[i])

		bitstream := convertToBitstream(byteArr)

		fmt.Printf("Running for key #%d\n", i+1)
		fmt.Printf("Monobit: %v\n", monobit(bitstream))
		fmt.Printf("Poker: %v\n", poker(bitstream))
		fmt.Printf("Runs: %v\n", runs(bitstream))
		fmt.Printf("Longrun: %v\n", longruns(bitstream))
		fmt.Println()
		fmt.Println()
	}
}
