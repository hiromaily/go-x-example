package text_test

import (
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//https://blog.golang.org/matchlang
//https://godoc.org/golang.org/x/text/message

func TestMessageFormat(t *testing.T) {
	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Println(123456.78) // Prints 123,456.78

	p.Printf("%d ducks in a row\n", 4331) // Prints 4,331 ducks in a row

	p = message.NewPrinter(message.MatchLanguage("de"))
	p.Printf("Hoogte: %.1f meter\n", 1244.9) // Prints Hoogte: 1,244.9 meter

	p = message.NewPrinter(message.MatchLanguage("ja"))
	p.Println(123456.78) // Prints 123,456.78
}

func TestMessageTranslation(t *testing.T) {
	p := message.NewPrinter(language.English)
	p.Printf("archive(noun)\n") // Prints "archive"
	p.Printf("archive(verb)\n") // Prints "archive"

	p = message.NewPrinter(language.German)
	p.Printf("archive(noun)\n") // Prints "Archiv"
	p.Printf("archive(verb)\n") // Prints "archivieren"
}
