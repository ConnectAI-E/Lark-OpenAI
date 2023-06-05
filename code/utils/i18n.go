package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 设置全局语言
var lang language.Tag = language.Vietnamese
var I18n = message.NewPrinter(lang)
