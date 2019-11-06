package main

import (
	"fmt"
	"strings"

	chinese "golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func main() {
	fmt.Print("faf\n")
	quote := `If you wish to make an apple pie from scratch, you must first invent the universe.`
	cipher := IterativeRotationCipher{}
	result := cipher.encode(10, quote)
	fmt.Println(result)
	result1 := cipher.decode(10, result)
	fmt.Println(result1)
}

func isCircularSorted0(r []int) bool {
	if len(r) < 3 {
		return true
	}

	breakPoint := 0

	for index := 1; index < len(r); index++ {
		if r[index] < r[index-1] {
			breakPoint = index
			break
		}
	}

	if breakPoint == 0 {
		return true
	}

	if r[len(r)] > r[0] {
		return false
	}

	if breakPoint == len(r)-1 {
		return true
	}

	for index := breakPoint + 1; index < len(r); index++ {
		if r[index] < r[index-1] {
			return false
		}
	}
	return true
}

func isCircularSorted(r []int) bool {
	dipSeen := false
	for i := 0; i < len(r); i++ {
		if r[i] < r[i-1] {
			if !dipSeen {
				if r[len(r)-1] > r[0] {
					return false
				}
				dipSeen = true
			} else {
				return false
			}
		}
	}
	return true
}

func UGConverter(from []byte, to Charset) []byte {
	var decodedBytes []byte
	switch to {
	case GB18030:
		decodedBytes, _ = chinese.GB18030.NewEncoder().Bytes(from)
	case UTF8:
		decodedBytes, _ = chinese.GB18030.NewDecoder().Bytes(from)
	}
	return decodedBytes
}

func InArray(array1 []string, array2 []string) []string {
	var result = []string{}

	if len(array2) == 0 {
		return result
	}

	for _, value := range array1 {
		for _, innerValue := range array2 {
			if strings.Contains(innerValue, value) {
				result = append(result, value)
				break
			}
		}
	}

	var filterted []string
	if len(result) > 2 {
		filterted = append(filterted, result[0])
	} else {
		return result
	}
	for i := 1; i < len(result); i++ {
		if result[i] != filterted[len(filterted)] {
			filterted = append(filterted, result[i])
		}
	}
	return result
}

func CountBits(in uint) int {
	result := 0

	for in != 0 {
		result += int(in & 1)
		in = in >> 1
	}

	return result
	// bits := 0
	// for b > 0 {
	//   b = b & (b - 1)
	//   bits++
	// }
	// return bits
}

type IterativeRotationCipher struct{}

func (i *IterativeRotationCipher) encode(times int, quote string) string {
	result := quote

	for i := 0; i < times; i++ {
		tempStr, tempSpaceIndex := dropSpace(result)
		tempStr = moveRight(tempStr, times)
		tempStr = insertSpace(tempStr, tempSpaceIndex)
		strSlice := strings.Split(tempStr, " ")
		for index, j := range strSlice {
			strSlice[index] = moveRight(j, times)
		}
		tempStr = strings.Join(strSlice, " ")
		result = tempStr
	}

	return "$times " + result
}

func (i *IterativeRotationCipher) decode(times int, quote string) string {
	result := quote[len("${times} "):]
	for i := 0; i < times; i++ {
		strSlice, tempSpaceIndex := splitBySpace(result)
		for index, j := range strSlice {
			strSlice[index] = moveLeft(j, times)
		}
		tempStr := strings.Join(strSlice,"")
		tempStr = moveLeft(tempStr, times)
		result = insertSpace(tempStr, tempSpaceIndex)
	}
	return result
}

func moveRight(str string, i int) string {
	strLen := len(str)

	i = i % strLen
	if i == 0 {
		return str
	} else {
		return str[strLen-i:] + str[:strLen-i]
	}
}
func moveLeft(str string, i int) string {
	strLen := len(str)
	i = i % strLen
	if i == 0 {
		return str
	} else {
		return str[i:] + str[:i]
	}
}

func dropSpace(str string) (res string, spaceIndex []int) {
	for index, ele := range str {
		if ele == ' ' {
			spaceIndex = append(spaceIndex, index)
		}
	}
	res = strings.ReplaceAll(str, " ", "")
	return
}

func splitBySpace(str string) (res []string, spaceIndex []int) {
	for index, ele := range str {
		if ele == ' ' {
			spaceIndex = append(spaceIndex, index)
		}
	}
	res = strings.Split(str, " ")
	return
}

func insertSpace(str string, spaceIndex []int) string {
	for _, i := range spaceIndex {
		str = str[:i] + " " + str[i:]
	}
	return str
}
