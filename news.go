package fortniteapi

type NewsParams LanguageParams
type BRNewsParams LanguageParams
type STWNewsParams LanguageParams
type CreativeNewsParams LanguageParams

type NewsMotd struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	TabTitle        string `json:"tabTitle"`
	Body            string `json:"body"`
	Image           string `json:"image"`
	TileImage       string `json:"tileImage"`
	SortingPriority int    `json:"sortingPriority"`
	Hidden          bool   `json:"hidden"`
	WebsiteURL      string `json:"websiteUrl"`
	VideoString     string `json:"videoString"`
	VideoID         string `json:"videoId"`
}

type NewsMessage struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Image   string `json:"image"`
	Adspace string `json:"adspace"`
}

type News struct {
	Hash     string        `json:"hash"`
	Date     string        `json:"date"`
	Image    string        `json:"image"`
	Motds    []NewsMotd    `json:"motds"`
	Messages []NewsMessage `json:"messages"`
}

type NewsResponse struct {
	BR       News `json:"br"`
	STW      News `json:"stw"`
	Creative News `json:"creative"`
}

type BRNewsResponse News
type STWNewsResponse News
type CreativeNewsResponse News
