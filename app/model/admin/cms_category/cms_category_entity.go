// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package cms_category

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table qxkj_cms_category.
type Entity struct {
	Id             uint64  `orm:"id,primary"      json:"id"`              // 分类id
	ParentId       uint64  `orm:"parent_id"       json:"parent_id"`       // 分类父id
	Status         uint    `orm:"status"          json:"status"`          // 状态,1:发布,0:不发布
	DeleteTime     uint    `orm:"delete_time"     json:"delete_time"`     // 删除时间
	ListOrder      float64 `orm:"list_order"      json:"list_order"`      // 排序
	Name           string  `orm:"name"            json:"name"`            // 分类名称
	Description    string  `orm:"description"     json:"description"`     // 分类描述
	SeoTitle       string  `orm:"seo_title"       json:"seo_title"`       //
	SeoKeywords    string  `orm:"seo_keywords"    json:"seo_keywords"`    //
	SeoDescription string  `orm:"seo_description" json:"seo_description"` //
	ListTpl        string  `orm:"list_tpl"        json:"list_tpl"`        // 分类列表模板
	OneTpl         string  `orm:"one_tpl"         json:"one_tpl"`         // 分类文章页模板
	More           string  `orm:"more"            json:"more"`            // 扩展属性
	CateType       uint    `orm:"cate_type"       json:"cate_type"`       // 分类类型
	CateAddress    string  `orm:"cate_address"    json:"cate_address"`    // 跳转地址
	CateContent    string  `orm:"cate_content"    json:"cate_content"`    // 单页内容
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
