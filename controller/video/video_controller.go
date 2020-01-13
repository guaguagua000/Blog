package c_video

import (
	"Blog/dao"
	"Blog/entity"
	"Blog/log"
	m_video "Blog/model/video"
	"Blog/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
查询视频列表接口
*/
func FindVideoList(ctx *gin.Context) {
	videoList := videoService.FindVideoList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}

/**
按条件查询视频列表接口
*/
func FindVideoByWhere(ctx *gin.Context) {
	/**
	绑定query参数到对象.
	注意:默认情况下, "query参数或form参数"和对象成员的大小写需保持一致. 除非为对象补充tag定义, 定义规则如下:
	1. 对于query参数, 需要如下定义
	form:"siteId" binding:"required"
	2. 对于form参数, 需要如下定义
	form:"siteId"
	3. 对于body的json参数, 无需定义tag, 且不区分大小写

	建议用ShouldBind, 不要使用其他bind方法, 否则可能会返回400错误, 参考链接:https://learnku.com/docs/gin-gonic/2018/gin-readme/3819
	*/
	videoQuery := m_video.Video{}
	if err := ctx.ShouldBind(&videoQuery); err != nil {
		log.Logger.Error("绑定请求参数到对象异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  "参数错误",
		})
		return
	}
	log.Logger.Info("绑定请求参数到对象", zap.Any("videoQuery", videoQuery))

	videoList := []m_video.Video{}
	dao.Db.Where("site_id = ? and title like ?", videoQuery.SiteId, "%7%").Find(&videoList)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}

/**
添加视频
*/
func AddVideo(ctx *gin.Context) {
	video := m_video.Video{}
	video.ID = util.GenUniqueId()
	video.Title = "title4"
	jsonTime := &entity.JsonTime{time.Now()}
	video.CreatedAt = jsonTime
	video.UpdatedAt = jsonTime
	dao.Db.Create(video)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
批量添加视频
*/
func BulkAddVideo(ctx *gin.Context) {
	videoService.BulkAddVideo()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
更新视频
*/
func UpdateVideo(ctx *gin.Context) {
	//绑定参数到对象
	videoBind := m_video.Video{}
	if ctx.ShouldBind(&videoBind) == nil {
		log.Logger.Info("绑定请求参数到对象", zap.Any("videoBind", videoBind))
	}

	dao.Db.Model(&m_video.Video{}).Where("id = ?", "id4").Update(videoBind)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
删除视频
*/
func DeleteVideo(ctx *gin.Context) {
	//绑定参数到对象
	videoBind := m_video.Video{}
	if ctx.ShouldBind(&videoBind) == nil {
		log.Logger.Info("绑定请求参数到对象", zap.Any("videoBind", videoBind))
	}

	dao.Db.Where("id = ?", videoBind.ID).Delete(&m_video.Video{})

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}
