package builtins

import (
	"io"
	"os"
	"fmt"
)

func Pwd(w io.Writer) error {
    wd, err := os.Getwd()
    if err != nil {
        return err
    }
    _, err = fmt.Fprintln(w, wd)
    return err
}