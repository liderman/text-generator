package textgenerator

import (
	"fmt"
)

func ExampleSimpleTemplate() {
	tg := New()
	template := "Good {morning|day}!"

	fmt.Print(tg.Generate(template))
}

func ExampleComplexTemplate() {
	tg := New()
	template := "{Good {morning|evening|day}|Goodnight|Hello}, {friend|brother}! {How are you|What's new with you}?"

	fmt.Print(tg.Generate(template))
}
