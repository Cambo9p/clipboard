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

    // setup front end  -- will need to pass clipboard in to view
    // need to make something that looks like
    // need to do some dependancy injectiong with an interface 
    // clipboardObj := clipboard.InitClipboard()
    // clipboardObj.pollclipboard() 
    // -- then we can pass the clipboard object into the UI 

    cp := clipboard.InitClipboard()

    go cp.PollClipboardHistory()


}
