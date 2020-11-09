package model

type Tag struct {
	// id
	Id int32 `json:"id"`
	// 标签名称
	Name string `json:"name"`
	// 状态 0为禁用、1为启用
	State int8 `json:"state"`
	// 创建时间
	CreatedOn int32 `json:"created_on"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 修改时间
	ModifiedOn int32 `json:"modified_on"`
	// 修改人
	ModifiedBy string `json:"modified_by"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
