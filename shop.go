package fortniteapi

type ShopParams LanguageParams

type ShopItemBundle struct {
	Name  string `json:"name"`
	Info  string `json:"info"`
	Image string `json:"image"`
}

type ShopItemBanner struct {
	Value        string `json:"value"`
	Intensity    string `json:"intensity"`
	BackendValue string `json:"backendValue"`
}

type ShopItemOfferTag struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type ShopItemLayoutTextureMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ShopItemLayoutStringMetadata ShopItemLayoutTextureMetadata
type ShopItemLayoutTextMetadata ShopItemLayoutTextureMetadata

type ShopItemLayout struct {
	ID                   string                          `json:"id"`
	Name                 string                          `json:"name"`
	Category             string                          `json:"category"`
	Index                int                             `json:"index"`
	Rank                 int                             `json:"rank"`
	ShowIneligibleOffers string                          `json:"showIneligibleOffers"`
	Background           string                          `json:"background"`
	UseWidePreview       bool                            `json:"useWidePreview"`
	DisplayType          string                          `json:"displayType"`
	TextureMetadata      []ShopItemLayoutTextureMetadata `json:"textureMetadata"`
	StringMetadata       []ShopItemLayoutStringMetadata  `json:"stringMetadata"`
	TextMetadata         []ShopItemLayoutTextMetadata    `json:"textMetadata"`
}

type ShopItemColors struct {
	Color1              string `json:"color1"`
	Color2              string `json:"color2"`
	Color3              string `json:"color3"`
	TextBackgroundColor string `json:"textBackgroundColor"`
}

type ShopItemNewDisplayAssetMaterialInstance struct {
	ID          string            `json:"id"`
	PrimaryMode string            `json:"primaryMode"`
	ProductTag  string            `json:"productTag"`
	Images      map[string]string `json:"Images"`   //nolint:tagliatelle
	Colors      map[string]string `json:"Colors"`   //nolint:tagliatelle
	Scalings    map[string]string `json:"Scalings"` //nolint:tagliatelle
	Flags       map[string]bool   `json:"Flags"`    //nolint:tagliatelle
}

type ShopItemNewDisplayAssetRenderImage struct {
	ProductTag string `json:"productTag"`
	FileName   string `json:"fileName"`
	Image      string `json:"image"`
}

type ShopItemNewDisplayAsset struct {
	ID                string                                    `json:"id"`
	CosmeticID        string                                    `json:"cosmeticId"`
	MaterialInstances []ShopItemNewDisplayAssetMaterialInstance `json:"materialInstances"`
	RenderImages      []ShopItemNewDisplayAssetRenderImage      `json:"renderImages"`
}

type ShopItem struct {
	RegularPrice           int                     `json:"regularPrice"`
	FinalPrice             int                     `json:"finalPrice"`
	DevName                string                  `json:"devName"`
	OfferID                string                  `json:"offerId"`
	InDate                 string                  `json:"inDate"`
	OutDate                string                  `json:"outDate"`
	Bundle                 ShopItemBundle          `json:"bundle,omitzero"`
	Banner                 ShopItemBanner          `json:"banner,omitzero"`
	OfferTag               ShopItemOfferTag        `json:"offerTag,omitzero"`
	Giftable               bool                    `json:"giftable"`
	Refundable             bool                    `json:"refundable"`
	SortPriority           int                     `json:"sortPriority"`
	LayoutID               string                  `json:"layoutId"`
	Layout                 ShopItemLayout          `json:"layout"`
	Colors                 ShopItemColors          `json:"colors"`
	TileBackgroundMaterial string                  `json:"tileBackgroundMaterial"`
	TileSize               string                  `json:"tileSize"`
	DisplayAssetPath       string                  `json:"displayAssetPath"`
	NewDisplayAssetPath    string                  `json:"newDisplayAssetPath"`
	NewDisplayAsset        ShopItemNewDisplayAsset `json:"newDisplayAsset"`
	BRItems                []BRCosmetic            `json:"brItems,omitempty"`
	Tracks                 []Track                 `json:"tracks,omitempty"`
	Instruments            []Instrument            `json:"instruments,omitempty"`
	Cars                   []Car                   `json:"cars,omitempty"`
	LegoKits               []LegoKit               `json:"legoKits,omitempty"`
}

type ShopResponse struct {
	Hash      string     `json:"hash"`
	Date      string     `json:"date"`
	VBuckIcon string     `json:"vbuckIcon"`
	Entries   []ShopItem `json:"entries"`
}
