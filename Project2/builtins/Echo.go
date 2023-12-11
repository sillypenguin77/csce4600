package builtins
import (
	"strings"
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}