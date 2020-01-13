package d_video

import (
	"Blog/dao"
	m_video "Blog/model/video"
)

type VideoDao struct {
	dao.BaseDao
}

/**
查找视频列表
*/
func (this *VideoDao) FindVideoList() []m_video.Video {
	var videoList []m_video.Video
	dao.Db.Find(&videoList)

	return videoList
}
