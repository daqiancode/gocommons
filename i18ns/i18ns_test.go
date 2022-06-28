package i18ns_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/daqiancode/gocommons/i18ns"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestLocalMessages(t *testing.T) {
	fmt.Println(os.Getwd())
	bundles := i18ns.NewBundle(language.English, "test.en.toml", "test.zh.toml")
	lm_en := i18ns.NewLocalMessages(bundles, "en")
	lm_zh := i18ns.NewLocalMessages(bundles, "zh")
	assert.Equal(t, "Hello jim", lm_en.Get("hello", "Name", "jim"))
	assert.Equal(t, "你好 jim", lm_zh.Get("hello", "Name", "jim"))
	assert.Equal(t, "jim has 1 unread email.", lm_en.GetPlural("UnreadEmails", 1, "Name", "jim"))
	assert.Equal(t, "jim has 2 unread emails.", lm_en.GetPlural("UnreadEmails", 2, "Name", "jim"))

}
