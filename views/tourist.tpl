<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">

{{template "header3.html" .}}
<style>
    .tourist-box {
        height: 540px;
    }

    .tourist-list-box {
        height: 442px;
        overflow: auto;
    }

    .tourist-list {
        display: block;
        margin: 5px;
        padding: 5px;
        border-bottom: 1px solid #e6e6e6;
    }

    .form-group {
        margin-bottom: 1rem;
    }

    .form-group label {
        float: left;
        display: block;
        padding: 9px 15px;
        width: 120px;
        font-weight: 400;
        line-height: 20px;
        text-align: right;
    }

    .form-group .form-control {
        width: 560px;
        height: 38px;
        line-height: 38px;
        border-radius: 2px;
        border: 1px solid #e6e6e6;
        padding-left: 10px;
    }

    .form-group small {
        padding-left: 15px;
        margin-left: 135px;
    }

    .form-text {
        display: block;
        font-size: 80%;
        font-weight: 400;
        margin-top: 0.25rem;
    }

    .text-muted {
        color: #999 !important;
    }

    .form-text.text-muted>a {
        color: #01AAED;
    }

    .card-header {
        border-bottom: 1px solid #eee;
        font-size: 14px;
        color: #333;
        overflow: hidden;
        background-color: #F8F8F8;
        border-radius: 2px 2px 0 0;
    }
</style>

<body>
    <div class="layui-card layadmin-header">
        <span class="layui-breadcrumb" id="breadcrumbTitle" lay-filter="breadcrumb" style="visibility: visible;">
            <a>模板</a>
            <a><cite>图形化配置</cite></a>
        </span>
    </div>
    <div class="layui-fluid">
        <div class="layui-row layui-col-md12 layui-col-space15">
            <div class="layui-col-md8">
                <div class="layui-card">
                    <div class="layui-card-header">配置区域</div>
                    <div class="layui-card-body tourist-box" id="t-container" ondrop="drop(event)"
                        ondragover="allowDrop(event)">
                    </div>
                </div>
            </div>
            <div class="layui-col-md4">
                <div class="layui-card">
                    <div class="layui-card-header">组件列表</div>
                    <div class="layui-card-body">
                        <div class="layui-tab">
                            <ul class="layui-tab-title">
                                <li class="layui-this">接收</li>
                                <li>通道</li>
                                <li>写入</li>
                            </ul>
                            <div class="layui-tab-content tourist-list-box">
                                <!-- tab1 -->
                                <div class="layui-tab-item layui-show">
                                {{range $key, $val := .tourists.SourceMap}}
                                    <a href="javascript:;" class="layui-row tourist-list" id="{{$key}}" data="{{$key}}" onclick="showSource(this)">
                                        <span draggable="true" ondragstart="drag(event)">{{$val}}</span>
                                    </a>
                                {{end}}
                                </div>
                                <!-- tab2 -->
                                <div class="layui-tab-item">
                                {{range $key, $val := .tourists.ChannelMap}}
                                    <a href="javascript:;" class="layui-row tourist-list" id="{{$key}}" data="{{$key}}" onclick="showChannel(this)">
                                        <span draggable="true" ondragstart="drag(event)">{{$val}}</span>
                                    </a>
                                {{end}}
                                </div>
                                <!-- tab3 -->
                                <div class="layui-tab-item">
                                {{range $key, $val := .tourists.SinkMap}}
                                    <a href="javascript:;" class="layui-row tourist-list" id="{{$key}}" data="{{$key}}" onclick="showSink(this)">
                                        <span draggable="true" ondragstart="drag(event)">{{$val}}</span>
                                    </a>
                                {{end}}
                                </div>
                            </div>
                        </div>
                    </div>
                    <button onclick="privew()" class="layui-btn layui-col-md12">保存</button>
                </div>
            </div>
        </div>
    </div>
    {{template "footer3.html" .}}
    <script type="text/javascript" src="/static/assets/js/echarts.min.js"></script>
    <script type="text/javascript" src="/static/js/tourist.js"></script>
</body>

</html>