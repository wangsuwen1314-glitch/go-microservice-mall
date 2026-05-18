// focus.go 管理后台轮播图模块控制器
// 该文件实现了轮播图（focus）的增删改查功能，包含页面渲染和表单处理逻辑。
// 使用 Gin 作为 HTTP 路由和页面渲染框架，操作数据库则通过 models.DB 完成。

package admin

import (
	"fmt"
	"ginshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FocusController 轮播图管理控制器
// 继承自 BaseController，支持统一的错误和成功提示处理。
type FocusController struct {
	BaseController
}

// Index 显示轮播图列表页面
// 查询所有轮播图数据并渲染到 admin/focus/index.html 模板。
func (con FocusController) Index(c *gin.Context) {
	// 定义一个 Focus 结构体切片用于接收查询结果
	focusList := []models.Focus{}
	// 从数据库中查询所有轮播图记录
	models.DB.Find(&focusList)
	// 渲染列表页面，并传入 focusList 数据
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})

}

// Add 显示新增轮播图页面
// 只渲染新增表单页面，不需要额外数据。
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

// DoAdd 处理新增轮播图表单提交
// 从表单获取轮播图信息，校验后保存到数据库。
func (con FocusController) DoAdd(c *gin.Context) {
	// 获取表单字段
	title := c.PostForm("title")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	// 参数校验：如果 focusType 或 status 转换失败，则视为非法请求
	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
		return
	}
	// 如果 sort 不是合法数字，则提示输入正确排序值
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
		return
	}

	// 上传图片文件，返回图片路径
	focusImgSrc, err4 := models.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}

	// 构建轮播图模型对象
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}

	// 将新数据写入数据库
	err5 := models.DB.Create(&focus).Error
	if err5 != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}

}

// Edit 显示编辑轮播图页面
// 根据传入 id 查询要编辑的轮播图，并渲染编辑页面。
func (con FocusController) Edit(c *gin.Context) {
	// 获取 URL 查询参数 id
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/focus")
		return
	}

	// 根据 id 查找对应的轮播图记录
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	// 渲染编辑页面，并传入当前轮播图数据
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

// DoEdit 处理编辑轮播图表单提交
// 从表单获取更新后的数据，并保存到数据库。
func (con FocusController) DoEdit(c *gin.Context) {
	// 获取表单字段
	id, err1 := models.Int(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))

	// 参数校验：id、focusType、status 转换失败视为非法请求
	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "非法请求", "/admin/focus")
		return
	}
	// sort 转换失败时，提示用户输入正确排序值
	if err3 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
		return
	}

	// 上传新的图片文件，如果用户没有上传则返回空字符串
	focusImg, _ := models.UploadImg(c, "focus_img")

	// 先查出当前记录，再更新字段
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	// 只有当用户上传了新图片，才替换原来的图片路径
	if focusImg != "" {
		focus.FocusImg = focusImg
	}

	// 保存修改后的记录到数据库
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+models.String(id))
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}

// Delete 删除轮播图记录
// 根据传入 id 删除对应数据，并返回删除结果。
func (con FocusController) Delete(c *gin.Context) {
	// 获取 URL 查询参数 id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
	} else {
		// 构建要删除的 Focus 记录并执行删除
		focus := models.Focus{Id: id}
		models.DB.Delete(&focus)
		// 如果需要，也可以删除对应的图片文件，示例代码如下：
		// os.Remove("static/upload/20210915/1631694117.jpg")
		con.Success(c, "删除数据成功", "/admin/focus")
	}
}
