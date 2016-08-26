package text_generator

import (
	"math/rand"
	"time"
)

type TextGenerator struct {
	startTag  rune
	endTag    rune
	separator rune
}

func New() *TextGenerator {
	rand.Seed(time.Now().UnixNano())
	return &TextGenerator{
		startTag:  '{',
		endTag:    '}',
		separator: '|',
	}
}

func (t *TextGenerator) Generate(text string) string {
	return string(t.scanAndReplace([]rune(text)))
}

func (t *TextGenerator) getRandomPart(text []rune) []rune {
	openLevel := 0
	lastPos := 0
	isIgnore := false
	parts := []string{}
	for i := 0; i < len(text); i++ {
		if text[i] == t.startTag {
			openLevel++
			isIgnore = true
			continue
		}
		if text[i] == t.endTag {
			openLevel--
			if openLevel == 0 {
				isIgnore = false
			}
			continue
		}
		if isIgnore == true {
			continue
		}
		if text[i] == t.separator {
			parts = append(parts, string(text[lastPos:i]))
			lastPos = i + 1
		}
	}

	parts = append(parts, string(text[lastPos:len(text)]))

	return []rune(parts[rand.Intn(len(parts))])
}

func (t *TextGenerator) scanAndReplace(text []rune) []rune {
	startSafePos := 0
	startPos := 0
	endPos := 0
	openLevel := 0
	isFind := false
	result := []rune{}

	for i := 0; i < len(text); i++ {
		if text[i] == '{' {
			if openLevel == 0 {
				startPos = i
				//
				result = append(result, text[startSafePos:startPos]...)
			}

			openLevel++
			continue
		}

		if text[i] == '}' {
			openLevel--

			if openLevel == 0 {
				isFind = true
				endPos = i

				//
				startSafePos = i + 1
				result = append(result, t.scanAndReplace(t.getRandomPart(text[startPos+1:endPos]))...)

				continue
			}
		}
	}

	if isFind == false {
		return text
	}

	return append(result, text[endPos+1:]...)
}
