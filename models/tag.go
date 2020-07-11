package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string `josn:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//获取所有标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 获取标签总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 判断标签存不存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name= ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id= ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//创建标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id= ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}

// 在创建数据之前
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// 在更新数据之前
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
