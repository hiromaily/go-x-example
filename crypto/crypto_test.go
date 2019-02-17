package crypto_test

import (
	"fmt"
	"syscall"
	"testing"

	"golang.org/x/crypto/ssh/terminal"
)

func TestCrypto(t *testing.T) {
	fmt.Print("パスワード > ")

	psw, _ := terminal.ReadPassword(int(syscall.Stdin))

	fmt.Println(psw)
}
