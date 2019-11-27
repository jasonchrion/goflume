<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header3.html" .}}

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a href="/template">模板维护</a>
            <a><cite>模板编辑</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-md12 layui-col-space15">
            <div class="layui-card">
                <div class="layui-card-header">模板配置</div>
                <div class="layui-card-body card-add-body">
                    <form class="layui-form" role="form" method="post" action="/template/save">
                        <div class="layui-form-item">
                            <label class="layui-form-label">名称</label>
                            <div class="layui-input-block">
                                <input type="text" name="name" value="{{.t.Name}}" lay-verify="required"
                                    placeholder="请输入名称" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">描述</label>
                            <div class="layui-input-block">
                                <input type="text" name="desc" value="{{.t.DESC}}" lay-verify="required"
                                    placeholder="请输入描述" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item layui-form-text">
                            <label class="layui-form-label">模板配置</label>
                            <div class="layui-input-block" style="height: 380px;overflow: hidden;">
                                <pre class="card-body-pre" style="height: 372px;">
                                    <code id="codeBody" style="height: 372px;margin-top: -24px;" 
                                    class="card-body-code properties">{{.t.Setting}}</code>
                                </pre>
                                <textarea id="codeBody2" name="setting" lay-verify="required" placeholder="请输入模板配置"
                                 rows="18" class="layui-textarea card-body-textarea">{{.t.Setting}}</textarea>
                            </div>
                        </div>
                        <input value="{{.t.ID}}" type="text" style="display: none;" name="tid" />
                        <div class="layui-form-item">
                            <div class="layui-input-block">
                                <button type="submit" class="layui-btn" lay-submit="" lay-filter="save">保存</button>
                                <a href="/template" class="layui-btn layui-btn-primary">关闭</a>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
    {{template "footer3.html" .}}
    <script>
        layui.use(['form', 'layedit'], function () { })
    </script>
</body>

</html>