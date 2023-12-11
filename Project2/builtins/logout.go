package builtins

import (
    "fmt"
    "os"
)
func Logout() {
        fmt.Println("Logging out.")
        os.Exit(0)
}

