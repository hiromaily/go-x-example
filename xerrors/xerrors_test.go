package xerrors_test

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/xerrors"
)

//https://godoc.org/golang.org/x/xerrors
//https://qiita.com/sonatard/items/9c9faf79ac03c20f4ae1

func TestXerrors(t *testing.T) {
	// New
	err := xerrors.New("error01")
	fmt.Println(err)

	// Errorf
	err = xerrors.Errorf("error02: %v", err)
	fmt.Println(err)

	// As
	_, err = os.Open("non-existing")
	if err != nil {
		var pathError *os.PathError
		if xerrors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		}
	}
}
