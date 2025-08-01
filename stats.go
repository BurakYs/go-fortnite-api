package fortniteapi

type BRStatsByNameParams struct {
	Name string `url:"name"`

	// Enum: "epic", "psn", "xbl"
	//
	// Default: "epic"
	AccountType string `url:"accountType,omitempty"`

	// Enum: "season", "lifetime"
	//
	// Default: "season"
	TimeWindow string `url:"timeWindow,omitempty"`

	// Enum: "all", "keyboardMouse", "gamepad", "touch"
	//
	// Default: *none*
	Image         string       `url:"image,omitempty"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

type BRStatsByIDParams struct {
	// Enum: "season", "lifetime"
	//
	// Default: "season"
	TimeWindow string `url:"timeWindow,omitempty"`

	// Enum: "all", "keyboardMouse", "gamepad", "touch"
	//
	// Default: *none*
	Image         string       `url:"image,omitempty"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

type BRStatsAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BRStatsBattlePass struct {
	Level    int `json:"level"`
	Progress int `json:"progress"`
}

type BRStatsData struct {
	Score           int     `json:"score"`
	ScorePerMin     float64 `json:"scorePerMin"`
	ScorePerMatch   float64 `json:"scorePerMatch"`
	Wins            int     `json:"wins"`
	Top3            int     `json:"top3"`
	Top5            int     `json:"top5"`
	Top6            int     `json:"top6"`
	Top10           int     `json:"top10"`
	Top12           int     `json:"top12"`
	Top25           int     `json:"top25"`
	Kills           int     `json:"kills"`
	KillsPerMin     float64 `json:"killsPerMin"`
	KillsPerMatch   float64 `json:"killsPerMatch"`
	Deaths          int     `json:"deaths"`
	KD              float64 `json:"kd"`
	Matches         int     `json:"matches"`
	WinRate         float64 `json:"winRate"`
	MinutesPlayed   int     `json:"minutesPlayed"`
	PlayersOutlived int     `json:"playersOutlived"`
	LastModified    string  `json:"lastModified"`
}

type BRStatsModes struct {
	Overall BRStatsData `json:"overall"`
	Solo    BRStatsData `json:"solo"`
	Duo     BRStatsData `json:"duo"`
	Trio    BRStatsData `json:"trio"`
	Squad   BRStatsData `json:"squad"`
	LTM     BRStatsData `json:"ltm"`
}

type BRStatsAll BRStatsModes
type BRStatsKeyboardMouse BRStatsModes
type BRStatsGamepad BRStatsModes
type BRStatsTouch BRStatsModes
type BRStatsStats struct {
	All           BRStatsAll           `json:"all"`
	KeyboardMouse BRStatsKeyboardMouse `json:"keyboardMouse"`
	Gamepad       BRStatsGamepad       `json:"gamepad"`
	Touch         BRStatsTouch         `json:"touch"`
}

type BRStatsByNameResponse struct {
	Account    BRStatsAccount    `json:"account"`
	BattlePass BRStatsBattlePass `json:"battlePass"`
	Image      string            `json:"image"`
	Stats      BRStatsStats      `json:"stats"`
}

type BRStatsByIDResponse BRStatsByNameResponse
