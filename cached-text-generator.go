package textgenerator

import (
	"crypto/md5"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

// CachedTextGenerator Implementing caching to prevent re-generation of phrases
type CachedTextGenerator struct {
	generator GeneratorInterface
	cache     *cache.Cache
}

// NewCached returns a new instance a cached text generator.
func NewCached(generator GeneratorInterface, ttl int) GeneratorInterface {
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

// Configure Configuring tags and separators
func (t *CachedTextGenerator) Configure(startTag rune, endTag rune, separator rune) GeneratorInterface {
	t.generator.Configure(startTag, endTag, separator)
	return t
}
