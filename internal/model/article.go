package model

type Article struct {
	// id
	Id int32 `json:"id"`
	// 文章标题
	Title string `json:"title"`
	// 文章简述
	Desc string `json:"desc"`
	// 封面图片地址
	CoverImageUrl string `json:"cover_image_url"`
	// 文章内容
	Content string `json:"content"`
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

func (a Article) TableName() string {
	return "blog_article"
}
