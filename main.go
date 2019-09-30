package main

import (
	"fmt"
	"math/rand"
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
	counter := 0

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
			if counter < 34 {
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

	if _, ok := lengths[34]; ok {
		return false
	}

	return true
}

func main() {
	bitstream := generateBits()

	fmt.Printf("Monobit: %v\n", monobit(bitstream))
	fmt.Printf("Poker: %v\n", poker(bitstream))
	fmt.Printf("Runs: %v\n", runs(bitstream))
	fmt.Printf("Longrun: %v\n", longruns(bitstream))
}
