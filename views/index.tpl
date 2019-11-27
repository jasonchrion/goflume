<!DOCTYPE html>
<html>

{{template "header.html" .}}
<style>
    .content-div {
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        overflow: hidden;
    }

    .content-iframe {
        position: absolute;
        width: 100%;
        height: 100%;
        left: 0;
        top: 0;
        right: 0;
        bottom: 0;
        border: none;
    }
</style>

<body class="layui-layout-body">
    <div class="layui-layout layui-layout-admin">
        <div class="layui-header">
            <div class="layui-logo">{{.Title}}</div>
            <ul class="layui-nav layui-layout-left">
                <li class="layui-nav-item">
                    <a href="/index" title="首页">
                        <i class="layui-icon layui-icon-home" style="color: #333;"></i>
                    </a>
                </li>
                <li class="layui-nav-item">
                    <a href="javascript:reloadPage();" title="刷新">
                        <i class="layui-icon layui-icon-refresh" style="color: #333;"></i>
                    </a>
                </li>
                <span class="layui-nav-bar" style="left: 94px; top: 48px; width: 0px; opacity: 0;"></span>
            </ul>
            <ul class="layui-nav layui-layout-right">
                <li class="layui-nav-item" lay-unselect="">
                    <a href="/readme" target="blank" style="color: #333;">
                        说明文档&nbsp;<i class="layui-icon layui-icon-help"></i>
                    </a>
                </li>
            </ul>
        </div>

        <div class="layui-side layui-bg-black">
            <div class="layui-side-scroll">
                <ul class="layui-nav layui-nav-tree" lay-filter="menu">
                    <li class="layui-nav-item layui-this">
                        <a data="/info" href="javascript:;">概览</a>
                    </li>
                    <li class="layui-nav-item">
                        <a data="/state" href="javascript:;">采集器监控</a>
                    </li>
                    <li class="layui-nav-item">
                        <a href="javascript:;">模板</a>
                        <dl class="layui-nav-child">
                            <dd><a data="/template" href="javascript:;">模板维护</a></dd>
                            <dd><a data="/tourist" href="javascript:;">图形化配置</a></dd>
                        </dl>
                    </li>
                    <li class="layui-nav-item">
                        <a data="/collect" href="javascript:;">采集器维护</a>
                    </li>
                    <li class="layui-nav-item">
                        <a data="/file" href="javascript:;">文件管理</a>
                    </li>
                    <li class="layui-nav-item">
                        <a href="javascript:;">调试</a>
                        <dl class="layui-nav-child">
                            <dd><a data="/flume" href="javascript:;">Flume管理</a></dd>
                            <dd><a data="/filewatcher" href="javascript:;">配置文件监控</a></dd>
                            <dd><a data="/test/js" href="javascript:;">JS测试</a></dd>
                            <dd><a data="/test/datafix" href="javascript:;">DataFix测试</a></dd>
                        </dl>
                    </li>
                </ul>
            </div>
        </div>

        <div class="layui-body">
            <!-- 内容主体区域 -->
            <div class="content-div">
                <iframe id="content_iframe" class="content-iframe" src="/info"></iframe>
            </div>
        </div>

        <div class="layui-footer">
            © flume.manager
        </div>
    </div>
    <script src="/static/assets/js/jquery-3.2.1.min.js"></script>
    <script src="/static/layui/layui.js"></script>
    <script src="/static/js/flumeui.js"></script>
    <script>
        function reloadPage() {
            var iframe = document.getElementById('content_iframe');
            iframe.src = iframe.src;
        }
        layui.use(['element'], function () {
            var element = layui.element;
            element.on('nav(menu)', function (elem) {
                var url = elem.attr("data");
                if (undefined != url && "" != url) {
                    $("#content_iframe").attr({ src: url });
                }
            });
        });
    </script>
</body>

</html>