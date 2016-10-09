package text_generator

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tg := New()
	t1 := tg.Generate("{aaa|aaa}")
	if !reflect.DeepEqual(t1, "aaa") {
		t.Errorf("Not equal: `%s` != `%s`", t1, "aaa")
	}

	t2 := tg.Generate("{aaa|aaa} {bbb|bbb}")
	if !reflect.DeepEqual(t2, "aaa bbb") {
		t.Errorf("Not equal: `%s` != `%s`", t2, "aaa bbb")
	}

	t3 := tg.Generate("aaa {bbb|bbb} ccc {ddd|ddd} eee")
	if !reflect.DeepEqual(t3, "aaa bbb ccc ddd eee") {
		t.Errorf("Not equal: `%s` != `%s`", t3, "aaa bbb ccc ddd eee")
	}

	t4 := tg.Generate("{{aaa|aaa}|{aaa|aaa}}")
	if !reflect.DeepEqual(t4, "aaa") {
		t.Errorf("Not equal: `%s` != `%s`", t4, "aaa")
	}

	t5 := tg.Generate("aaa{bbb{ccc|ccc}ddd|bbb{ccc|ccc}ddd}eee")
	if !reflect.DeepEqual(t5, "aaabbbcccdddeee") {
		t.Errorf("Not equal: `%s` != `%s`", t5, "aaabbbcccdddeee")
	}

	t6 := tg.Generate("aaa")
	if !reflect.DeepEqual(t6, "aaa") {
		t.Errorf("Not equal: `%s` != `%s`", t6, "aaa")
	}

	t7 := tg.Generate("{}")
	if !reflect.DeepEqual(t7, "") {
		t.Errorf("Not equal: `%s` != `%s`", t7, "")
	}

	t8 := tg.Generate("{|}")
	if !reflect.DeepEqual(t8, "") {
		t.Errorf("Not equal: `%s` != `%s`", t8, "")
	}
}

func TestCustomGenerateConfig(t *testing.T) {
	tg := New().Configure('[', ']', '!')
	t1 := tg.Generate("[aaa!aaa]")
	if !reflect.DeepEqual(t1, "aaa") {
		t.Errorf("Not equal: `%s` != `%s`", t1, "aaa")
	}
}

func TestCachedGenerate(t *testing.T) {
	tg := New()
	tgcached := NewCached(tg, 10)
	prevResult := ""
	for i := 0; i <= 100; i++ {
		result := tgcached.Generate("{aaa|bbb|ccc|ddd|eee|fff|ggg|hhh}")
		if i == 0 {
			prevResult = result
			continue
		}
		if !reflect.DeepEqual(result, prevResult) {
			t.Errorf("Not equal: `%s` != `%s`", result, prevResult)
			break
		}
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
		if !reflect.DeepEqual(result, prevResult) {
			t.Errorf("Not equal: `%s` != `%s`", result, prevResult)
			break
		}
	}
}
