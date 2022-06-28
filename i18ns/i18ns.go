package i18ns

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type I18n struct {
	bundle *i18n.Bundle
}

func NewBundle(defaultLanguage language.Tag, messageFiles ...string) *i18n.Bundle {

	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, v := range messageFiles {
		bundle.MustLoadMessageFile(v)
	}
	return bundle
}

type LocalMessage struct {
	localizer *i18n.Localizer
	bundle    *i18n.Bundle
}

func NewLocalMessages(bundle *i18n.Bundle, langs ...string) *LocalMessage {
	r := &LocalMessage{
		localizer: i18n.NewLocalizer(bundle, langs...),
	}
	return r
}

// Get messageID
// args can be, 1. empty, 2.map[string]interface, 3. multple args like Get(id, "name","jim" , "age" , 1)
func (s *LocalMessage) Get(messageID string, args ...interface{}) string {
	var arg interface{}
	if len(args) == 1 {
		arg = args[0]
	}
	if len(args) > 1 {
		argMap := make(map[string]interface{}, len(args)/2)
		for i := 0; i < len(args); i += 2 {
			if i+1 < len(args) {
				argMap[args[i].(string)] = args[i+1]
			}
		}
		arg = argMap
	}
	return s.localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID, TemplateData: arg})
}

func (s *LocalMessage) GetPlural(messageID string, count int, args ...interface{}) string {
	var arg interface{}
	if len(args) == 1 {
		arg = args[0]
		if marg, ok := arg.(map[string]interface{}); ok {
			marg["Count"] = count
			arg = marg
		}

	}
	if len(args) > 1 {
		argMap := make(map[string]interface{}, len(args)/2)
		for i := 0; i < len(args); i += 2 {
			if i+1 < len(args) {
				argMap[args[i].(string)] = args[i+1]
			}
		}
		argMap["Count"] = count
		arg = argMap
	}
	return s.localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID, PluralCount: count, TemplateData: arg})
}
