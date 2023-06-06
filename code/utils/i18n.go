package utils

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 设置全局语言
var lang language.Tag = language.Vietnamese
var I18n = message.NewPrinter(lang)

func SetI18n(s string) {
	lang = language.MustParse(s)
	I18n = message.NewPrinter(lang)
	fmt.Printf("SetI18n: %v\n", I18n.Sprintf("已选择AI模式:"))
}
