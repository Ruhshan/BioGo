package biogo

import (
	"io/ioutil"
	"strings"
)

type cre_feature struct {
	ID     string
	Values []string
}

func getCres() []string {

	cres := []string{}

	dat, _ := ioutil.ReadFile("place.dat.txt")
	data_string := string(dat)
	splitted := strings.Split(data_string, "\n")

	for _, line := range splitted {
		splited_line := strings.Fields(line)
		//fmt.Println(splited_line[2])

		cres = append(cres, splited_line[2])
	}

	return cres
}

func MakeCreFeatures(fastas []*fasta) []*cre_feature {
	cre_features := []*cre_feature{}

	cres := getCres()

	for _, fasta := range fastas {
		new_cre_feature := new(cre_feature)
		new_cre_feature.ID = fasta.Id

		for _, cre := range cres {
			iud := IUPACDna{Sequence: cre}
			is_matched := iud.Match(fasta.Sequence)
			new_cre_feature.Values = append(new_cre_feature.Values, is_matched)
		}

		cre_features = append(cre_features, new_cre_feature)

	}

	return cre_features

}
