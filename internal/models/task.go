package models

import "net/url"

type Task struct {
	Url     url.URL
	Tittles []string
}