package biogo

import (
	"errors"
	"fmt"
)


type propeller_twist struct{
	ID string
	Values [] string
}

func calcPropTwist(sequence string)string{
	
	stacking_energy := map[string]float64{
		"AA":-18.66,"AC":-13.10,"AG":-14.00,"AT":-15.01,
        "CA":-9.45,"CC":-8.11,"CG":-10.03,"CT":-14.00,
        "GA":-14.48,"GC":-11.08,"GG":-8.11,"GT":-13.10,
        "TA":-11.85,"TC":-13.48,"TG":-9.45,"TT":-18.66,
	}
	energy := 0.0
	for i:=0;i<len(sequence)-1;i+=2{
		energy+=stacking_energy[sequence[i:i+2]]
	}

	return fmt.Sprintf("%.6f", energy)
}

func PropellerTwistSlidingWindow(fastas []*fasta, window int, step int, max_len int) ([]*propeller_twist, error) {
	propeller_twists := []*propeller_twist{}

	for _, fasta := range fastas {

		if len(fasta.Sequence) < max_len {
			return propeller_twists, errors.New("Sequence Length is smaller than max length")
		}

		sequence := fasta.Sequence

		new_propeller_twist := new(propeller_twist)

		new_propeller_twist.ID = fasta.Id

		for i := 0; i < max_len-window; i += step {
			sub_sequence := sequence[i : i+window]
			energy := calcPropTwist(sub_sequence)
			new_propeller_twist.Values = append(new_propeller_twist.Values, energy)
		}

		propeller_twists = append(propeller_twists, new_propeller_twist)

	}

	return propeller_twists, nil
}