package textgenerator

import (
	"math/rand"
	"time"
)

// GeneratorInterface Text generator interface for expanding functionality
type GeneratorInterface interface {
	Generate(text string) string
	Configure(startTag rune, endTag rune, separator rune) GeneratorInterface
}

// TextGenerator Generator of unique texts by mask
type TextGenerator struct {
	startTag  rune
	endTag    rune
	separator rune
}

// New returns a new instance a text generator.
func New() *TextGenerator {
	rand.Seed(time.Now().UnixNano())
	return &TextGenerator{
		startTag:  '{',
		endTag:    '}',
		separator: '|',
	}
}

// Configure method configures the parser
func (t *TextGenerator) Configure(startTag rune, endTag rune, separator rune) GeneratorInterface {
	t.startTag = startTag
	t.endTag = endTag
	t.separator = separator
	return t
}

// Generate generates and returns a new text.
// Use the rules for generating a plurality of texts.
// Example mask: `Good {morning|day}!`
func (t *TextGenerator) Generate(text string) string {
	return string(t.scanAndReplace([]rune(text)))
}

func (t *TextGenerator) getRandomPart(text []rune) []rune {
	openLevel := 0
	lastPos := 0
	isIgnore := false
	var parts []string
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
		if isIgnore {
			continue
		}
		if text[i] == t.separator {
			parts = append(parts, string(text[lastPos:i]))
			lastPos = i + 1
		}
	}

	parts = append(parts, string(text[lastPos:]))

	return []rune(parts[rand.Intn(len(parts))])
}

func (t *TextGenerator) scanAndReplace(text []rune) []rune {
	startSafePos := 0
	startPos := 0
	endPos := 0
	openLevel := 0
	isFind := false
	var result []rune

	for i := 0; i < len(text); i++ {
		if text[i] == t.startTag {
			if openLevel == 0 {
				startPos = i
				//
				result = append(result, text[startSafePos:startPos]...)
			}

			openLevel++
			continue
		}

		if text[i] == t.endTag {
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

	if !isFind {
		return text
	}

	return append(result, text[endPos+1:]...)
}
