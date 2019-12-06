<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header.html" .}}
<style>
    .layui-card-header>span {
        color: #009688;
        font-size: 24px;
        padding-left: 10px;
    }

    dd.layui-col-md4 {
        padding: 5px;
    }

    legend>span {
        font-size: 12px;
        color: #666;
        margin-left: 10px;
    }

    .layui-elem-quote-cmd {
        height: 160px;
        word-break: break-all;
        overflow: auto;
        color: #666;
    }

    .code-cmd {
        color: #009688;
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

    .card-header {
        border-bottom: none;
    }
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a lay-href=""><cite>采集器监控</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-space15">
            <div class="layui-col-md12">
                <div class="layui-card" style="text-align: center;">
                    <div class="layui-card-header card-header layui-col-md3">采集器总数<span id="totalCount">0</span></div>
                    <div class="layui-card-header card-header layui-col-md3">运行<span id="runCount">0</span></div>
                    <div class="layui-card-header card-header layui-col-md3">停止<span id="stopCount">0</span></div>
                    <div class="layui-card-header card-header">重启<span id="restartCount">0</span></div>
                </div>
            </div>
        </div>
        <div class="layui-row layui-col-space15">
            <dl class="layui-col-md12">
                {{range $s := .states}}
                <dd class="layui-col-md4">
                    <input id="M_{{$s.ID}}" name="metricPort" value="{{$s.State.Port}}" style="display: none">
                    <div class="layui-card">
                        <div class="layui-card-header"><strong>{{$s.Name}}</strong></div>
                        <div class="layui-card-body">
                            <div class="layui-form-item">
                                <label class="layui-form-label">状态</label>
                                <div class="layui-input-block">
                                 {{if eq $s.State.Run 1}}
                                    <input type="text" disabled class="layui-input layui-bg-green" value="运行">
                                {{else if eq $s.State.Run 2}}
                                    <input type="text" disabled class="layui-input layui-bg-blue" value="重启">
                                {{else}}
                                    <input type="text" disabled class="layui-input layui-bg-red" value="关闭">
                                {{end}}
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">进程PID</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled value="{{$s.State.PID}}">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">缓存/容量</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled id="C_{{$s.ID}}" value="NaN">
                                </div>
                            </div>
                            <div class="layui-form-item">
                                <label class="layui-form-label">接收/发送</label>
                                <div class="layui-input-block">
                                    <input type="text" class="layui-input" disabled id="ST_{{$s.ID}}" value="NaN">
                                </div>
                            </div>
                            <fieldset class="layui-elem-field layui-field-title">
                                <legend>启动命令<span>ID:&nbsp;{{$s.ID}}</span></legend>
                            </fieldset>
                            <blockquote class="layui-elem-quote layui-elem-quote-cmd">{{$s.CMD}}</blockquote>
                            <div class="layui-row">
                                <a href="/collect/start?cid={{$s.ID}}"
                                    class="layui-btn layui-col-md2 layui-col-md-offset2" title="启动">
                                    <i class="layui-icon layui-icon-play"></i>
                                </a>
                                <a href="/collect/stop?cid={{$s.ID}}" class="layui-btn layui-col-md2" title="关闭">
                                    <i class="layui-icon layui-icon-close-fill"></i>
                                </a>
                                <a href="/log/download?cid={{$s.ID}}" class="layui-btn layui-col-md2" title="日志下载">
                                    <i class="layui-icon layui-icon-download-circle"></i>
                                </a>
                                <a href="javascript:;" name="watchdog" data="{{$s.ID}}" class="layui-btn layui-col-md2"
                                    title="日志查看">
                                    <i class="layui-icon layui-icon-layer"></i>
                                </a>
                            </div>
                        </div>
                    </div>
                </dd>
                {{end}}
            </dl>
        </div>
    </div>
</body>
{{template "footer2.html" .}}
<script>
    layui.use(['layer', 'element', 'tools', 'ws'], function () {
        layui.tools.setWaitShade("a[title='启动']");
        layui.tools.setWaitShade("a[title='关闭']");

        var total = $("dd").length;
        var dtotal = $(".layui-input.layui-bg-red").length;
        var rtotal = total - dtotal;
        var wtotal = $(".layui-input.layui-bg-blue").length;
        $("#totalCount").text(total);
        $("#runCount").text(rtotal);
        $("#stopCount").text(dtotal);
        $("#restartCount").text(wtotal);

        $("a[name='watchdog']").click(function () {
            var cid = $(this).attr("data");
            var layer = layui.layer;
            var log = layui.ws.log("/ws/log?cid="+cid);
            var height = window.screen.availHeight - 265
            layer.open({
                type: 1
                , area: ['80%', height + 'px']
                , offset: '50px'
                , id: 'logBox'
                , content: '<div id="log_content" class="layui-row layui-col-md12" style="padding: 10px;overflow: auto;max-height: '
                    + (height - 45) + 'px;"></div>'
                , skin: 'layui-bg-black'
                , btnAlign: 'r'
                , title: "日志"
                , yes: function () {
                    layer.closeAll();
                }
                , cancel: function () {
                    log.close();
                }
                , success: function () {
                    $("#logBox").height(height);
                    log.open("log_content"
                        , cid
                        , 2000
                        , /Exception:|(.*\.java:\d)|^\s+at\s|\sERROR\s|^Caused by:|Exception$|\d more/);
                }
            });
        })

        var ws = layui.ws.create("/ws/jmx/metric");
        var ok = false;
        var loadMetric = function () {
            $("input[name='metricPort']").each(function (a, b) {
                var m = $(b);
                var port = m.val();
                if (port >= 40000 && port =< 50000) {
                    var id = m.attr("id").substr(2);
                    ws.send(id);
                }
            })
        }
        ws.onmessage = function (a) {
            if (ok) {
                var id = a.data.substr(0, 32);
                var data = JSON.parse(a.data.substr(33));
                var cap = 0;
                var size = 0;
                var put = 0;
                var take = 0;
                for (var c in data) {
                    if (c.startsWith("CHANNEL")) {
                        cap += parseInt(data[c].ChannelCapacity);
                        size += parseInt(data[c].ChannelSize);
                        put += parseInt(data[c].EventPutSuccessCount);
                        take += parseInt(data[c].EventTakeSuccessCount);
                    }
                }
                $("#C_" + id).val(size + "/" + cap);
                $("#ST_" + id).val(put + "/" + take);
                return
            }
            if (a.data == '200') {
                console.log("管道监控连接建立成功");
                ok = true;
                loadMetric();
            }
        }
        setInterval(loadMetric, 3000);
    })
</script>

</html>