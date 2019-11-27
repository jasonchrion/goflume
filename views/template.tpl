<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header.html" .}}
<style>
    dd.layui-col-md4 {
        padding: 5px;
        height: 246px;
    }

    .card-add-body {
        line-height: 173px;
        height: 173px;
        text-align: center;
        cursor: pointer;
    }
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a>模板</a>
            <a><cite>模板维护</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-md12 layui-col-space15">
            <div class="layui-card layui-col-md6">
                <div class="layui-card-body">
                    <form enctype="multipart/form-data" method="post" action="/pkg/upload">
                        <input type="file" required="required" name="inputfile" id="inputfile" />
                        <input type="submit" class="layui-btn layui-bg-green" value="导入模板" />
                        <div class="layui-inline layui-word-aux" style="line-height: 38px;margin-left: 10px;">
                            文件大小最大为10MB
                        </div>
                    </form>
                </div>
            </div>
            <div class="layui-card layui-col-md6">
                <div class="layui-card-body">
                    <a href="/pkg/download" class="layui-btn layui-bg-green" id="download">
                        导出模板
                    </a>
                </div>
            </div>
        </div>
        <div class="layui-row layui-col-md12 layui-col-space15">
            <dl>
                <dd class="layui-col-md4">
                    <div class="layui-card">
                        <div class="layui-card-header">新增</div>
                        <div class="layui-card-body card-add-body">
                            <a title="新增" class="layui-col-md12" href="/template/new">
                                <i class="layui-icon layui-icon-add-1" style="font-size: 50px;"></i>
                            </a>
                        </div>
                    </div>
                </dd>
                {{range $t := .templates}}
                <dd class="layui-col-md4">
                    <div class="layui-card">
                        <div class="layui-card-header"><strong>{{$t.Name}}</strong></div>
                        <div class="layui-card-body">
                            <fieldset class="layui-elem-field">
                                <legend>描述</legend>
                                <div class="layui-field-box" style="word-break: break-all;height: 80px;">
                                    {{$t.DESC}}
                                </div>
                            </fieldset>
                            <div class="layui-row">
                                <a href="/template/update?tid={{$t.ID}}" title="修改"
                                    class="layui-btn layui-col-md2 layui-col-md-offset4">
                                    <i class="layui-icon layui-icon-edit"></i>
                                </a>
                                <a data="/template/delete?tid={{$t.ID}}" title="删除"
                                    class="layui-btn layui-col-md2">
                                    <i class="layui-icon layui-icon-delete"></i>
                                </a>
                            </div>
                        </div>
                    </div>
                </dd>
                {{end}}
            </dl>
        </div>
    </div>
    {{template "footer2.html" .}}
    <script>
        layui.use(['tools'], function () {
            layui.tools.setDeleteConfirm("a[title='删除']");
        })
    </script>
</body>

</html>