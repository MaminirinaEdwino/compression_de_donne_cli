package module

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Symbole struct {
	Symb   string
	Freq   int
	Codage string
}
type Noeud struct {
	Symbole      Symbole
	NoeudSuivant string
	Cote         int
}

func Huffman(mot string) string {
	var symb []string
	tab := strings.Split(mot, "")
	var freqTab []Symbole
	// intab := false
	// fmt.Println(len(tab))
	for _, symbole := range tab {
		if !slices.Contains(symb, symbole) {
			symb = append(symb, symbole)
		}
	}
	for _, symbole := range symb {
		freqTab = append(freqTab, Symbole{
			Symb: symbole,
			Freq: 0,
		})
	}

	for i, symbole := range freqTab {
		for _, s := range tab {
			// fmt.Println(symbole.Symb, s)
			if symbole.Symb == s {
				freqTab[i].Freq++
			}
		}
	}
	sort.Slice(freqTab, func(i, j int) bool {
		return freqTab[i].Freq > freqTab[j].Freq
	})
	// fmt.Println(freqTab)
	var arbreHuffman []Noeud
	var originalNode []Noeud
	var usedNode []Noeud

	for _, element := range freqTab {
		originalNode = append(originalNode, Noeud{
			Symbole: element,
		})
	}
	arbreHuffman = originalNode
	nbr := 0
	for len(arbreHuffman) > 1 {
		nbr += 1

		newNode := Noeud{
			Symbole: Symbole{
				Symb: strconv.Itoa(nbr),
				Freq: arbreHuffman[len(arbreHuffman)-1].Symbole.Freq + arbreHuffman[len(arbreHuffman)-2].Symbole.Freq,
			},
		}
		arbreHuffman[len(arbreHuffman)-2].NoeudSuivant = strconv.Itoa(nbr)
		arbreHuffman[len(arbreHuffman)-2].Cote = 0
		arbreHuffman[len(arbreHuffman)-1].NoeudSuivant = strconv.Itoa(nbr)
		arbreHuffman[len(arbreHuffman)-1].Cote = 1
		usedNode = append(usedNode, arbreHuffman[len(arbreHuffman)-2])
		usedNode = append(usedNode, arbreHuffman[len(arbreHuffman)-1])

		arbreHuffman = arbreHuffman[:len(arbreHuffman)-2]
		arbreHuffman = append(arbreHuffman, newNode)
		sort.Slice(arbreHuffman, func(i, j int) bool {
			return arbreHuffman[i].Symbole.Freq > arbreHuffman[j].Symbole.Freq
		})

	}
	// fmt.Println(usedNode)

	// fmt.Println(arbreHuffman)
	racine := arbreHuffman[0]
	usedNode = append(usedNode, racine)

	// GetCodage(usedNode[2])
	// fmt.Print(usedNode[1].Symbole.Symb)

	var f = make(map[string]string)
	for _, n := range usedNode {
		if _, err := strconv.Atoi(n.Symbole.Symb); err != nil {
			codage := ""
			GetCodage(n, usedNode, &codage)
			// fmt.Println("inv", codage)
			realcode := ""
			for i := len(codage); i > 0; i-- {
				realcode += codage[i-1 : i]
			}
			// fmt.Println(n.Symbole.Symb, realcode)
			f[n.Symbole.Symb] = realcode
		}
	}

	// file, err := os.OpenFile("compression.tay", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	// if err != nil{
	// 	panic(err)
	// }

	data := ""

	for _, i := range tab {
		// fmt.Println(f[i])
		data += f[i]

	}

	// fmt.Fprint(file, ".")
	// for idx, i := range f {
	// 	fmt.Fprintf(file, "%s%s",idx, i)
	// }
	fmt.Println(data)
	return data
}
func GetCodage(element Noeud, usedNode []Noeud, codage *string) int {
	if element.NoeudSuivant != "" {
		// fmt.Println(element)
		var suivant Noeud
		for _, i := range usedNode {
			if element.NoeudSuivant == i.Symbole.Symb {
				// fmt.Print(element.Cote)
				*codage += fmt.Sprintf("%d", element.Cote)
				suivant = i
				break
			}
		}
		GetCodage(suivant, usedNode, codage)
	}
	return element.Cote
}
