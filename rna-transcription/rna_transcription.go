package strand

import (
	"strings"
)

var transcribe = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U",
)

func ToRNA(dna string) string {
	return transcribe.Replace(dna)
}
