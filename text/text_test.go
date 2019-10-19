package text_test

import (
	"testing"
	//"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//https://blog.golang.org/matchlang
//https://godoc.org/golang.org/x/text/message

func TestMessageFormat(t *testing.T) {
	p := message.NewPrinter(message.MatchLanguage("en"))
	if _, err := p.Println(123456.78); err != nil {
		t.Fatalf("fail to call p.Println(123456.78) %v", err)
	} // Prints 123,456.78

	if _, err := p.Printf("%d ducks in a row\n", 4331); err != nil {
		t.Fatalf(`fail to call p.Printf("%d ducks in a row\n", 4331)`, err)
	} // Prints 4,331 ducks in a row

	p = message.NewPrinter(message.MatchLanguage("de"))
	if _, err := p.Printf("Hoogte: %.1f meter\n", 1244.9); err != nil {
		t.Fatalf(`"Hoogte: %.1f meter\n", 1244.9)`, err)
	} // Prints Hoogte: 1,244.9 meter

	p = message.NewPrinter(message.MatchLanguage("ja"))
	if _, err := p.Println(123456.78); err != nil {
		t.Fatalf("fail to call p.Println(123456.78) %v", err)
	} // Prints 123,456.78

	//この方法はカンマが表示されない
	//p = message.NewPrinter(language.Japanese)
	//cur, _ := currency.FromTag(language.Japanese)
	//p.Printf("%d\n", currency.NarrowSymbol(cur.Amount(120000.0)))
}

func TestMessageTranslation(t *testing.T) {
	p := message.NewPrinter(language.English)
	p.Printf("archive(noun)\n") // Prints "archive"
	p.Printf("archive(verb)\n") // Prints "archive"

	p = message.NewPrinter(language.German)
	p.Printf("archive(noun)\n") // Prints "Archiv"
	p.Printf("archive(verb)\n") // Prints "archivieren"

}
