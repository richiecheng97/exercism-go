package protein

import "errors"

var (
	ErrStop        = errors.New("Err Stop")
	ErrInvalidBase = errors.New("Err Invalid Base")
)

var mp = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	var res []string
	start := 0
	for start != len(rna) {
		v, err := FromCodon(rna[start : start+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		res = append(res, v)
		start += 3
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	if v, ok := mp[codon]; !ok {
		return "", ErrInvalidBase
	} else {
		if v == "STOP" {
			return "", ErrStop
		}
		return v, nil
	}
}
