package module

import (
	"fmt"
	"strconv"
	"strings"
)


func RLE(mot string){
	fmt.Println("RLE")
	var result []string
	resultString := ""
	tab := strings.Split(mot, "")
	count := -1
	for _, symbole := range tab{
		if len(result) == 0 {
			result = append(result, symbole)
			count++
		}
		if result[len(result)-1] == symbole {
			count++
		}else{
			resultString+=fmt.Sprintf("%d%s", count, result[len(result)-1])
			result = append(result, symbole)
			count= 1
		}
	}
	resultString+=fmt.Sprintf("%d%s", count, result[len(result)-1])
	fmt.Println(resultString)
}

func DecodeRLE(mot string){
	tab := strings.Split(mot, "")
	result := ""
	count := ""
	for _, e := range tab{
		if _, err := strconv.Atoi(e); err == nil {
			count+= e
		}else{
			val, err := strconv.Atoi(count)
			if err != nil {
				panic(err)
			}
			for range val {
				result+=e
			}
			count = ""
		}
	}
	fmt.Println(result)
}