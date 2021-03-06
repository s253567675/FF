package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// 定义表模型-苹果IPA包表
type Download struct {
	DownloadID string    `gorm:"primary_key;column:download_id;type:varchar(100)" comment:"下载id"`
	Path       string    `gorm:"not null;column:path;type:varchar(2000)" comment:"文件路径"`
	CreatedAt  time.Time `gorm:"not null" comment:"创建时间"`
}

// 设置表名
func (d Download) TableName() string {
	return "download"
}

// 创建初始化表
func initDownload() {
	if !db.HasTable(&Download{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Download{}).Error; err != nil {
			panic(err)
		}
	}
}

// 添加
func InsertDownloadPath(path string) (string, error) {
	var id = fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	if err := db.Create(&Download{
		Path:       path,
		DownloadID: id,
	}).Error; err != nil {
		return "", err
	}
	return id, nil
}

func GetDownloadPathByID(id string) (string, error) {
	var (
		download Download
		err      error
	)
	if err = db.Where("download_id = ?", id).First(&download).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return download.Path, nil
}

// 根据id删除
func DeleteDownloadPathByID(id string) error {
	return db.Where("download_id = ?", id).Delete(&Download{}).Error
}
