package biogo

import (
	"errors"
	"fmt"
)

type base_stacking struct {
	ID     string
	Values []string
}

func calcBaseStaking(sequence string) string {

	stacking_energy := map[string]float64{
		"AA": -5.37, "AC": -10.51, "AG": -6.78, "AT": -6.57,
		"CA": -6.57, "CC": -8.26, "CG": -9.69, "CT": -6.78,
		"GA": -9.81, "GC": -14.59, "GG": -8.26, "GT": -10.51,
		"TA": -3.82, "TC": -9.81, "TG": -6.57, "TT": -5.37,
	}
	energy := 0.0
	for i := 0; i < len(sequence)-1; i += 2 {
		energy += stacking_energy[sequence[i:i+2]]
	}

	return fmt.Sprintf("%.6f", energy)
}

func BaseStackingSlidingWindow(fastas []*fasta, window int, step int, max_len int) ([]*base_stacking, error) {
	base_stackings := []*base_stacking{}

	for _, fasta := range fastas {

		if len(fasta.Sequence) < max_len {
			return base_stackings, errors.New("Sequence Length is smaller than max length")
		}

		sequence := fasta.Sequence

		new_base_stacking := new(base_stacking)

		new_base_stacking.ID = fasta.Id

		for i := 0; i < max_len-window; i += step {
			sub_sequence := sequence[i : i+window]
			energy := calcBaseStaking(sub_sequence)
			new_base_stacking.Values = append(new_base_stacking.Values, energy)
		}

		base_stackings = append(base_stackings, new_base_stacking)

	}

	return base_stackings, nil
}
