package module

import (
	"fmt"
	"os"
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

func Huffman(mot string) (string, map[string]string) {
	var symb []string
	tab := strings.Split(mot, "")
	var freqTab []Symbole
	var arbreHuffman []Noeud
	var originalNode []Noeud
	var usedNode []Noeud
	var f = make(map[string]string)

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
			if symbole.Symb == s {
				freqTab[i].Freq++
			}
		}
	}
	sort.Slice(freqTab, func(i, j int) bool {
		return freqTab[i].Freq > freqTab[j].Freq
	})

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
	racine := arbreHuffman[0]
	usedNode = append(usedNode, racine)

	for _, n := range usedNode {
		if _, err := strconv.Atoi(n.Symbole.Symb); err != nil {
			codage := ""
			GetCodage(n, usedNode, &codage)
			realcode := ""
			for i := len(codage); i > 0; i-- {
				realcode += codage[i-1 : i]
			}
			f[n.Symbole.Symb] = realcode
		}
	}
	data := ""
	for _, i := range tab {
		data += f[i]
	}
	fmt.Println("data", f)
	return data, f
}
func GetCodage(element Noeud, usedNode []Noeud, codage *string) int {
	if element.NoeudSuivant != "" {
		var suivant Noeud
		for _, i := range usedNode {
			if element.NoeudSuivant == i.Symbole.Symb {
				*codage += fmt.Sprintf("%d", element.Cote)
				suivant = i
				break
			}
		}
		GetCodage(suivant, usedNode, codage)
	}
	return element.Cote
}

func DecodeHuffman(sortie string, codage string) {
	sortieFichier, err := os.ReadFile(sortie)
	if err != nil {
		panic(err)
	}
	codageFichier, err := os.ReadFile(codage)
	if err != nil {
		panic(err)
	}
	cd := strings.Split(string(codageFichier), "\n")
	
	cdg := make(map[string]string)
	for _, i := range cd {
		str := strings.Split(i, "")
		if len(str)>0 {
			cdg[str[0]] = strings.Join(str[1:], "")
		} 
	}
	fmt.Println(cdg)
	contenuFichier := ""
	for _, i := range sortieFichier{
		contenuFichier+=ToBinary(int(i))
	}
	code := strings.Split(contenuFichier, "")
	used := ""
	for _ ,i := range code{
		used+=i
		for k, cod := range cdg {
			// fmt.Println(used)
			if used == cod {
				// fmt.Println(used)
				fmt.Println(k)
				used = ""
			}
		}
	}

}
