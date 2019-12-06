<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header3.html" .}}
<style>
    .template-list-box {
        height: 540px;
        overflow: auto;
    }

    .template-list {
        display: block;
        margin: 5px;
        padding: 5px;
        border-bottom: 1px solid #e6e6e6;
    }
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a href="/collect">采集器维护</a>
            <a><cite>采集器编辑</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-md12 layui-col-space15">
            <div class="layui-col-md8">
                <div class="layui-card">
                    <div class="layui-card-header">采集器配置</div>
                    <div class="layui-card-body card-add-body">
                        <form class="layui-form" role="form" method="post" action="/collect/save">
                            <div class="layui-form-item">
                                <label class="layui-form-label">名称</label>
                                <div class="layui-input-block">
                                    <input type="text" name="name" value="{{.c.Name}}" lay-verify="required"
                                        placeholder="请输入名称" class="layui-input">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">厂商</label>
                                <div class="layui-input-block">
                                    <input type="text" name="company" value="{{.c.Company}}" placeholder="请输入厂商"
                                        class="layui-input">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">产品名称</label>
                                <div class="layui-input-block">
                                    <input type="text" name="product" value="{{.c.Product}}" placeholder="请输入产品名称"
                                        class="layui-input">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">产品版本</label>
                                <div class="layui-input-block">
                                    <input type="text" name="productVersion" value="{{.c.ProductVersion}}"
                                        placeholder="请输入产品版本" class="layui-input">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">描述</label>
                                <div class="layui-input-block">
                                    <input type="text" name="desc" value="{{.c.DESC}}" lay-verify="required"
                                        placeholder="请输入描述" class="layui-input">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">内存大小(MB)</label>
                                <div class="layui-input-block">
                                    {{if eq .c.MemSize ""}}
                                        <input type="text" name="memSize" value="2048" placeholder="请输入内存大小" class="layui-input">
                                    {{else}}
                                        <input type="text" name="memSize" value="{{.c.MemSize}}" placeholder="请输入内存大小" class="layui-input">
                                    {{end}}
                                </div>
                            </div>
                            <div class="layui-form-item layui-form-text">
                                <label class="layui-form-label">配置 <i id="bigSetting"
                                        class="layui-icon layui-icon-fonts-code layui-bg-green"
                                        style="border-radius: 2px;padding: 2px;cursor: pointer;"></i>
                                </label>
                                <div class="layui-input-block">
                                    <textarea name="setting" lay-verify="required" placeholder="请输入采集配置" rows="6"
                                        class="layui-textarea card-body-textarea-show">{{.c.Setting}}</textarea>
                                </div>
                            </div>
                            <input value="{{.c.ID}}" type="text" hidden="true" name="cid" />
                            <div class="layui-form-item">
                                <div class="layui-input-block">
                                    <button type="submit" class="layui-btn" lay-submit="" lay-filter="save">保存</button>
                                    <a href="/collect" class="layui-btn layui-btn-primary">关闭</a>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <div class="layui-col-md4">
                <div class="layui-card">
                    <div class="layui-card-header">模板参考列表</div>
                    <div class="layui-card-body card-add-body template-list-box">
                    {{range $t := .templates}}
                        <a href="javascript:;" class="layui-row template-list" title="{{$t.Setting}}">
                            <span>{{$t.Name}}</span>
                        </a>
                    {{end}}
                    </div>
                </div>
            </div>
        </div>
    </div>
    {{template "footer3.html" .}}
    <script>
        layui.use(['form', 'layedit', 'layer'], function () {
            var layer = layui.layer;
            $("#bigSetting").click(function () {
                layer.open({
                    type: 1
                    , area: ['80%', '600px']
                    , offset: '50px'
                    , id: 'settingBox'
                    , move: false
                    , content: `
                    <pre class="card-body-pre" style="height: 492px;">
                        <code id="codeBody" style="height: 100%;margin-top: -16px;" class="card-body-code properties">...</code>
                    </pre>
                    <textarea class="layui-textarea card-body-textarea" rows="24" id="codeBody2"></textarea>
                    `
                    , title: "采集器配置"
                    , btn: ['确定']
                    , yes: function () {
                        var setting = $("#codeBody2").val();
                        $("textarea[name='setting']").text(setting);
                        layer.closeAll();
                    }
                    , cancel: function () {
                        layer.closeAll();
                    }
                    , success: function () {
                        var setting = $("textarea[name='setting']").text();
                        $("#codeBody").text(setting);
                        $("#codeBody2").text(setting);
                        loadhljs();
                    }
                })
            });
            $("a.template-list").click(function () {
                var setting = $(this).attr("title");
                $("textarea[name='setting']").text(setting);
            })
        })
    </script>
</body>

</html>