package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main () {
    fmt.Println("Collecting Key Bindings from KDE")

    kdeConfigPath := filepath.Join(os.Getenv("HOME"), ".config", "kglobalshortcutsrc")
    outputFilePath := "kde_shortcuts.txt"

    shortcuts, err := parseShortcuts(kdeConfigPath, "Global Shortcuts")
    if err != nil {
        fmt.Printf("Error parsing shortcuts: %v\n", err)
        return
    }

    err = saveShortcutsToFIle(shortcuts, outputFilePath)
    if err != nil {
        fmt.Printf("Error Saving shortcuts to file: %v\n", err)
        return
    }
}

func parseShortcuts(filepath string, sourceName string) ([][3]string, error) {
    content, err := os.ReadFile(filepath)
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(content), "\n")
    var shortcuts [][3]string

    for _, line := range lines {
        if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
            section := strings.TrimSuffix(strings.TrimPrefix(line, "["), "]"),

    }
}
