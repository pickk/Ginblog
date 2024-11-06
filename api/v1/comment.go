package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"github.com/wejectchen/ginblog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddComment 新增评论
func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)

	code := model.AddComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetComment 获取单个评论信息
func GetComment(c *gin.Context) {
	// 从 URL 参数中获取评论的 ID，将其转换为整数
	// strconv.Atoi 将字符串转换为整数，如果转换失败则返回 0
	id, _ := strconv.Atoi(c.Param("id"))

	// 调用 model.GetComment 函数，通过评论 ID 获取评论数据和状态码
	data, code := model.GetComment(id)

	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                   // 状态码，用于表示操作是否成功
		"data":    data,                   // 评论数据
		"message": errmsg.GetErrMsg(code), // 根据状态码获取相应的错误信息
	})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteComment(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCommentCount 获取评论数量
func GetCommentCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	total := model.GetCommentCount(id)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

// GetCommentList 后台查询评论列表
func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetCommentList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetCommentListFront 展示页面显示评论列表
func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetCommentListFront(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// CheckComment 通过审核
func CheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.CheckComment(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// UncheckComment 撤下评论审核
func UncheckComment(c *gin.Context) {
	var data model.Comment // 声明一个 Comment 类型的变量来存储请求中的数据

	// 从请求体中解析 JSON 数据并绑定到 data 变量
	_ = c.ShouldBindJSON(&data)

	// 从请求的 URL 路径中获取 id 参数，并尝试将其转换为整型
	id, _ := strconv.Atoi(c.Param("id"))

	// 调用 model 中的 UncheckComment 函数处理撤审操作，并获取返回的状态码
	code := model.UncheckComment(id, &data)

	// 使用 HTTP 状态码 200 返回 JSON 格式的响应，包括操作的状态码和对应的消息文本
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                   // 操作的状态码
		"message": errmsg.GetErrMsg(code), // 根据状态码获取并返回错误消息
	})
}
