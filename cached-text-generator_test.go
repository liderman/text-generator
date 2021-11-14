package textgenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCachedTextGenerator_Generate(t *testing.T) {
	tg := New()
	tgcached := NewCached(tg, 10)
	prevResult := ""
	for i := 0; i <= 100; i++ {
		result := tgcached.Generate("{aaa|bbb|ccc|ddd|eee|fff|ggg|hhh}")
		if i == 0 {
			prevResult = result
			continue
		}
		assert.Equal(t, result, prevResult)
	}
}

func TestCustomCachedGenerate(t *testing.T) {
	tg := New().Configure('[', ']', '!')
	tgcached := NewCached(tg, 10)
	prevResult := ""
	for i := 0; i <= 100; i++ {
		result := tgcached.Generate("[aaa!bbb!ccc!ddd!eee!fff!ggg!hhh]")
		if i == 0 {
			prevResult = result
			continue
		}
		assert.Equal(t, result, prevResult)
	}
}

func TestCachedTextGenerator_CustomGenerateConfig(t *testing.T) {
	tg := New()
	NewCached(tg, 10).Configure('[', ']', '!')
	assert.Equal(t, tg.startTag, rune('['))
	assert.Equal(t, tg.endTag, rune(']'))
	assert.Equal(t, tg.separator, rune('!'))
}
