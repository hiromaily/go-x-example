package crypto_test

import (
	"fmt"
	"syscall"
	"testing"

	"golang.org/x/crypto/ssh/terminal"
)

func TestCrypto(t *testing.T) {
	//This code doesn't run from test, it should run as main func
	fmt.Print("password > ")

	password, _ := terminal.ReadPassword(int(syscall.Stdin))

	fmt.Println(string(password))
}
