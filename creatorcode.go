package fortniteapi

type CreatorCodeParams struct {
	Name          string       `url:"name"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

type CreatorCodeAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatorCodeResponse struct {
	Code     string             `json:"code"`
	Account  CreatorCodeAccount `json:"account"`
	Status   string             `json:"status"`
	Verified bool               `json:"verified"`
}
