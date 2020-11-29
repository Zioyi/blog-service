package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(2001003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(2001004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
)
