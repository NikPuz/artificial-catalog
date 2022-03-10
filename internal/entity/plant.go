package entity

type Plant struct {
	Id uint32 `json:"id"`
	Name string `json:"name"`
	Image_name string `json:"image_name"`
	Height int `json:"height"`
	Preparation int `json:"preparation"`
	Tags []string `json:"tags"`
}

