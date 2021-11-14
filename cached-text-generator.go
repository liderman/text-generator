// Package text_generator Fast text generator on a mask.
package text_generator

import (
	"crypto/md5"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type CachedTextGenerator struct {
	generator TextGeneratorInterface
	cache     *cache.Cache
}

// NewCached returns a new instance a cached text generator.
func NewCached(generator TextGeneratorInterface, ttl int) TextGeneratorInterface {
	return &CachedTextGenerator{
		generator: generator,
		cache:     cache.New(time.Duration(ttl)*time.Second, 30*time.Second),
	}
}

// Generate generates and returns a new text.
// Use the rules for generating a plurality of texts.
// Example mask: `Good {morning|day}!`
func (t *CachedTextGenerator) Generate(text string) string {
	key := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	data, found := t.cache.Get(key)
	if found {
		return data.(string)
	}

	result := t.generator.Generate(text)
	t.cache.Set(key, result, cache.DefaultExpiration)
	return result
}

func (t *CachedTextGenerator) Configure(startTag rune, endTag rune, separator rune) TextGeneratorInterface {
	t.generator.Configure(startTag, endTag, separator)
	return t
}
