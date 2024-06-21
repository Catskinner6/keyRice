import configparser
import os

def parse_shortcuts(file_path, source_name):
    config = configparser.RawConfigParser()
    config.read(file_path)

    shortcuts = []
    for section in config.sections():
        for key, value in config.items(section):
            if "none" not in value.lower():
                action = key
                keybinding = value.split(',')[0].strip()
                shortcuts.append((keybinding, action, source_name))

    return shortcuts

def save_shortcuts_to_file(shortcuts, output_file):
    with open(output_file, 'w') as f:
        f.write(f"{'Keybinding':<25} {'Action':<50} {'Source':<30}\n")
        f.write("="*105 + "\n")
        for keybinding, action, source in shortcuts:
            f.write(f"{keybinding:<25} {action:<50} {source:<30}\n")

def main():
    kde_config_path = os.path.expanduser("~/.config/kglobalshortcutsrc")
    output_file = "kde_shortcuts.txt"
    
    if not os.path.exists(kde_config_path):
        print("KDE global shortcuts configuration file not found.")
        return

    shortcuts = parse_shortcuts(kde_config_path, "Global Shortcuts")
    save_shortcuts_to_file(shortcuts, output_file)
    print(f"Shortcuts have been saved to {output_file}")

if __name__ == "__main__":
    main()

