package global

import "gorm.io/gorm"

var (
	DB            *gorm.DB          // 数据库接口
	FEED_NUM      = 30              // 每次返回的视频数量
	MAX_FILE_SIZE = int64(10 << 20) // 上传文件大小限制
)
