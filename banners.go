package fortniteapi

type BannersParams LanguageParams

type BannerImages struct {
	SmallIcon string `json:"smallIcon"`
	Icon      string `json:"icon"`
}

type BannerRarity BRCosmeticRarity
type BannerSeries BRCosmeticSeries
type BannerSet BRCosmeticSet
type BannerIntroduction BRCosmeticIntroduction

type Banner struct {
	ID              string             `json:"id"`
	DevName         string             `json:"devName"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Category        string             `json:"category"`
	FullUsageRights bool               `json:"fullUsageRights"`
	Rarity          BannerRarity       `json:"rarity"`
	Series          BannerSeries       `json:"series,omitzero"`
	Set             BannerSet          `json:"set,omitzero"`
	Introduction    BannerIntroduction `json:"introduction"`
	Images          BannerImages       `json:"images"`
}

type BannersResponse []Banner

type BannerColors struct {
	ID               string `json:"id"`
	Color            string `json:"color"`
	Category         string `json:"category"`
	SubCategoryGroup int    `json:"subCategoryGroup"`
}

type BannerColorsResponse []BannerColors
