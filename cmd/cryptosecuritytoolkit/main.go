// cmd/cryptosecuritytoolkit/main.go
package main

import (
"flag"
"log"
"os"

"cryptosecuritytoolkit/internal/cryptosecuritytoolkit"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := cryptosecuritytoolkit.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
