// access.go 管理后台权限模块控制器
// 该文件实现了权限项（access）的增删改查功能，包括页面渲染和表单处理。
// 使用 Gin 框架处理 HTTP 请求，通过 gRPC 调用 RBAC 服务进行数据操作。

package admin

import (
	"context"
	"ginshop/models"
	pbRbac "ginshop/proto/rbacAccess"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AccessController 权限控制器结构体
// 继承自 BaseController，提供权限管理的相关方法。
type AccessController struct {
	BaseController
}

// Index 显示权限列表页面
// 该方法获取所有权限项列表，并渲染到 admin/access/index.html 模板。
func (con AccessController) Index(c *gin.Context) {

	rbacClient := pbRbac.NewRbacAccessService("rbac", models.RbacClient)
	res, _ := rbacClient.AccessGet(context.Background(), &pbRbac.AccessGetRequest{})

	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": res.AccessList,
	})

}

// Add 显示新增权限页面
// 获取所有权限项列表（用于选择上级模块），渲染到 admin/access/add.html 模板。
func (con AccessController) Add(c *gin.Context) {
	//获取顶级模块
	rbacClient := pbRbac.NewRbacAccessService("rbac", models.RbacClient)
	res, _ := rbacClient.AccessGet(context.Background(), &pbRbac.AccessGetRequest{})

	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": res.AccessList,
	})
}

// DoAdd 处理新增权限表单提交
// 从表单中读取权限信息，进行校验后调用角色权限服务添加新权限项。
func (con AccessController) DoAdd(c *gin.Context) {
	// 从表单获取模块名称，并去掉两侧空格
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	// 获取动作名称
	actionName := c.PostForm("action_name")
	// 获取权限类型，并转换为 int
	accessType, err1 := models.Int(c.PostForm("type"))
	// 获取 URL
	url := c.PostForm("url")
	// 获取上级模块 ID，并转换为 int
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	// 获取排序值，并转换为 int
	sort, err3 := models.Int(c.PostForm("sort"))
	// 获取状态，并转换为 int
	status, err4 := models.Int(c.PostForm("status"))
	// 获取描述
	description := c.PostForm("description")
	// 检查所有数字字段是否转换成功
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入参数错误", "/admin/access/add")
		return
	}
	// 检查模块名称是否为空
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}

	// 创建 gRPC 客户端
	rbacClient := pbRbac.NewRbacAccessService("rbac", models.RbacClient)
	// 调用添加权限项的接口
	res, _ := rbacClient.AccessAdd(context.Background(), &pbRbac.AccessAddRequest{
		ModuleName:  moduleName,
		Type:        int64(accessType),
		ActionName:  actionName,
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
	})

	// 检查添加是否成功
	if !res.Success {
		con.Error(c, "增加数据失败", "/admin/access/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/access")

}

// Edit 显示编辑权限页面
// 根据查询参数 id 获取要编辑的权限项和所有权限列表，渲染到 admin/access/edit.html 模板。
func (con AccessController) Edit(c *gin.Context) {
	//获取要修改的数据的 ID
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/access")
	}
	// 获取当前 id 对应的权限项数据
	rbacClient := pbRbac.NewRbacAccessService("rbac", models.RbacClient)
	access, _ := rbacClient.AccessGet(context.Background(), &pbRbac.AccessGetRequest{
		Id: int64(id),
	})

	//获取顶级模块
	resAccess, _ := rbacClient.AccessGet(context.Background(), &pbRbac.AccessGetRequest{})

	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access.AccessList[0],
		"accessList": resAccess.AccessList,
	})
}

// DoEdit 处理编辑权限表单提交
// 从表单读取更新后的权限信息，进行校验后调用角色权限服务更新权限项。
func (con AccessController) DoEdit(c *gin.Context) {
	// 获取权限项 ID
	id, err1 := models.Int(c.PostForm("id"))
	// 获取模块名称并去掉空格
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	// 获取动作名称
	actionName := c.PostForm("action_name")
	// 获取权限类型
	accessType, err2 := models.Int(c.PostForm("type"))
	// 获取 URL
	url := c.PostForm("url")
	// 获取上级模块 ID
	moduleId, err3 := models.Int(c.PostForm("module_id"))
	// 获取排序值
	sort, err4 := models.Int(c.PostForm("sort"))
	// 获取状态
	status, err5 := models.Int(c.PostForm("status"))
	// 获取描述
	description := c.PostForm("description")
	// 检查所有数字字段转换是否成功
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		con.Error(c, "传入参数错误", "/admin/access")
		return
	}
	// 检查模块名称是否为空
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id="+models.String(id))
		return
	}
	// 创建 gRPC 客户端
	rbacClient := pbRbac.NewRbacAccessService("rbac", models.RbacClient)
	// 调用更新权限项的接口
	accessRes, _ := rbacClient.AccessEdit(context.Background(), &pbRbac.AccessEditRequest{
		Id:          int64(id),
		ModuleName:  moduleName,
		Type:        int64(accessType),
		ActionName:  actionName,
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
	})

	// 检查更新是否成功
	if !accessRes.Success {
		con.Error(c, "修改数据失败", "/admin/access/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/access/edit?id="+models.String(id))
	}

}

// Delete 删除权限项
// 根据查询参数 id 删除指定的权限项，通过角色权限服务进行删除。
func (con AccessController) Delete(c *gin.Context) {
	// 获取要删除的权限项 ID
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/access")
	} else {
		// 获取要删除的权限项相关服务客户端
		rbacClient := pbRbac.NewRbacAccessService("rbac", models.RbacClient)
		accessRes, _ := rbacClient.AccessDelete(context.Background(), &pbRbac.AccessDeleteRequest{
			Id: int64(id),
		})
		// 检查删除是否成功
		if !accessRes.Success { //顶级模块
			con.Error(c, accessRes.Message, "/admin/access")
		} else { //操作 或者菜单
			con.Success(c, "删除数据成功", "/admin/access")
		}

	}
}
