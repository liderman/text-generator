package text_generator_test

import (
	"fmt"
	"github.com/liderman/text-generator"
)

func ExampleSimpleTemplate() {
	tg := text_generator.New()
	template := "Good {morning|day}!"

	fmt.Print(tg.Generate(template))
}

func ExampleComplexTemplate() {
	tg := text_generator.New()
	template := "{Good {morning|evening|day}|Goodnight|Hello}, {friend|brother}! {How are you|What's new with you}?"

	fmt.Print(tg.Generate(template))
}
