package dto

type HomeDetailResp struct {
	Id      int64            `json:"id"`
	Name    string           `json:"name"`
	Belong  int64            `json:"belong"`
	Members []HomeMemberResp `json:"members"`
}

type HomeMemberResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
	AllowDel bool   `json:"allowDel"` // 当前用户是否允许移除家庭成员
}
