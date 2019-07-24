package net_test

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestHTML(t *testing.T) {
	str := "<h1 class=”test”>test</h1>"
	sr := strings.NewReader(str)
	token := html.NewTokenizer(sr)

	for {
		tt := token.Next()
		switch tt {
		case html.ErrorToken:
			break
		case html.TextToken:
			fmt.Println(token.Text())
		case html.StartTagToken:
			//func (z *Tokenizer) TagName() (name []byte, hasAttr bool)
			fmt.Println(token.TagName())
		case html.EndTagToken:
			fmt.Println(token.TagName())
			break
		case html.SelfClosingTagToken:
		case html.CommentToken:
		case html.DoctypeToken:
		}

		if tt.String() == "Error" {
			break
		}

		fmt.Println(tt.String())
	}
}
