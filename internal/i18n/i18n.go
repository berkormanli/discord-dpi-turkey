package i18n

import (
	"sync"
)

var (
	currentLanguage = "EN"
	mu              sync.RWMutex
	translations    = map[string]map[string]string{
		"EN": englishTranslations,
		"TR": turkishTranslations,
		"RU": russianTranslations,
	}
)

// Init initializes the translation system with a language
func Init(lang string) {
	SetLanguage(lang)
}

// SetLanguage changes the current language
func SetLanguage(lang string) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := translations[lang]; ok {
		currentLanguage = lang
	} else {
		currentLanguage = "EN" // fallback
	}
}

// T returns the translation for a key
func T(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	if trans, ok := translations[currentLanguage][key]; ok {
		return trans
	}

	// Fallback to English
	if trans, ok := translations["EN"][key]; ok {
		return trans
	}

	return key // Return key if no translation found
}

// GetCurrentLanguage returns the current language code
func GetCurrentLanguage() string {
	mu.RLock()
	defer mu.RUnlock()
	return currentLanguage
}

// GetAvailableLanguages returns list of available languages
func GetAvailableLanguages() []string {
	return []string{"EN", "TR", "RU"}
}
