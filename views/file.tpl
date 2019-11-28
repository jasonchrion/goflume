<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header.html" .}}
<style>
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a lay-href=""><cite>文件管理</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-space15">
            <div class="layui-col-md12">
                <div class="layui-card">
                    <div class="layui-card-body">
                        <form enctype="multipart/form-data" method="post" action="/file/upload">
                            <input type="file" required="required" name="inputfile" id="inputfile" />
                            <input type="submit" class="layui-btn layui-bg-green" value="上传文件" />
                            <div class="layui-inline layui-word-aux" style="line-height: 38px;margin-left: 10px;">
                                文件大小最大为10MB
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
        <div class="layui-row layui-col-space15">
            <div class="layui-col-md12">
                <div class="layui-card">
                    <div class="layui-card-body">
                        <table lay-filter="fileTable" style="display: none;">
                            <thead>
                                <tr>
                                    <th lay-data="{field:'num', width:80, sort:true}">序号</th>
                                    <th lay-data="{field:'fileName', width:200, sort:true}">文件名</th>
                                    <th lay-data="{field:'path', sort:true}">路径</th>
                                    <th lay-data="{field:'modifyTime', width:200, sort:true}">修改时间</th>
                                    <th lay-data="{field:'act', width:200, align: 'center'}"></th>
                                </tr>
                            </thead>
                            <tbody>
                            {{range $i,$f := .files}}
                                <tr>
                                    <td>{{$i}}</td>
                                    <td>{{$f.FileName}}</td>
                                    <td>{{$f.FilePath}}</td>
                                    <td>{{$f.UpdateTime}}</td>
                                    <td>
                                        <a class="layui-btn layui-btn-xs" href="/file/update?name={{$f.ShortPath}}" title="修改">
                                            修改
                                        </a>
                                        <a class="layui-btn layui-btn-danger layui-btn-xs" data="/file/delete?name={{$f.ShortPath}}" title="删除">
                                            删除
                                        </a>
                                        <a class="layui-btn layui-btn-primary layui-btn-xs" href="/file/download?name={{$f.ShortPath}}" title="下载">
                                            下载
                                        </a>
                                    </td>
                                </tr>
                            {{end}}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
    {{template "footer2.html" .}}
    <script>
        layui.use(['table', 'tools'], function () {
            layui.table.init('fileTable', { limit: 5000 });
            layui.tools.setDeleteConfirm("a[title='删除']");
        })
    </script>
</body>

</html>