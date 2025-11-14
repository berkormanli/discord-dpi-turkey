package gui

import (
	"fmt"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/berkormanli/discord-dpi-turkey/internal/config"
	"github.com/berkormanli/discord-dpi-turkey/internal/dpi"
	"github.com/berkormanli/discord-dpi-turkey/internal/i18n"
)

// MainUI represents the main application UI
type MainUI struct {
	window  fyne.Window
	config  *config.Config
	tabs    *container.AppTabs
	dpiMgr  *dpi.DPIManager
}

// NewMainUI creates a new main UI
func NewMainUI(window fyne.Window, cfg *config.Config) *MainUI {
	dpiMgr, err := dpi.NewDPIManager(cfg)
	if err != nil {
		// Log error but continue - some features may not work
		dpiMgr = nil
	}
	
	return &MainUI{
		window: window,
		config: cfg,
		dpiMgr: dpiMgr,
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
		if m.dpiMgr == nil {
			m.showError("Service manager not available")
			return
		}
		if err := m.dpiMgr.InstallWireSockStandard(); err != nil {
			m.showError(err.Error())
		} else {
			m.showInfo("WireSock Standard Installation completed successfully")
		}
	})

	alternativeBtn := widget.NewButton(i18n.T("ws_alternative_install"), func() {
		if m.dpiMgr == nil {
			m.showError("Service manager not available")
			return
		}
		if err := m.dpiMgr.InstallWireSockAlternative(); err != nil {
			m.showError(err.Error())
		} else {
			m.showInfo("WireSock Alternative Installation completed successfully")
		}
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
			m.showInfo("Folder customization feature requires external binaries. This feature will be available after installing required tools.")
		}),
	)
}

// buildByeDPITab creates the ByeDPI configuration tab
func (m *MainUI) buildByeDPITab() fyne.CanvasObject {
	splitTunnelBtn := widget.NewButton(i18n.T("byedpi_split_tunnel"), func() {
		if m.dpiMgr == nil {
			m.showError("Service manager not available")
			return
		}
		if err := m.dpiMgr.InstallByeDPISplitTunnel(); err != nil {
			m.showError(err.Error())
		} else {
			m.showInfo("ByeDPI Split Tunneling installed successfully")
		}
	})

	dllInstallBtn := widget.NewButton(i18n.T("byedpi_dll_install"), func() {
		if m.dpiMgr == nil {
			m.showError("Service manager not available")
			return
		}
		if err := m.dpiMgr.InstallByeDPIDLL(); err != nil {
			m.showError(err.Error())
		} else {
			m.showInfo("ByeDPI DLL installation completed successfully")
		}
	})

	uninstallBtn := widget.NewButton(i18n.T("byedpi_uninstall"), func() {
		if m.dpiMgr == nil {
			m.showError("Service manager not available")
			return
		}
		if err := m.dpiMgr.UninstallByeDPI(); err != nil {
			m.showError(err.Error())
		} else {
			m.showInfo("ByeDPI uninstalled successfully")
		}
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
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.InstallZapretAuto(scanSpeed.Selected); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("Zapret auto install completed successfully")
			}
		}),
		container.NewHBox(
			widget.NewLabel(i18n.T("zapret_scan_speed")),
			scanSpeed,
		),
		widget.NewButton(i18n.T("zapret_preset_install"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.InstallZapretPreset("default"); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("Zapret preset install completed successfully")
			}
		}),
		widget.NewButton(i18n.T("zapret_preset_once"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.RunZapretOnce("default"); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("Zapret started temporarily")
			}
		}),
		widget.NewButton(i18n.T("zapret_uninstall"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.UninstallZapret(); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("Zapret uninstalled successfully")
			}
		}),
	)
}

// buildGoodbyeDPITab creates the GoodbyeDPI configuration tab
func (m *MainUI) buildGoodbyeDPITab() fyne.CanvasObject {
	useBlacklist := false
	blacklistCheck := widget.NewCheck(i18n.T("gdpi_use_blacklist"), func(checked bool) {
		useBlacklist = checked
	})
	
	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_goodbyedpi")),
		widget.NewSeparator(),
		widget.NewButton(i18n.T("gdpi_install"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.InstallGoodbyeDPI("default", useBlacklist); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("GoodbyeDPI installed successfully")
			}
		}),
		widget.NewButton(i18n.T("gdpi_run_once"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.RunGoodbyeDPIOnce("default", useBlacklist); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("GoodbyeDPI started temporarily")
			}
		}),
		blacklistCheck,
		widget.NewButton(i18n.T("gdpi_uninstall"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.UninstallGoodbyeDPI(); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("GoodbyeDPI uninstalled successfully")
			}
		}),
	)
}

// buildRepairTab creates the repair tab
func (m *MainUI) buildRepairTab() fyne.CanvasObject {
	cleanInstall := false
	cleanInstallCheck := widget.NewCheck(i18n.T("repair_clean_install"), func(checked bool) {
		cleanInstall = checked
	})
	
	var standardStatus, ptbStatus *widget.Label
	standardStatus = widget.NewLabel("")
	ptbStatus = widget.NewLabel("")
	
	updateStatus := func() {
		if m.dpiMgr != nil {
			standard, ptb := m.dpiMgr.GetDiscordStatus()
			if standard {
				standardStatus.SetText(i18n.T("repair_discord_standard") + " " + i18n.T("repair_installed"))
			} else {
				standardStatus.SetText(i18n.T("repair_discord_standard") + " " + i18n.T("repair_not_installed"))
			}
			if ptb {
				ptbStatus.SetText(i18n.T("repair_discord_ptb") + " " + i18n.T("repair_installed"))
			} else {
				ptbStatus.SetText(i18n.T("repair_discord_ptb") + " " + i18n.T("repair_not_installed"))
			}
		}
	}
	updateStatus()
	
	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_repair")),
		widget.NewSeparator(),
		widget.NewButton(i18n.T("repair_discord"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.RepairDiscord(); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("Discord repair completed")
				updateStatus()
			}
		}),
		widget.NewButton(i18n.T("repair_install_ptb"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.InstallDiscordPTB(cleanInstall); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("Discord PTB installation initiated")
				updateStatus()
			}
		}),
		cleanInstallCheck,
		widget.NewLabel(i18n.T("repair_status_checks")),
		standardStatus,
		ptbStatus,
	)
}

// buildAdvancedTab creates the advanced configuration tab
func (m *MainUI) buildAdvancedTab() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(i18n.T("tab_advanced")),
		widget.NewSeparator(),
		widget.NewLabel(i18n.T("advanced_services")),
		widget.NewButton(i18n.T("advanced_remove_all"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.RemoveAllServices(); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("All services removed successfully")
			}
		}),
		widget.NewButton(i18n.T("advanced_reset_dns"), func() {
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.ResetDNS(); err != nil {
				m.showError(err.Error())
			} else {
				m.showInfo("DNS settings reset successfully")
			}
		}),
		widget.NewButton(i18n.T("advanced_uninstall_app"), func() {
			m.showInfo("To uninstall SplitWire-Turkey, please remove all services first, then delete the application using your system's package manager or manually delete the application files.")
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
			if m.dpiMgr == nil {
				m.showError("Service manager not available")
				return
			}
			if err := m.dpiMgr.OpenLogsFolder(); err != nil {
				m.showError(err.Error())
			}
		}),
	)
}

// showInfo displays an information dialog
func (m *MainUI) showInfo(message string) {
	dialog.ShowInformation("Information", message, m.window)
}

// showError displays an error dialog
func (m *MainUI) showError(message string) {
	dialog.ShowError(fmt.Errorf("%s", message), m.window)
}
