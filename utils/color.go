package utils

import "fmt"

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextPurple
	TextCyan
	TextWhite
)

func Black(str string) string {
	return textColor(TextBlack, str)
}

func Red(str string) string {
	return textColor(TextRed, str)
}
func Yellow(str string) string {
	return textColor(TextYellow, str)
}
func Green(str string) string {
	return textColor(TextGreen, str)
}
func Cyan(str string) string {
	return textColor(TextCyan, str)
}
func Blue(str string) string {
	return textColor(TextBlue, str)
}
func Purple(str string) string {
	return textColor(TextPurple, str)
}
func White(str string) string {
	return textColor(TextWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}
