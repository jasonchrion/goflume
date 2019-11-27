<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header.html" .}}
<style>
    dd.layui-col-md4 {
        padding: 5px;
        height: 458px;
    }

    .card-add-body {
        line-height: 385px;
        height: 385px;
        text-align: center;
        cursor: pointer;
    }
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a lay-href=""><cite>采集器维护</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-space15">
            <dl class="layui-col-md12">
                <dd class="layui-col-md4">
                    <div class="layui-card">
                        <div class="layui-card-header">新增</div>
                        <div class="layui-card-body card-add-body">
                            <a title="新增" class="layui-col-md12" href="/collect/new">
                                <i class="layui-icon layui-icon-add-1" style="font-size: 80px;"></i>
                            </a>
                        </div>
                    </div>
                </dd>
                {{range $c := .collects}}
                <dd class="layui-col-md4">
                    <div class="layui-card">
                        <div class="layui-card-header"><strong>{{$c.Name}}</strong></div>
                        <div class="layui-card-body">
                            <div class="layui-form-item">
                                <label class="layui-form-label">厂商</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled value="{{$c.Company}}">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">产品</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled value="{{$c.Product}}">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">产品版本</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled value="{{$c.ProductVersion}}">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">创建时间</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled value="{{$c.CreateTime}}">
                                </div>
                            </div>
                            <fieldset class="layui-elem-field">
                                <legend>描述</legend>
                                <div class="layui-field-box" style="word-break: break-all;height: 80px;">
                                    {{$c.DESC}}
                                </div>
                            </fieldset>
                            <div class="layui-row">
                                <a href="/collect/update?cid={{$c.ID}}" title="修改"
                                    class="layui-btn layui-col-md2 layui-col-md-offset3">
                                    <i class="layui-icon layui-icon-edit"></i>
                                </a>
                                <a data="/collect/delete?cid={{$c.ID}}" title="删除" class="layui-btn layui-col-md2">
                                    <i class="layui-icon layui-icon-delete"></i>
                                </a>
                                <a href="/collect/package?cid={{$c.ID}}" class="layui-btn layui-col-md2" title="配置下载">
                                    <i class="layui-icon layui-icon-download-circle"></i>
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