package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
)

const clipboardSize = 10

type ClipboardHistory interface {
    PollClipboardHistory()
}

// the most recent clipboard should be in the 0 spot 
type clipboardHistoryImpl struct {
    history []string
}

func InitClipboard() ClipboardHistory {
    return &clipboardHistoryImpl {
        history: make([]string, 50, 50),
    }
}

func (ch clipboardHistoryImpl) currentNewestclipboard() string {
    return ch.history[0]
}

func (ch *clipboardHistoryImpl) UpdateClipboard()  {
    fmt.Printf("updating clipboard\n")
    nch, err := GetCurrentClipboard() 
    if err != nil {
        // TODO
        return 
    }
    ch.insertNewestClipboard(nch)
}

// inserts the new string into the newest pos and cascade down
func (ch *clipboardHistoryImpl) insertNewestClipboard(newClipboard string)  {
    oldslice := ch.history
    fmt.Printf(fmt.Sprintf("the old slice was %v\n", oldslice))

    // cascade the elements down one and copy the newest into the zero slice
    newSlice := make([]string, clipboardSize, clipboardSize)
    copy(newSlice[1:], oldslice[:len(oldslice)-1])
    newSlice[0] = newClipboard
    fmt.Printf(fmt.Sprintf("the new slice is %v\n", newSlice))

    ch.history = newSlice
}

func (ch* clipboardHistoryImpl) HasChanged() bool {
    newClipboard, err := GetCurrentClipboard()
    if err != nil {
        return false
    }

    if newClipboard != ch.currentNewestclipboard() {
         return true
    } else {
        return false
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

func ClipboardChanged(clipboard *clipboardHistoryImpl) bool  {
    newClipboard, err := GetCurrentClipboard()
    if err != nil {
        return false
    }

    if newClipboard != clipboard.currentNewestclipboard() {
         return true
    } else {
        return false
    }
}

// main entrypoint
func (ch clipboardHistoryImpl) PollClipboardHistory() {
    fmt.Printf("polling clipboard")

    for {
        if ch.HasChanged() {
            // update the clipboard
            ch.UpdateClipboard()
        }
    }
}
