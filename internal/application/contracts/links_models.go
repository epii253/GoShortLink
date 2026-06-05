package links

type LinkData struct {
	Url string `json:"full_url"`
}

type ShortLinkRequest struct {
	ShortLink string `json:"url"`
}

type LinkAddResponse struct {
	ShortedUrl string `json:"url"`
}

type LinkExtractResponse struct {
	FullUrl string `json:"full_url"`
}

type DeleteLinkRequest struct {
	ShortedUrl string `json:"short_url"`
}

type DeleteLinkResponse struct {
	TotalClicks uint64 `json:"total_clicks"`
}
