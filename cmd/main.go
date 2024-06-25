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

    err = saveShortcutsToFile(shortcuts, outputFilePath)
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
      u      section := strings.TrimSuffix(strings.TrimPrefix(line, "["), "]")
            continue // Skip section headers
        }

        if len(line) == 0 || strings.HasPrefix(line, "#") {
            continue // Skip empty lines and comments
        }

        parts := strings.Split(line, "=")
        if len(parts) < 2 {
            continue // Slip malformed lines
        }

        action := strings.TrimSpace(parts[0])
        keybinding := strings.TrimSpace(strings.Split(parts[1], ",")[0])

        if keybinding == "none" {
            continue // Skip shortcuts set to "none"
        }

        shortcuts = append(shortcuts, [3]string{keybinding, action, sourceName})
    }

    return shortcuts, nil
}

func saveShortcutsToFile(shortcuts [][3]string, outputFilePath string) error {
    f, err := os.Create(outputFilePath)
    if err != nil {
        return err
    }
    defer f.Close

    // Write Headers
    fmt.Fprintf(f, "%-25s %-50s %-30s\n?", "Keybinding", "Action", "Source")
    fmt.Fprintf(f, "%s\n", strings.Repeat("=", 105))

    // Write Shortcuts
    for _, shortcut := range shortcuts {
        fmt.Fprintf(f, "%-25s %-50s %-30s\n", shortcut[0], shortcut[1], shortcut[2])
    }
    
    fmt.Fprintf("Shortcuts have beem saved to %s\n", outputFilePath)
    return nil
}
