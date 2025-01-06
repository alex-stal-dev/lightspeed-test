package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "./sample-ips"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var uniqueCounter uint32 = 0
	ipBitSet := make([]bool, math.MaxUint32)

	for scanner.Scan() {
		line := scanner.Text()

		ipAsInt, err := ipStringToInt(line)

		if err == nil {
			ipExists := ipBitSet[ipAsInt]
			if !ipExists {
				uniqueCounter++
				ipBitSet[ipAsInt] = true
			}
		}
	}

	fmt.Println("The number of unique addresses is ", uniqueCounter)
}

// ipStringToInt represents IPv4-address string as a 4-byte unsigned integer since
// IPv4 has 4 sections with 1-byte value each.
func ipStringToInt(ipStr string) (uint32, error) {
	tokens := strings.Split(ipStr, ".")
	if len(tokens) != 4 {
		return 0, errors.New("invalid ip string")
	}

	var fourByteIP uint32
	for _, token := range tokens {
		intValue, err := strconv.Atoi(token)
		if err != nil || intValue < 0 || intValue > 255 {
			return 0, errors.New("invalid ip string")
		}
		fourByteIP = (fourByteIP << 8) | uint32(intValue)
	}

	return fourByteIP, nil
}
