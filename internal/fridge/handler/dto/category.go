package dto

type CategoryRsp struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Admin   string `json:"admin"`
	ImgPath string `json:"imgPath"`
	CanEdit bool   `json:"canEdit"`
	Type    int    `json:"type"`
}
