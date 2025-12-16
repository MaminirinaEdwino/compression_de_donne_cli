package module

import (
	"fmt"
	"strings"
)

func BWT(mot string) {
	tab := strings.Split(mot, "")
	var result []string
	var final string
	var idFinal int
	for id := range tab {
		results := fmt.Sprintf("%s%s", strings.Join(tab[id:], ""), strings.Join(tab[:id], ""))
		result = append(result, results)
	}

	for id := range result {
		for id1 := range result {
			if result[id] < result[id1] {
				tmp := result[id1]
				result[id1] = result[id]
				result[id] = tmp
			}
		}
	}

	for _, element := range result {
		tmp := strings.Split(element, "")
		final += tmp[len(tmp)-1]
	}
	for id, res := range result {
		if mot == res {
			idFinal = id
		}
	}
	fmt.Println(final, idFinal+1)
}

func tri(C2 []string) []string {
	for i := range C2 {
		for j := range C2 {
			if C2[i] < C2[j] {
				tmp := C2[i]
				C2[i] = C2[j]
				C2[j] = tmp
			}
		}
	}
	return C2
}

func addChar(tab []string, C1 []string) []string {
	for i := range tab {
		tab[i] = fmt.Sprintf("%s%s", C1[i], tab[i])
	}
	return tab
}

func DecodeBWT(mot string, position int) {
	var C1 []string
	var C2 []string
	// var C3 []string

	C1 = strings.Split(mot, "")
	C2 = append(C2, C1...)

	C2 = tri(C2)
	for len(C2[len(C2)-1]) < len(C1) {
		C2 = addChar(C2, C1)
		C2 = tri(C2)
	}

	fmt.Println(C2[position -1])
}
