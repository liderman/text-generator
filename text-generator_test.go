package textgenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextGenerator_Generate(t *testing.T) {
	tg := New()
	t1 := tg.Generate("{aaa|aaa}")
	assert.Equal(t, t1, "aaa")

	t2 := tg.Generate("{aaa|aaa} {bbb|bbb}")
	assert.Equal(t, t2, "aaa bbb")

	t3 := tg.Generate("aaa {bbb|bbb} ccc {ddd|ddd} eee")
	assert.Equal(t, t3, "aaa bbb ccc ddd eee")

	t4 := tg.Generate("{{aaa|aaa}|{aaa|aaa}}")
	assert.Equal(t, t4, "aaa")

	t5 := tg.Generate("aaa{bbb{ccc|ccc}ddd|bbb{ccc|ccc}ddd}eee")
	assert.Equal(t, t5, "aaabbbcccdddeee")

	t6 := tg.Generate("aaa")
	assert.Equal(t, t6, "aaa")

	t7 := tg.Generate("{}")
	assert.Equal(t, t7, "")

	t8 := tg.Generate("{|}")
	assert.Equal(t, t8, "")

	t9 := tg.Generate("{")
	assert.Equal(t, t9, "{")

	t10 := tg.Generate("}")
	assert.Equal(t, t10, "}")

	t11 := tg.Generate("|")
	assert.Equal(t, t11, "|")
}

func TestTextGenerator_CustomGenerateConfig(t *testing.T) {
	tg := New().Configure('[', ']', '!')
	t1 := tg.Generate("[aaa!aaa]")
	assert.Equal(t, t1, "aaa")
}
