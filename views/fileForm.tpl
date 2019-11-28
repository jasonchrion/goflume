<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header3.html" .}}

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a lay-href="">文件管理</a>
            <a><cite>文件编辑</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-md12 layui-col-space15">
            <div class="layui-card">
                <div class="layui-card-header">文件内容</div>
                <div class="layui-card-body card-add-body">
                    <form class="layui-form" role="form" method="post" action="/file/save">
                        <div class="layui-form-item">
                            <label class="layui-form-label">名称</label>
                            <div class="layui-input-block">
                                <input type="text" name="name" id="codeName" value="{{.fi.FileName}}" disabled
                                    class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item layui-form-text">
                            <label class="layui-form-label">文件内容</label>
                            <div class="layui-input-block" style="height: 420px;overflow: hidden;">
                                <pre class="card-body-pre" style="height: 420px;">
                                    <code id="codeBody" style="height: 400px;margin-top: -24px;" class="card-body-code">{{.fi.Content}}</code>
                                </pre>
                                <textarea name="content" id="codeBody2" placeholder="请输入内容"
                                    rows="20" class="layui-textarea card-body-textarea">{{.fi.Content}}</textarea>
                            </div>
                        </div>
                        <input value="{{.fi.ShortPath}}" type="text" hidden="true" name="path" />
                        <div class="layui-form-item">
                            <div class="layui-input-block">
                                <button type="submit" class="layui-btn" lay-submit="" lay-filter="save">保存</button>
                                <a href="/file" class="layui-btn layui-btn-primary">关闭</a>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
    {{template "footer3.html" .}}
    <script>
        layui.use(['form', 'layedit'], function () {

        })
    </script>
</body>

</html>