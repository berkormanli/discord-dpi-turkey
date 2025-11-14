package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/berkormanli/discord-dpi-turkey/internal/config"
	"github.com/berkormanli/discord-dpi-turkey/internal/i18n"
)

// MainUI represents the main application UI
type MainUI struct {
	window fyne.Window
	config *config.Config
	tabs   *container.AppTabs
}

// NewMainUI creates a new main UI
func NewMainUI(window fyne.Window, cfg *config.Config) *MainUI {
	return &MainUI{
		window: window,
		config: cfg,
	}
}

// Build constructs the main UI layout
func (m *MainUI) Build() fyne.CanvasObject {
	// Create tabs for different sections
	m.tabs = container.NewAppTabs(
		container.NewTabItem(i18n.T("tab_wiresock"), m.buildWireSockTab()),
		container.NewTabItem(i18n.T("tab_byedpi"), m.buildByeDPITab()),
		container.NewTabItem(i18n.T("tab_zapret"), m.buildZapretTab()),
		container.NewTabItem(i18n.T("tab_goodbyedpi"), m.buildGoodbyeDPITab()),
		container.NewTabItem(i18n.T("tab_repair"), m.buildRepairTab()),
		container.NewTabItem(i18n.T("tab_advanced"), m.buildAdvancedTab()),
		container.NewTabItem(i18n.T("tab_about"), m.buildAboutTab()),
	)

	// Create top bar with language selector and theme toggle
	topBar := m.buildTopBar()

	// Main layout
	return container.NewBorder(
		topBar, // top
		nil,    // bottom
		nil,    // left
		nil,    // right
		m.tabs, // center
	)
}

// buildTopBar creates the top bar with controls
func (m *MainUI) buildTopBar() fyne.CanvasObject {
	// Language selector
	langSelect := widget.NewSelect(i18n.GetAvailableLanguages(), func(lang string) {
		i18n.SetLanguage(lang)
		m.config.Language = lang
		m.config.Save()
		// Refresh UI
		m.refreshUI()
	})
	langSelect.SetSelected(m.config.Language)

	// Theme toggle button
	var themeBtn *widget.Button
	updateThemeBtn := func() {
		if m.config.IsDarkMode {
			themeBtn.SetText("☀️ Light")
		} else {
			themeBtn.SetText("🌙 Dark")
		}
	}

	themeBtn = widget.NewButton("", func() {
		m.config.IsDarkMode = !m.config.IsDarkMode
		m.config.Save()
		fyne.CurrentApp().Settings().SetTheme(NewCustomTheme(m.config.IsDarkMode))
		updateThemeBtn()
	})
	updateThemeBtn()

	return container.NewHBox(
		widget.NewLabel(i18n.T("app_name")),
		widget.NewLabel(" - "),
		widget.NewLabel("Language:"),
		langSelect,
		themeBtn,
	)
}

// refreshUI refreshes all UI elements with new translations
func (m *MainUI) refreshUI() {
	// Rebuild tabs with new translations
	m.tabs.Items[0].Text = i18n.T("tab_wiresock")
	m.tabs.Items[1].Text = i18n.T("tab_byedpi")
	m.tabs.Items[2].Text = i18n.T("tab_zapret")
	m.tabs.Items[3].Text = i18n.T("tab_goodbyedpi")
	m.tabs.Items[4].Text = i18n.T("tab_repair")
	m.tabs.Items[5].Text = i18n.T("tab_advanced")
	m.tabs.Items[6].Text = i18n.T("tab_about")
	m.tabs.Refresh()
}

// buildWireSockTab creates the WireSock configuration tab
func (m *MainUI) buildWireSockTab() fyne.CanvasObject {
	standardBtn := widget.NewButton(i18n.T("ws_standard_install"), func() {
		m.showInfo("WireSock Standard Installation - Coming Soon")
	})

	alternativeBtn := widget.NewButton(i18n.T("ws_alternative_install"), func() {
		m.showInfo("WireSock Alternative Installation - Coming Soon")
	})

	browserCheck := widget.NewCheck(i18n.T("ws_tunnel_browsers"), func(checked bool) {
		m.config.BrowserTunnel = checked
		m.config.Save()
	})
	browserCheck.SetChecked(m.config.BrowserTunnel)

	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_wiresock")),
		widget.NewSeparator(),
		standardBtn,
		alternativeBtn,
		browserCheck,
		widget.NewButton(i18n.T("ws_customize_folders"), func() {
			m.showInfo("Customize Folders - Coming Soon")
		}),
	)
}

// buildByeDPITab creates the ByeDPI configuration tab
func (m *MainUI) buildByeDPITab() fyne.CanvasObject {
	splitTunnelBtn := widget.NewButton(i18n.T("byedpi_split_tunnel"), func() {
		m.showInfo("ByeDPI Split Tunneling - Coming Soon")
	})

	dllInstallBtn := widget.NewButton(i18n.T("byedpi_dll_install"), func() {
		m.showInfo("ByeDPI DLL Installation - Coming Soon")
	})

	uninstallBtn := widget.NewButton(i18n.T("byedpi_uninstall"), func() {
		m.showInfo("Uninstall ByeDPI - Coming Soon")
	})

	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_byedpi")),
		widget.NewSeparator(),
		splitTunnelBtn,
		dllInstallBtn,
		widget.NewCheck(i18n.T("byedpi_tunnel_browsers"), nil),
		uninstallBtn,
	)
}

// buildZapretTab creates the Zapret configuration tab
func (m *MainUI) buildZapretTab() fyne.CanvasObject {
	scanSpeed := widget.NewSelect(
		[]string{i18n.T("zapret_fast"), i18n.T("zapret_standard"), i18n.T("zapret_full")},
		func(string) {},
	)
	scanSpeed.SetSelected(i18n.T("zapret_standard"))

	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_zapret")),
		widget.NewSeparator(),
		widget.NewButton(i18n.T("zapret_auto_install"), func() {
			m.showInfo("Zapret Auto Install - Coming Soon")
		}),
		container.NewHBox(
			widget.NewLabel(i18n.T("zapret_scan_speed")),
			scanSpeed,
		),
		widget.NewButton(i18n.T("zapret_preset_install"), func() {
			m.showInfo("Zapret Preset Install - Coming Soon")
		}),
		widget.NewButton(i18n.T("zapret_preset_once"), func() {
			m.showInfo("Zapret Run Once - Coming Soon")
		}),
		widget.NewButton(i18n.T("zapret_uninstall"), func() {
			m.showInfo("Uninstall Zapret - Coming Soon")
		}),
	)
}

// buildGoodbyeDPITab creates the GoodbyeDPI configuration tab
func (m *MainUI) buildGoodbyeDPITab() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_goodbyedpi")),
		widget.NewSeparator(),
		widget.NewButton(i18n.T("gdpi_install"), func() {
			m.showInfo("GoodbyeDPI Install - Coming Soon")
		}),
		widget.NewButton(i18n.T("gdpi_run_once"), func() {
			m.showInfo("GoodbyeDPI Run Once - Coming Soon")
		}),
		widget.NewCheck(i18n.T("gdpi_use_blacklist"), nil),
		widget.NewButton(i18n.T("gdpi_uninstall"), func() {
			m.showInfo("Uninstall GoodbyeDPI - Coming Soon")
		}),
	)
}

// buildRepairTab creates the repair tab
func (m *MainUI) buildRepairTab() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_repair")),
		widget.NewSeparator(),
		widget.NewButton(i18n.T("repair_discord"), func() {
			m.showInfo("Repair Discord - Coming Soon")
		}),
		widget.NewButton(i18n.T("repair_install_ptb"), func() {
			m.showInfo("Install Discord PTB - Coming Soon")
		}),
		widget.NewCheck(i18n.T("repair_clean_install"), nil),
		widget.NewLabel(i18n.T("repair_status_checks")),
		widget.NewLabel(i18n.T("repair_discord_standard") + " " + i18n.T("repair_not_installed")),
		widget.NewLabel(i18n.T("repair_discord_ptb") + " " + i18n.T("repair_not_installed")),
	)
}

// buildAdvancedTab creates the advanced configuration tab
func (m *MainUI) buildAdvancedTab() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_advanced")),
		widget.NewSeparator(),
		widget.NewLabel(i18n.T("advanced_services")),
		widget.NewButton(i18n.T("advanced_remove_all"), func() {
			m.showInfo("Remove All Services - Coming Soon")
		}),
		widget.NewButton(i18n.T("advanced_reset_dns"), func() {
			m.showInfo("Reset DNS - Coming Soon")
		}),
		widget.NewButton(i18n.T("advanced_uninstall_app"), func() {
			m.showInfo("Uninstall App - Coming Soon")
		}),
	)
}

// buildAboutTab creates the about tab
func (m *MainUI) buildAboutTab() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(i18n.T("app_name")),
		widget.NewLabel(i18n.T("about_version")+" 2.0.0"),
		widget.NewSeparator(),
		widget.NewLabel(i18n.T("about_description")),
		widget.NewSeparator(),
		widget.NewLabel(i18n.T("about_author")),
		widget.NewLabel(i18n.T("about_license")),
		widget.NewButton(i18n.T("about_github"), func() {
			m.showInfo("GitHub: https://github.com/berkormanli/discord-dpi-turkey")
		}),
		widget.NewButton(i18n.T("about_logs_folder"), func() {
			m.showInfo("Logs folder - Coming Soon")
		}),
	)
}

// showInfo displays an information dialog
func (m *MainUI) showInfo(message string) {
	dialog := widget.NewLabel(message)
	m.window.SetContent(container.NewVBox(
		dialog,
		widget.NewButton(i18n.T("btn_close"), func() {
			m.window.SetContent(m.Build())
		}),
	))
}
