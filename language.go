package fortniteapi

type Language string

const (
	LanguageArabic             Language = "ar"
	LanguageGerman             Language = "de"
	LanguageEnglish            Language = "en"
	LanguageSpanish            Language = "es"
	LanguageSpanish419         Language = "es-419"
	LanguageFrench             Language = "fr"
	LanguageIndonesian         Language = "id"
	LanguageItalian            Language = "it"
	LanguageJapanese           Language = "ja"
	LanguageKorean             Language = "ko"
	LanguagePolish             Language = "pl"
	LanguagePortugueseBrazil   Language = "pt-BR"
	LanguageRussian            Language = "ru"
	LanguageThai               Language = "th"
	LanguageTurkish            Language = "tr"
	LanguageVietnamese         Language = "vi"
	LanguageChineseSimplified  Language = "zh-Hans"
	LanguageChineseTraditional Language = "zh-Hant"
)

type LanguageParams struct {
	Language      Language     `url:"language,omitempty"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}
