package generate

import (
	"fmt"
	"github.com/beego/bee/logger/colors"
	"github.com/beego/bee/utils"
	"github.com/iufansh/iutils"
	"os"
	"path"
	"strings"

	beeLogger "github.com/beego/bee/logger"
)

func GenerateMcv(sname, currpath string) {
	beeLogger.Log.Infof("是否一次性生成 Model func, Controller, View, Router. [Y：一次性全部生成 | N：每个生成前询问] ")

	if utils.AskForConfirmation() {
		GenerateModelFunc(sname, currpath)
		GenerateController("back/"+sname, currpath)
		GenerateView("back/"+sname, currpath)
		GenerateRouter("back/"+sname, currpath)
		return
	}
	beeLogger.Log.Infof("是否生成 Model func: '%s'? [Y|N] ", sname)

	// Generate the model
	if utils.AskForConfirmation() {
		GenerateModelFunc(sname, currpath)
	}

	beeLogger.Log.Infof("是否生成 Controller: '%s'? [Y|N] ", sname)

	// Generate the model
	if utils.AskForConfirmation() {
		GenerateController("back/"+sname, currpath)
	}

	beeLogger.Log.Infof("是否生成 View: '%s'? [Y|N] ", sname)

	// Generate the model
	if utils.AskForConfirmation() {
		GenerateView("back/"+sname, currpath)
	}

	beeLogger.Log.Infof("是否生成 Router: '%s'? [Y|N] ", sname)

	// Generate the model
	if utils.AskForConfirmation() {
		GenerateRouter("back/"+sname, currpath)
	}

	beeLogger.Log.Successf("All done! Don't forget to add  beego.NSNamespace(\"/%s\" ,beego.NSInclude(&controllers.%sController{})) to routers/route.go\n", sname, strings.Title(sname))
}

func GenerateRouter(cname, currpath string) {
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

	fp := path.Join(currpath, "routers")
	fpath := path.Join(fp, "generate_"+fname+".txt")

	if f, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)

		content := strings.Replace(TplRouter, "{{controllerName}}", controllerName, -1)
		content = strings.Replace(content, "{{controllerNameFirstLower}}", iutils.LowerFirst(controllerName), -1)
		content = strings.Replace(content, "{{packageName}}", iutils.LowerFirst(packageName), -1)
		f.WriteString(content)

		// Run 'gofmt' on the generated source code
		utils.FormatSourceCode(fpath)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", fpath, "\x1b[0m")
	} else {
		beeLogger.Log.Fatalf("Could not create router txt file: %s", err)
	}
}

var TplRouter = `
beego.NSNamespace("{{controllerNameFirstLower}}", beego.NSInclude(&{{packageName}}.{{controllerName}}Controller{})),

{Id: 2020, Pid: 0, Enabled: 1, Display: 1, Description: "实体{{controllerName}}列表", Url: "{{controllerName}}Controller.List", Name: "实体{{controllerName}}列表", Icon: "", Sort: 100},
{Id: 2021, Pid: 2020, Enabled: 1, Display: 0, Description: "删除单条实体{{controllerName}}", Url: "{{controllerName}}Controller.Del", Name: "删除单条实体{{controllerName}}", Icon: "", Sort: 100},
{Id: 2022, Pid: 2020, Enabled: 1, Display: 0, Description: "添加实体{{controllerName}}", Url: "{{controllerName}}Controller.Add", Name: "添加实体{{controllerName}}", Icon: "", Sort: 100},
{Id: 2023, Pid: 2020, Enabled: 1, Display: 0, Description: "提交新增实体{{controllerName}}", Url: "{{controllerName}}Controller.Create", Name: "提交新增实体{{controllerName}}", Icon: "", Sort: 100},
{Id: 2024, Pid: 2020, Enabled: 1, Display: 0, Description: "编辑实体{{controllerName}}", Url: "{{controllerName}}Controller.Edit", Name: "编辑实体{{controllerName}}", Icon: "", Sort: 100},
{Id: 2025, Pid: 2020, Enabled: 1, Display: 0, Description: "提交更新实体{{controllerName}}", Url: "{{controllerName}}Controller.Update", Name: "提交更新实体{{controllerName}}", Icon: "", Sort: 100},
`
