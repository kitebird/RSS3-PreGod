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
	PoapEventInfo PoapEventInfo
	TokenId       string
	Owner         string
}

type PoapResponse struct {
	PoapAction PoapAction
	Created    string
}

type PoapSupply struct {
	Total int `json:"total"`
	Order int `json:"order"`
}

type TokenResponse struct {
	PoapEventInfo PoapEventInfo
	PoapSupply    PoapSupply
}
