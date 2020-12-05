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

// recipe
// admin/recipe
func GenerateView(viewpath, currpath string) {
	w := colors.NewColorWriter(os.Stdout)

	beeLogger.Log.Info("Generating view...")

	p, f := path.Split(viewpath)
	modelName := strings.Title(f)
	if strings.Contains(modelName, "_") {
		arr := strings.Split(modelName, "_")
		modelName = ""
		for _, v := range arr {
			modelName = modelName + iutils.UpperFirst(v)
		}
	}
	modelPath := path.Join(currpath, "models", f+".go")

	var modelExist bool
	var content string
	var modelAttrs []ModelAttr
	if _, err := os.Stat(modelPath); err == nil {
		modelExist = true
		modelAttrs = AnalysisModel(modelPath)
	}

	absViewPath := path.Join(currpath, "views", p, iutils.LowerFirst(modelName))
	err := os.MkdirAll(absViewPath, os.ModePerm)
	if err != nil {
		beeLogger.Log.Fatalf("Could not create '%s' view: %s", viewpath, err)
	}

	cfile := path.Join(absViewPath, "list.html")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)

		content = strings.ReplaceAll(TplList, "{{modelName}}", modelName)
		if modelExist {
			var querySections string
			var thSections string
			var tdSections string
			for _, v := range modelAttrs {
				querySection := strings.ReplaceAll(TplListQuerySection, "{{modelFieldNameFirstLower}}", iutils.LowerFirst(v.Name))
				querySection = strings.ReplaceAll(querySection, "{{modelFieldDesc}}", v.Description)
				querySections = querySections + querySection

				thSections = thSections + strings.ReplaceAll(TplListThSection, "{{modelFieldDesc}}", v.Description)

				switch v.DataType {
				case "time.Time":
					tdSections = tdSections + strings.ReplaceAll(TplListTdSectionDate, "{{modelFieldName}}", v.Name)
					break
				default:
					tdSections = tdSections + strings.ReplaceAll(TplListTdSection, "{{modelFieldName}}", v.Name)
				}
			}
			content = strings.ReplaceAll(content, "{{querySections}}", querySections)
			content = strings.ReplaceAll(content, "{{thSections}}", thSections)
			content = strings.ReplaceAll(content, "{{tdSections}}", tdSections)
		}

		f.WriteString(content)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		beeLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "add.html")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)

		content = strings.ReplaceAll(TplAdd, "{{modelName}}", modelName)
		if modelExist {
			var addSections string
			for _, v := range modelAttrs {
				if v.Name == "Id" || v.Name == "CreateDate" || v.Name == "ModifyDate" || v.Name == "Creator" ||
					v.Name == "Modifior" || v.Name == "Version" {
					continue
				}
				section := strings.ReplaceAll(TplAddSection, "{{modelFieldName}}", v.Name)
				section = strings.ReplaceAll(section, "{{modelFieldDesc}}", v.Description)
				if strings.Contains(v.DataType, "int") || strings.Contains(v.DataType, "float") {
					section = strings.ReplaceAll(section, "{{modelFieldType}}", "number")
				} else {
					section = strings.ReplaceAll(section, "{{modelFieldType}}", "text")
				}
				addSections = addSections + section
			}
			content = strings.ReplaceAll(content, "{{addSections}}", addSections)
		}

		f.WriteString(content)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		beeLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "edit.html")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)

		content = strings.ReplaceAll(TplEdit, "{{modelName}}", modelName)
		if modelExist {
			var editSections string
			for _, v := range modelAttrs {
				if v.Name == "Id" || v.Name == "CreateDate" || v.Name == "ModifyDate" || v.Name == "Creator" ||
					v.Name == "Modifior" || v.Name == "Version" {
					continue
				}
				section := strings.ReplaceAll(TplEditSection, "{{modelFieldName}}", v.Name)
				section = strings.ReplaceAll(section, "{{modelFieldDesc}}", v.Description)
				if strings.Contains(v.DataType, "int") || strings.Contains(v.DataType, "float") {
					section = strings.ReplaceAll(section, "{{modelFieldType}}", "number")
				} else {
					section = strings.ReplaceAll(section, "{{modelFieldType}}", "text")
				}
				editSections = editSections + section
			}
			content = strings.ReplaceAll(content, "{{editSections}}", editSections)
		}

		f.WriteString(content)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		beeLogger.Log.Fatalf("Could not create view file: %s", err)
	}

}
