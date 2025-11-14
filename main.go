package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/berkormanli/discord-dpi-turkey/internal/config"
	"github.com/berkormanli/discord-dpi-turkey/internal/gui"
	"github.com/berkormanli/discord-dpi-turkey/internal/i18n"
)

const (
	AppName    = "SplitWire-Turkey"
	AppVersion = "2.0.0"
)

func main() {
	// Check if running with admin/root privileges
	if !hasAdminPrivileges() {
		fmt.Fprintf(os.Stderr, "This application requires administrator/root privileges.\n")
		os.Exit(1)
	}

	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		cfg = config.Default()
	}

	// Initialize translations
	i18n.Init(cfg.Language)

	// Create and run the GUI application
	a := app.New()
	a.Settings().SetTheme(gui.NewCustomTheme(cfg.IsDarkMode))

	w := a.NewWindow(fmt.Sprintf("%s v%s", AppName, AppVersion))
	w.Resize(fyne.NewSize(900, 600))

	// Create main UI
	mainUI := gui.NewMainUI(w, cfg)
	w.SetContent(mainUI.Build())

	// Save configuration on close
	w.SetOnClosed(func() {
		if err := cfg.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to save configuration: %v\n", err)
		}
	})

	w.ShowAndRun()
}

func hasAdminPrivileges() bool {
	// Platform-specific implementation
	return checkAdminPrivileges()
}
