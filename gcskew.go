package biogo

import (
	"errors"
	"fmt"
)

type gcskew struct {
	ID     string
	Values []string
}

func calcGcskew(sequence string) string {
	a, t, g, c := 1.0, 1.0, 1.0, 1.0

	for _, s := range sequence {

		switch s {
		case 65:
			a++
		case 84:
			t++
		case 67:
			c++
		case 71:
			g++

		}
	}

	//return string((g - c) / (g + c))
	return fmt.Sprintf("%.6f", (g-c)/(g+c))
}

func GCSkewSlidingWindow(fastas []*fasta, window int, step int, max_len int) ([]*gcskew, error) {

	gcskews := []*gcskew{}

	for _, fasta := range fastas {

		if len(fasta.Sequence) < max_len {
			return gcskews, errors.New("Sequence Length is smaller than max length")
		}

		sequence := fasta.Sequence

		new_gcskew := new(gcskew)

		new_gcskew.ID = fasta.Id

		for i := 0; i < max_len-window; i += step {
			sub_sequence := sequence[i : i+window]
			skew := calcGcskew(sub_sequence)
			new_gcskew.Values = append(new_gcskew.Values, skew)
		}

		gcskews = append(gcskews, new_gcskew)

	}

	return gcskews, nil

}
