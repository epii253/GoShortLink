package links

type LinkData struct {
	Url string `json:"full_url"`
}

type ShortLinkData struct {
	ShortLink string `json:"url"`
}

type LinkAddResult struct {
	ShortedUrl string `json:"url"`
}

type LinkExtractResult struct {
	FullUrl string `json:"full_url"`
}
