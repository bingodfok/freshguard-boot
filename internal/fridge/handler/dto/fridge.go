package dto

type CreateFridgeReq struct {
	Name string `json:"name" binding:"required"`
}

type FridgeRep struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Admin   string `json:"admin"`
	CanEdit bool   `json:"canEdit"`
}
