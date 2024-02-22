package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
)

const clipboardSize = 10

// the most recent clipboard should be in the 0 spot 
type clipboardHistory struct {
    history []string
}

func (ch clipboardHistory) getcurrentclipboard() string {
    return ch.history[0]
}

// inserts the new string into the newest pos and cascade down
func (ch *clipboardHistory) insertNewestClipboard(string newClipboard)  {
    oldslice := ch.history

    // cascade the elements down one and copy the newest into the zero slice
    newSlice := make([]string, clipboardSize, clipboardSize)
    copy(newSlice[1:], oldslice[:len(oldslice)-1])
    newSlice[0] = value
}


func initClipboard() *clipboardHistory {
    return &clipboardHistory {
        history: make([]string, 50, 50),
    }
}

func GetCurrentClipboard() (string, error) {
    cmd := exec.Command("xclip", "-o")
    out, err := cmd.Output()
    if err != nil {
        return "", err
    }

    return strings.TrimSpace(string(out)), nil
}

func ClipboardChanged(clipboard *clipboardHistory) bool  {
    newClipboard, err := GetCurrentClipboard()
    if err != nil {
        return false
    }

    // TODO add getting the previous clipboard here 
    if newClipboard != /*TODO*/ {
         return true
    } else {
        return false
    }
}

func PollClipboard() {
    clipboard := initClipboard()
    for {

        fmt.Printf("polling clipboard")
        if ClipboardChanged(clipboard) {
            fmt.Printf("the clipboard has changed to X")
        }


    }
}
