package module

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Binary(nombre int) {
	result := ""
	for nombre >= 2 {
		// fmt.Print(nombre % 2)
		result += fmt.Sprintf("%d", nombre%2)
		nombre = nombre / 2
	}
	result += fmt.Sprintf("%d", nombre%2)
	result2 := strings.Split(result, "")
	for i := len(result); i > 0; i-- {
		fmt.Print(result2[i-1])
	}
}

func DecodeBinary(code string) int {
	binary := strings.Split(code, "")
	res := 0
	ln := len(binary) - 1
	for i := range len(binary) {
		resu, err := strconv.Atoi(binary[i])
		if err != nil{
			panic(err)
		}
		fmt.Println(resu)
		res+=int(math.Pow(2, float64(ln)))*resu
		ln--
	}
	
	fmt.Println(res)
	return res
}
