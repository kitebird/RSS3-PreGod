package protocol

type RSS3Index031 struct {
	Version        string `json:"version"`      // "rss3.io/version/v0.3.1"
	ID             string `json:"id"`           // EVM+ address
	DateCreated    string `json:"date_created"` // "2021-08-15T03:00:57.449Z"
	DateUpdated    string `json:"date_updated"` // "2022-02-21T01:56:48.880Z"
	Signature      string `json:"signature"`
	AgentID        string `json:"agent_id"`
	AgentSignature string `json:"agent_signature"`
	Profile        struct {
		Name     string   `json:"name"`
		Bio      string   `json:"bio"`
		Avatar   []string `json:"avatar"`
		Accounts []struct {
			ID        string   `json:"id"`
			Tags      []string `json:"tags"`
			Signature string   `json:"signature"`
		} `json:"accounts"`
	} `json:"profile"`
	Links []struct {
		ID   string `json:"id"`
		List string `json:"list"`
	} `json:"links"`
	Backlinks []struct {
		Auto bool   `json:"auto"` // true
		ID   string `json:"id"`
		List string `json:"list"`
	} `json:"backlinks"`
	Items struct {
		ListAuto string `json:"list_auto"`
	} `json:"items"`
	Assets struct {
		ListAuto string `json:"list_auto"`
	} `json:"assets"`

	CustomFiled_Pass struct {
		Assets []struct {
			ID    string `json:"id"`
			Class string `json:"class"`
			Order int    `json:"order"`
		} `json:"assets"`
	} `json:"_pass"` // nolint:tagliatelle // cause the json field just named "_pass"
}

type RSS3AutoAssets031 struct {
	Version     string   `json:"version"` // "rss3.io/version/v0.3.1"
	ID          string   `json:"id"`
	DateCreated string   `json:"date_created"`
	DateUpdated string   `json:"date_updated"`
	Auto        bool     `json:"auto"` // true
	List        []string `json:"list"`
}

type RSS3Links031 struct {
	Version        string   `json:"version"` // "rss3.io/version/v0.3.1"
	ID             string   `json:"id"`
	DateCreated    string   `json:"date_created"`
	DateUpdated    string   `json:"date_updated"`
	List           []string `json:"list"`
	Signature      string   `json:"signature"`
	AgentID        string   `json:"agent_id"`
	AgentSignature string   `json:"agent_signature"`
}

type RSS3Backlinks031 struct {
	Version     string   `json:"version"` // "rss3.io/version/v0.3.1"
	ID          string   `json:"id"`
	DateCreated string   `json:"date_created"`
	DateUpdated string   `json:"date_updated"`
	Auto        bool     `json:"auto"` // true
	List        []string `json:"list"`
}
