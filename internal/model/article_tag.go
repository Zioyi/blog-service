package model

type ArticleTag struct {
	// id
	Id int32 `json:"id"`
	// 文章ID
	ArticleId int32 `json:"article_id"`
	// 标签ID
	TagId int32 `json:"tag_id"`
	// 创建时间
	CreatedOn int32 `json:"created_on"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 修改时间
	ModifiedOn int32 `json:"modified_on"`
	// 修改人
	ModifiedBy string `json:"modified_by"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
