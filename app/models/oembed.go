// oEmbeded formart refs http://oembed.com/
//

package models

import ()

type Oembed struct {
	Version      int    `json:"version"`
	Type         string `json:"type"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	EmbedHtml    string `json:"embed_html"`
	AuthorName   string `json:"author_name"`
	AuthorUrl    string `json:"author_url"`
	ProviderName string `json:"provider_name"`
	ProviderUrl  string `json:"provider_url"`
}
