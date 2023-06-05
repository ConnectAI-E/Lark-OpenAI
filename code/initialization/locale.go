package initialization

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LoadLocalizer() *i18n.Localizer {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("../locales/en.toml")

	localizer := i18n.NewLocalizer(bundle, language.English.String())
	return localizer
}
