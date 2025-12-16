package module

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Symbole struct {
	Symb   string
	Freq   int
	Codage string
}
type Noeud struct {
	Symbole     Symbole
	NoeudGauche *Noeud
	NoeudDroite *Noeud
}

func Huffman(mot string) {
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
		nbr += 2
		usedNode = append(usedNode, arbreHuffman[len(arbreHuffman)-2])
		usedNode = append(usedNode, arbreHuffman[len(arbreHuffman)-1])

		newNode := Noeud{
			Symbole: Symbole{
				Freq: arbreHuffman[len(arbreHuffman)-1].Symbole.Freq + arbreHuffman[len(arbreHuffman)-2].Symbole.Freq,
			},
			NoeudGauche: &usedNode[nbr-2],
			NoeudDroite: &usedNode[nbr-1],
		}

		arbreHuffman = arbreHuffman[:len(arbreHuffman)-2]
		arbreHuffman = append(arbreHuffman, newNode)
		sort.Slice(arbreHuffman, func(i, j int) bool {
			return arbreHuffman[i].Symbole.Freq > arbreHuffman[j].Symbole.Freq
		})

	}
	// fmt.Println(usedNode)

	fmt.Println(arbreHuffman)
	racine := arbreHuffman[0]
	usedNode = append(usedNode, racine)
	for _, element := range usedNode{
		fmt.Println(element)
	}
	
	if usedNode[len(usedNode)-1].NoeudGauche != nil  {
		getnodegauche(*usedNode[len(usedNode)-1].NoeudGauche)
		getnodedroite(*usedNode[len(usedNode)-1].NoeudDroite)
	}

}
func getnodegauche(n Noeud){
	if n.NoeudGauche != nil {
		fmt.Println(n.NoeudGauche.Symbole.Symb)
		getnodegauche(*n.NoeudGauche)
		getnodedroite(*n.NoeudDroite)
	}
}
func getnodedroite(n Noeud){
	if n.NoeudDroite != nil {
		fmt.Println(n.NoeudDroite.Symbole.Symb)
		getnodegauche(*n.NoeudGauche)
		getnodedroite(*n.NoeudDroite)
	}
	
}
