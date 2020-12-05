package generate

var TplListQuerySection = `
                        <div class="layui-inline">
                            <div class="layui-input-inline">
                                <input type="text" name="{{modelFieldNameFirstLower}}" value="{{.cond.{{modelFieldNameFirstLower}}}}" placeholder="{{modelFieldDesc}}" class="layui-input">
                            </div>
                        </div>`

var TplListThSection = `
									<th>{{modelFieldDesc}}</th>`

var TplListTdSection = `
									<td>{{$vo.{{modelFieldName}}}}</td>`
var TplListTdSectionDate = `
									<td>{{date $vo.{{modelFieldName}} "Y-m-d H:i:s"}}</td>`

var TplList = `<!DOCTYPE html>
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
                    <li class="layui-this">实体{{modelName}}列表</li>
                    <li><a href='{{urlfor "{{modelName}}Controller.Add"}}'>添加实体{{modelName}}</a></li>
                </ul>
                <div class="layui-tab-content">
                    <form class="layui-form layui-form-pane" action='{{urlfor "{{modelName}}Controller.List"}}' method="get">
                        <!-- TODO 参考用，不用就删掉
                        <div class="layui-inline">
                            <div class="layui-input-inline">
                                <select name="status" placeholder="状态" style="width: 60px;">
                                    <option value="">全部状态</option>
                                    <option value="1" selected="selected">成功</option>
                                    <option value="2" selected="selected">失败</option>
                                </select>
                            </div>
                        </div>
                        <div class="layui-inline">
                            <div class="layui-input-inline">
                                <input type="text" name="timeStart" value="" placeholder="起始时间" class="layui-input">
                            </div>
                            <div class="layui-input-inline">
                                <input type="text" name="timeEnd" value="" placeholder="截止时间" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-inline">
                            <div class="layui-input-inline">
                                <input type="text" name="param1" value="" placeholder="订单号 | 名称 | 手机号" class="layui-input">
                            </div>
                        </div>
                        -->{{querySections}}
                        <div class="layui-inline">
                            <button class="layui-btn"><i class="layui-icon layui-icon-search layuiadmin-button-btn"></i></button>
                        </div>
                    </form>
                    <hr>
                    <div class="layui-tab-item layui-show">
                        <table class="layui-table">
                            <thead>
                                <tr>{{thSections}}
									<th></th>
                                </tr>
                            </thead>
                            <tbody>
                            {{range $i, $vo := .data}}
                                <tr>{{tdSections}}
                                    <td>
                                        <a href='{{urlfor "{{modelName}}Controller.Edit" ":id" $vo.Id}}'
                                           class="layui-btn layui-btn-normal layui-btn-xs">编辑</a>
                                        <button href='{{urlfor "{{modelName}}Controller.Del" ":id" $vo.Id}}'
                                                class="layui-btn layui-btn-danger layui-btn-xs ajax-delete">删除
                                        </button>
                                    </td>
                                </tr>
                            {{else}}
                                <tr>
                                    <td colspan="50" style="text-align:center;">没有数据</td>
                                </tr>
                            {{end}}
                            </tbody>
                        </table>
                        {{.Pagination}}
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
{{.Scripts}}
<!-- TODO 参考用，不用这删掉
<script>
    layui.use(['layer','laydate'], function(){
        var laydate = layui.laydate;

        laydate.render({
            elem: 'input[name="timeStart"]',
            type: 'datetime'
        });
        laydate.render({
            elem: 'input[name="timeEnd"]',
            type: 'datetime'
        });
    });
</script>
-->
</body>
</html>
`