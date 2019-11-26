layui.define(function (exports) {
    var obj = {
        log: function (url) {
            var ws = new WebSocket("ws://" + location.host + url);
            ws.onclose = function () {
                console.log("关闭日志通道");
            }
            var currentLines = 0;
            return {
                open: function (content_id, hello, maxLines, err_test) {
                    var logContent = $('#' + content_id);
                    logContent.html("");
                    var ok = false;
                    ws.onmessage = function (a) {
                        if (ok) {
                            if (-1 != maxLines && currentLines % maxLines == 0) {
                                logContent.html("");
                            }
                            if ((err_test).test(a.data)) {
                                logContent.append("<p class='console-err'>" + a.data + "</p>");
                            } else {
                                logContent.append("<p class='console-p'>" + a.data + "</p>");
                            }
                            if (-1 != maxLines) {
                                currentLines++;
                            }
                            logContent.scrollTop(logContent.prop('scrollHeight'));
                            return;
                        }
                        if (a.data == '200') {
                            console.log("建立日志通道成功");
                            ok = true;
                            if ("" != hello) {
                                ws.send(hello);
                            }
                        }
                    }
                }
                , close: function () {
                    ws.close();
                }
            }
        }
        , create: function (url) {
            var ws = new WebSocket("ws://" + location.host + url);
            ws.onclose = function () {
                console.log("关闭日志通道");
            }
            return ws;
        }
    }

    exports('ws', obj);
});