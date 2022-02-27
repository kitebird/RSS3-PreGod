package types

type PoapEventInfo struct {
	Id          int    `json:"id"`
	FancyId     string `json:"fancy_id"`
	Name        string `json:"name"`
	EventUrl    string `json:"event_url"`
	ImageUrl    string `json:"image_url"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Supply      int    `json:"supply"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ExpiryDate  string `json:"expiry_date"`
}

type PoapAction struct {
	PoapEventInfo PoapEventInfo `json:"event"`
	TokenId       string        `json:"tokenId"` // nolint:tagliatelle // cause the json field just named "tokenId"
	Owner         string        `json:"owner"`
	Chain         string        `json:"chain"`
}

type PoapResponse struct {
	PoapAction
	Created string `json:"created"`
}

type PoapSupply struct {
	Total int `json:"total"`
	Order int `json:"order"`
}

type TokenResponse struct {
	PoapEventInfo PoapEventInfo `json:"event"`
	PoapSupply    PoapSupply    `json:"supply"`
}
