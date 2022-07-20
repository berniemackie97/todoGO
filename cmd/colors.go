package todo

import "fmt"

const (
	colorDefualt = "\x1b[39m"

	colorRed   = "\x1b[91m"
	colorGreen = "\x1b[32m"
	colorBlue  = "\x1b[94m"
	colorGrey  = "\x1b[90m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", colorRed, s, colorDefualt)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", colorGreen, s, colorDefualt)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", colorBlue, s, colorDefualt)
}

func grey(s string) string {
	return fmt.Sprintf("%s%s%s", colorGrey, s, colorDefualt)
}
