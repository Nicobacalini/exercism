package anagram

import (
	"sort"
	"strings"
	"unicode/utf8"
)

func normalizar(s string) string {
	runas := []rune(strings.ToLower(s))
	sort.Slice(runas, func(i, j int) bool {
		return runas[i] < runas[j]
	})
	return string(runas)
}


func Detect(subject string, candidates []string) []string {
	var resultado []string

    subjectLower := strings.ToLower(subject)
	subjectSorted := normalizar(subject)
	subjectRuneCount := utf8.RuneCountInString(subject)

    for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		if candidateLower == subjectLower {
			continue
		}

		if utf8.RuneCountInString(candidate) != subjectRuneCount {
			continue
		}

		if normalizar(candidate) == subjectSorted {
			resultado = append(resultado, candidate)
		}
	}

	return resultado
}