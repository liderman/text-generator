package textgenerator

import (
	"testing"
)

func BenchmarkGenerateEasyText(b *testing.B) {
	tg := New()
	for i := 0; i < b.N; i++ {
		tg.Generate("Good {morning|day}!")
	}
}

func BenchmarkGenerateComplexText(b *testing.B) {
	tg := New()
	for i := 0; i < b.N; i++ {
		tg.Generate("{Good {morning|evening|day}|Goodnight|Hello}, {friend|brother}! {How are you|What's new with you}?")
	}
}
