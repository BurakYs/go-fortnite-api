package fortniteapi

type PlaylistsParams LanguageParams
type PlaylistByIDParams PlaylistsParams

type PlaylistsImages struct {
	Showcase    string `json:"showcase"`
	MissionIcon string `json:"missionIcon"`
}

type Playlist struct {
	ID                       string          `json:"id"`
	Name                     string          `json:"name"`
	SubName                  string          `json:"subName"`
	Description              string          `json:"description"`
	GameType                 string          `json:"gameType"`
	RatingType               string          `json:"ratingType"`
	MinPlayers               int             `json:"minPlayers"`
	MaxPlayers               int             `json:"maxPlayers"`
	MaxTeams                 int             `json:"maxTeams"`
	MaxTeamSize              int             `json:"maxTeamSize"`
	MaxSquads                int             `json:"maxSquads"`
	MaxSquadSize             int             `json:"maxSquadSize"`
	IsDefault                bool            `json:"isDefault"`
	IsTournament             bool            `json:"isTournament"`
	IsLimitedTimeMode        bool            `json:"isLimitedTimeMode"`
	IsLargeTeamGame          bool            `json:"isLargeTeamGame"`
	AccumulateToProfileStats bool            `json:"accumulateToProfileStats"`
	Images                   PlaylistsImages `json:"images"`
	GameplayTags             []string        `json:"gameplayTags"`
	Path                     string          `json:"path"`
	Added                    string          `json:"added"`
}

type PlaylistsResponse []Playlist
type PlaylistByIDResponse Playlist
