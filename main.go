package main

import (
	"fmt"

	"github.com/cambo9p/clipboard/clipboard"
)


func main() {
    s, err := clipboard.GetCurrentClipboard()
    if err != nil {
        fmt.Printf("something went wrong")
    }
    fmt.Printf(fmt.Sprintf("the clipboard is %s", s))

    go clipboard.PollClipboard()

    // setup 


}
