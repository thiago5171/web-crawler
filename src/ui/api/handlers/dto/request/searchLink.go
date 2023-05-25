package request

type SearchLink struct {
	Url         string `json:"url"`
	Email       string `json:"email"`
	NumberLinks int    `json:"number_links"`
}
