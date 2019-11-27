<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header.html" .}}
<style>
    .layui-backlog-body {
        height: 80px;
        display: block;
        background-color: #eee;
        padding: 10px;
        border-radius: 2px;
    }

    .layui-backlog-body h3 {
        color: #666;
    }

    .layui-backlog-body p {
        font-size: 40px;
        line-height: 60px;
    }

    .console-p,
    .console-err {
        margin: 0;
        padding: 2px 20px 2px 20px;
        font-size: 12px;
    }

    .console-err {
        color: #FF5722;
    }
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a><cite>概览</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-space15">
            <div class="layui-col-md6">
                <div class="layui-card" style="height: 270px;">
                    <div class="layui-card-header">采集器信息</div>
                    <div class="layui-card-body">
                        <ul class="layui-row layui-col-space10 layui-this">
                            <li class="layui-col-xs6">
                                <a class="layui-backlog-body">
                                    <h3>总数</h3>
                                    <p><cite th:text="${totalCount}" style="color: #333;">0</cite></p>
                                </a>
                            </li>
                            <li class="layui-col-xs6">
                                <a class="layui-backlog-body">
                                    <h3>运行</h3>
                                    <p><cite th:text="${runCount}" style="color: #009688;">0</cite></p>
                                </a>
                            </li>
                            <li class="layui-col-xs6">
                                <a class="layui-backlog-body">
                                    <h3>关闭</h3>
                                    <p><cite th:text="${totalCount - runCount}" style="color: #FF5722;">0</cite></p>
                                </a>
                            </li>
                            <li class="layui-col-xs6">
                                <a class="layui-backlog-body">
                                    <h3>重启</h3>
                                    <p><cite th:text="${restartCount}" style="color: #FFB800;">0</cite></p>
                                </a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <div class="layui-col-md6">
                <div class="layui-card" style="height: 270px;">
                    <div class="layui-card-header">运行信息</div>
                    <div class="layui-card-body">
                        <table class="layui-table">
                            <colgroup>
                                <col width="120">
                                <col>
                            </colgroup>
                            <tbody>
                                <tr>
                                    <td>连接名称</td>
                                    <td th:text="${jvm.name}">-</td>
                                </tr>
                                <tr>
                                    <td>JVM 版本</td>
                                    <td th:text="${jvm.vmVersion}">-</td>
                                </tr>
                                <tr>
                                    <td>启动时间</td>
                                    <td th:text="${jvm.startTime}">-</td>
                                </tr>
                                <tr>
                                    <td>JVM 内存</td>
                                    <td th:text="${jvm.heap}">-</td>
                                </tr>
                                <tr>
                                    <td>CPU 核心数</td>
                                    <td th:text="${jvm.cpu}">-</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
        <div class="layui-row layui-col-space15">
            <div class="layui-col-md12">
                <div class="layui-card" style="height: 340px;">
                    <div class="layui-card-header">日志</div>
                    <div class="layui-card-body">
                        <div id="log_content" class="layui-row layui-bg-black layui-col-md12"
                            style="padding: 10px;overflow: auto;height: 280px;">
                            <p class='console-err'>UI 日志异常</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</body>
{{template "footer2.html" .}}
<script>
    layui.use(['ws'], function () {
        var builder = layui.ws;
        var log = builder.log("/ws/uilog");
        log.open("log_content"
            , location.host
            , 1000
            , /^at\s|Exception:\s|\[ERROR\]|\([^.]+.java:\d+\)|\(Native Method\)/);
    })
</script>

</html>