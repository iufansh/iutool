package generate

var TplAddSection = `
                            <div class="layui-form-item">
                                <label class="layui-form-label">{{modelFieldDesc}}</label>
                                <div class="layui-input-block">
                                    <input type="{{modelFieldType}}" name="{{modelFieldName}}" value="" placeholder="请输入{{modelFieldDesc}}"
                                           class="layui-input" required lay-verify="required">
                                </div>
                            </div>`

var TplAdd = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    {{.HtmlHead}}
</head>
<body>
<div class="layui-fluid">
    <div class="layui-row layui-col-space10">
        <div class="layui-col-xs12 layui-col-sm12 layui-col-md12">
            <!--tab标签-->
            <div class="layui-tab layui-tab-brief">
                <ul class="layui-tab-title">
                    <li><a href='{{urlfor "{{modelName}}Controller.List"}}'>实体{{modelName}}列表</a></li>
                    <li class="layui-this">添加实体{{modelName}}</li>
                </ul>
                <div class="layui-tab-content">
                    <div class="layui-tab-item layui-show">
                        <form class="layui-form form-container" action='{{urlfor "{{modelName}}Controller.Create"}}'
                              method="post">{{addSections}}
                            <div class="layui-form-item">
                                <div class="layui-input-block">
                                    <button class="layui-btn" lay-submit lay-filter="*">保存</button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
{{.Scripts}}
</body>
</html>
`