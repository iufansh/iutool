// Copyright 2013 bee authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package generate

import (
	"fmt"
	"github.com/iufansh/iutils"
	"os"
	"path"
	"strings"

	beeLogger "github.com/beego/bee/logger"
	"github.com/beego/bee/logger/colors"
	"github.com/beego/bee/utils"
)

func GenerateController(cname, currpath string) {
	w := colors.NewColorWriter(os.Stdout)

	p, fname := path.Split(cname)
	controllerName := strings.Title(fname)
	if strings.Contains(controllerName, "_") {
		arr := strings.Split(controllerName, "_")
		controllerName = ""
		for _, v := range arr {
			controllerName = controllerName + iutils.UpperFirst(v)
		}
	}
	packageName := "controllers"

	if p != "" {
		i := strings.LastIndex(p[:len(p)-1], "/")
		packageName = p[i+1 : len(p)-1]
	}

	beeLogger.Log.Infof("Using '%s' as controller name", fname)
	beeLogger.Log.Infof("Using '%s' as package name", packageName)

	fp := path.Join(currpath, "controllers", p)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		// Create the controller's directory
		if err := os.MkdirAll(fp, 0777); err != nil {
			beeLogger.Log.Fatalf("Could not create controllers directory: %s", err)
		}
	}

	fpath := path.Join(fp, fname+".go")
	if f, err := os.OpenFile(fpath, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)

		modelPath := path.Join(currpath, "models", fname+".go")

		var content string
		if _, err := os.Stat(modelPath); err == nil {
			modelAttrs := AnalysisModel(modelPath)
			var querySections string
			var updateCols string
			for _, v := range modelAttrs {
				querySections = querySections + strings.ReplaceAll(querySection, "{{modelFieldNameFirstLower}}", iutils.LowerFirst(v.Name))
				if v.Name != "Id" && v.Name != "CreateDate" && v.Name != "Creator" {
					updateCols = updateCols + strings.ReplaceAll(updateColSection, "{{modelFieldName}}", v.Name)
				}
			}

			beeLogger.Log.Infof("Using matching model '%s'", controllerName)
			content = strings.Replace(controllerModelTpl, "{{packageName}}", packageName, -1)
			pkgPath := getPackagePath(currpath)
			content = strings.Replace(content, "{{pkgPath}}", pkgPath, -1)
			content = strings.Replace(content, "{{querySections}}", querySections, -1)
			content = strings.Replace(content, "{{updateCols}}", strings.TrimRight(updateCols, ","), -1)
			content = strings.Replace(content, "{{controllerNameFirstLower}}", iutils.LowerFirst(controllerName), -1)
		} else {
			content = strings.Replace(controllerTpl, "{{packageName}}", packageName, -1)
		}

		content = strings.Replace(content, "{{controllerName}}", controllerName, -1)
		f.WriteString(content)

		// Run 'gofmt' on the generated source code
		utils.FormatSourceCode(fpath)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", fpath, "\x1b[0m")
	} else {
		beeLogger.Log.Fatalf("Could not create controller file: %s", err)
	}
}

var updateColSection = `"{{modelFieldName}}",`

var querySection = `if v := strings.TrimSpace(c.GetString("{{modelFieldNameFirstLower}}")); v != "" {
	query["{{modelFieldNameFirstLower}}"] = v
	cond["{{modelFieldNameFirstLower}}"] = v
}
`

var controllerTpl = `package {{packageName}}

import (
	"github.com/astaxie/beego"
)

// {{controllerName}}Controller operations for {{controllerName}}   test111111111111111111
type {{controllerName}}Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *{{controllerName}}Controller) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Name Create
// @Description create {{controllerName}}
// @Param	body		body 	models.{{controllerName}}	true		"body for {{controllerName}} content"
// @Success 201 {object} models.{{controllerName}}
// @Failure 403 body is empty
// @router / [post]
func (c *{{controllerName}}Controller) Post() {

}

// GetOne ...
// @Name GetOne
// @Description get {{controllerName}} by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.{{controllerName}}
// @Failure 403 :id is empty
// @router /:id [get]
func (c *{{controllerName}}Controller) GetOne() {

}

// GetAll ...
// @Name GetAll
// @Description get {{controllerName}}
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.{{controllerName}}
// @Failure 403
// @router / [get]
func (c *{{controllerName}}Controller) GetAll() {

}

// Put ...
// @Name Put
// @Description update the {{controllerName}}
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.{{controllerName}}	true		"body for {{controllerName}} content"
// @Success 200 {object} models.{{controllerName}}
// @Failure 403 :id is not int
// @router /:id [put]
func (c *{{controllerName}}Controller) Put() {

}

// Delete ...
// @Name Delete
// @Description delete the {{controllerName}}
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *{{controllerName}}Controller) Delete() {

}
`

var controllerModelTpl = `package {{packageName}}

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/iufansh/iufans/controllers/sysmanage"
	"github.com/iufansh/iufans/utils"
	"{{pkgPath}}/models"
	"strings"
)

//  {{controllerName}}Controller operations for {{controllerName}}
type {{controllerName}}Controller struct {
	sysmanage.BaseController
}

// validate {{controllerName}} value
func validate{{controllerName}}(v *models.{{controllerName}}) (hasError bool, errMsg string) {
	valid := validation.Validation{}
	// valid.Required(v.Name, "errmsg").Message("名称必填")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return true, err.Message
		}
	}
	return false, ""
}

// List ...
// @Name List
// @Description get {{controllerName}} list page
// @Success 200 tpl
// @Failure 403
// @router /list [get]
func (c *{{controllerName}}Controller) List() {
	var query = make(map[string]string)
	var fields []string
	var sortby []string
	var order []string

	var cond = make(map[string]string)
    // query 条件设置方式参考 https://beego.me/docs/mvc/model/query.md
	{{querySections}}

	limit, _, offset := c.GetPaginateParam()
	l, total, err := models.GetPaginate{{controllerName}}(query, fields, sortby, order, offset, limit, true)
	if err != nil {
		logs.Error("{{controllerName}}Controller.List GetPaginate{{controllerName}} err:", err)
		c.Msg = "查询异常"
	}
	c.Dta = &l
	c.SetPaginator(limit, total)
	c.SetTplCondition(cond)
	c.RetTpl("{{packageName}}/{{controllerNameFirstLower}}/list.html")
}

// Del ...
// @Name Del
// @Description delete the {{controllerName}}
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /del/:id [post]
func (c *{{controllerName}}Controller) Del() {
	defer c.RetJSON()
	id := c.GetModelId()
	if err := models.Delete{{controllerName}}(id); err == nil {
		c.Code = utils.CODE_OK
		c.Msg = "删除成功"
	} else {
		logs.Error("{{controllerName}}Controller.Del models.Delete err:", err)
		c.Msg = err.Error()
	}
}

// Add ...
// @Name Add
// @Description open add tpl
// @Success add template
// @Failure 404 page not found
// @router /add [get]
func (c *{{controllerName}}Controller) Add() {
	c.RetTpl("{{packageName}}/{{controllerNameFirstLower}}/add.html")
}

// Create ...
// @Name Create
// @Description create {{controllerName}}
// @Param	form		form 	models.{{controllerName}}	true		"form for {{controllerName}} content"
// @Success 201 ok
// @Failure 403 error
// @router /create [post]
func (c *{{controllerName}}Controller) Create() {
	defer c.RetJSON()
	v := models.{{controllerName}}{}
	if err := c.ParseForm(&v); err != nil {
		logs.Error("{{controllerName}}Controller.Create ParseForm err:", err)
		c.Msg = "Param Error"
		return
	} else if hasError, errMsg := validate{{controllerName}}(&v); hasError {
		c.Msg = errMsg
		return
	}
	v.Creator = c.LoginAdminId
	if _, err := models.Add{{controllerName}}(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Code = utils.CODE_OK
		c.Msg = "添加成功"
	} else {
		logs.Error("{{controllerName}}Controller.Create models.Add err:", err)
		c.Msg = err.Error()
	}
}

// Edit ...
// @Name Edit
// @Description get {{controllerName}} by id
// @Param	id		path 	string	true		"The key for {{controllerName}}"
// @Success 200 edit template
// @Failure 403 :id is empty
// @router /edit/:id [get]
func (c *{{controllerName}}Controller) Edit() {
	id := c.GetModelId()
	v, err := models.Get{{controllerName}}ById(id)
	if err != nil {
		logs.Error("{{controllerName}}Controller.Edit models.Get err:", err)
		c.Msg = err.Error()
	} else {
		c.Dta = v
	}
	c.RetTpl("{{packageName}}/{{controllerNameFirstLower}}/edit.html")
}

// Update ...
// @Name Update
// @Description update the {{controllerName}}
// @Param	id		path 	string	true		"The id you want to update"
// @Param	form		form 	models.{{controllerName}}	true		"form for {{controllerName}} content"
// @Success 200 ok
// @Failure 403 :id is not int
// @router /update/:id [post]
func (c *{{controllerName}}Controller) Update() {
	defer c.RetJSON()
	id := c.GetModelId()
	v := models.{{controllerName}}{}
	if err := c.ParseForm(&v); err != nil {
		logs.Error("{{controllerName}}Controller.Update ParseForm err:", err)
		c.Msg = "Param Error"
		return
	} else if hasError, errMsg := validate{{controllerName}}(&v); hasError {
		c.Msg = errMsg
		return
	}
	v.Id = id
	v.Modifior = c.LoginAdminId
	var cols = []string{{{updateCols}}}
	if err := models.Update{{controllerName}}ById(&v, cols...); err == nil {
		c.Code = utils.CODE_OK
		c.Msg = "更新成功"
	} else {
		logs.Error("{{controllerName}}Controller.Update models.Update err:", err)
		c.Msg = err.Error()
	}
}
`
