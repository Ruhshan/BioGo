package biogo

import "regexp"

type IUPACDna struct {
	Sequence string
}

func (d IUPACDna) ToRegex() string {
	s := ""
	// source: https://www.bioinformatics.org/sms/iupac.html
	irm := map[int]string{
		82: "[AG]", 89: "[CT]", 83: "[GC]", 87: "[AT]",
		75: "[GT]", 77: "[AC]", 66: "[CGT]", 68: "[AGT]",
		72: "[ACT]", 86: "[ACG]", 78: "[ATCG]",
		65: "A", 84: "T", 67: "C", 71: "G",
	}

	for _, c := range d.Sequence {
		s += irm[int(c)]
	}

	return s
}

func (d IUPACDna) Match(s string) string {
	pattern := d.ToRegex()
	match, _ := regexp.MatchString(pattern, s)

	if match == true {
		return "1"
	} else {
		return "0"
	}
}
