var mutex = false;
var cmdQueue = [];
var curCmd = "";

$(document).ready(function() {
    translate();
    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
    
    $("#displayList").empty().append("<p>无执行记录</p>");
    getData();
    setInterval(timerTask, 100);
});

function translate() {
    $("#UploadFileSelect").text(top.DISPLAY.Transfer.UploadFileSelect);
    $("#CmdRes").text(top.DISPLAY.Advance.CmdRes);
    $("#CmdFile").text(top.DISPLAY.Advance.CmdFile);
    $(".btn-conf").text(top.DISPLAY.Advance.Config);
    $(".btn-clear").text(top.DISPLAY.Advance.Clear);
    $(".btn-upload").text(top.DISPLAY.Advance.Upload);
    $(".btn-import").text(top.DISPLAY.Advance.Import);
}

function bindBtnEvent() {
    $('.btn-minimize').click(function(e) {
        e.preventDefault();
        var $target = $(this).parent().parent().next('.box-content');
        if ($target.is(':visible')) {
            $('i', $(this)).removeClass('chevron-up').addClass('chevron-down');
            $(this).removeClass('btn-minimize').addClass('btn-maximize');
        } else {
            $('i', $(this)).removeClass('chevron-down').addClass('chevron-up');
            $(this).removeClass('btn-maximize').addClass('btn-minimize');
        }
        $target.slideToggle();
        bindToolTip();
    });
    $("input:file").not('[data-no-uniform="true"],#uniform-is-ajax').uniform();
}

function bindToolTip() {
    $('.btn-minimize').attr("data-original-title", top.DISPLAY.BTN.Minimize);
    $('.btn-maximize').attr("data-original-title", top.DISPLAY.BTN.Maximize);
    $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
        "placement" : "bottom",
        delay : {
            show : 400,
            hide : 200
        }
    });
}

function getData() {
    top.CGI.ajaxText("../advance.setting", false, function(data) {
        showCmdList(data);
    });

}

function showCmdList(data) {

    $("#cmdList").empty();
    if ("" == $.trim(data).replace(/\r/g, "").replace(/\n/g, "")) {
        $("#cmdList").append("<p>无执行记录</p>");

    } else {
        var cmdArr = data.split("\r\n")
        var start = 0, end = 0, line_html = [];
        for (var i = 0; i < cmdArr.length; i++) {
            if (i == cmdArr.length - 1) {
                end = i;
                while (start <= end) {
                    if ($.trim(cmdArr[end]) != "") {
                        break;
                    }
                    end--;
                }
                if (start <= end) {
                    line_html.push(createLine(cmdArr, "", start, end, true));
                }
            } else if ($.trim(cmdArr[i]) != "" && $.trim(cmdArr[i + 1]) == "") {
                end = i;
                line_html.push(createLine(cmdArr, "", start, end, true));
                start = i + 1;
            }
        }
        $("#cmdList").append(line_html.join(""));
    }
    $("#cmdList").parent().parent().scrollTop($("#cmdList").prop("scrollHeight"));
}

function createLine(cmdArr, cmdRes, start, end, icon_flag) {
    var html = [], icon_class = "";
    if (icon_flag) {
        icon_class = "icon-pencil";
    } else {
        if (top.CGI.isSuccess(cmdRes)) {
            icon_class = "icon-ok";
        } else {
            icon_class = "icon-remove";
        }
    }
    html.push('<li class="left nopadding">');
    html.push(' <div class="timeslot">');
    html.push('     <div class="icon left-1percent"><i class="', icon_class, '"></i></div>');
    html.push('     <span class="message mleft35"><span class="arrow"></span>');
    html.push('     <span class="text">');
    if (!icon_flag) {
        html.push('     <div>');
        html.push('         <div>命令：</div>');
        html.push('         <div>');
        for (var i = start; i <= end; i++) {
            html.push(cmdArr[i].replace(/, /g, ",").replace(/,/g, ", "));
        }
        html.push('         </div>');
        html.push('     </div>');
        
        if(cmdArr[start].indexOf("config") == 0) {
            html.push(composeUpdateRessult(cmdRes));
            
        } else if(cmdArr[start].indexOf("add") == 0) {
            html.push(composeUpdateRessult(cmdRes));
            
        } else if(cmdArr[start].indexOf("delete") == 0) {
            html.push(composeUpdateRessult(cmdRes));
            
        } else if(cmdArr[start].indexOf("query") == 0) {
            html.push(composeQueryResult(top.CGI.getJsonContent(cmdRes).singleData, cmdRes));
            
        } else if(cmdArr[start].indexOf("mix_query") == 0) {
            html.push(composeQueryResult(top.CGI.getJsonContent(cmdRes).mixData, cmdRes));
            
        } else if(cmdArr[start].indexOf("multi_query") == 0) {
            html.push(composeQueryResult(top.CGI.getJsonContent(cmdRes).multiData, cmdRes));
        } else {
        	html.push(composeUpdateRessult(cmdRes));
        }
    } else {
        for (var i = start; i <= end; i++) {
            html.push(cmdArr[i].replace(/, /g, ',').replace(/,/g, ", "));
        }
    }
    html.push('     </span>');
    html.push('     </span>');
    html.push(' </div>');
    html.push('</li>');
    return html.join("");
}

function composeUpdateRessult(cmdRes) {
    var html = [];
	var data = top.CGI.getJsonContent(cmdRes);
	if (typeof (data) != "undefined" && data != "") {
		html.push('<div>');
        html.push(' <div>返回值：</div>');
        html.push(' <div>');

        if ( typeof (data) == "object") {
            for (var key in data) {
                html.push(key + '=' + data[key].replace(/,/g, ", ") + '<br>');
            }
        } else {
            html.push(data.replace(/,/g, ", "));
        }

        html.push(' </div>');
        html.push('</div>');
    } 

    return html.join("");
}

function composeQueryResult(data, cmdRes) {
    
    var html = [];
    if ( typeof (data) != "undefined" && data != "") {
        html.push('<div>');
        html.push(' <div>返回值：</div>');
        html.push(' <div>');

        if ( typeof (data) == "object") {
            for (var j = 0; j < data.length; j++) {
                for (var key in data[j]) {
                    html.push(key + '=' + data[j][key].replace(/,/g, ", ") + '<br>');
                }
            }
        } else {
            for (var j = 0; j < data.length; j++) {
                html.push(data[j].replace(/,/g, ", "));
            }
        }

        html.push(' </div>');
        html.push('</div>');
    }
    
    return html.join("");
}

function autoExeCmd() {
    top.CGI.ajaxText("../advance.setting", false, function(data) {
        var splitCmds = data.split("\r\n");
        for (var i = 0; i < splitCmds.length; i++) {
            if ($.trim(splitCmds[i]) != "" && splitCmds[i].indexOf("#") != 0) {
                cmdQueue.push($.trim(splitCmds[i]));
            }
        }
    });
}

function timerTask() {
    if (cmdQueue.length > 0 && mutex == false) {
        processCmd();
    }
}

function processCmd() {

    mutex = true;
    curCmd = cmdQueue.shift().replace(/[\r\n]+$/, '');
    
    var index = curCmd.indexOf(" ");
    if (index == -1) {
        mutex = false;
        return;
    }
    
    var cmd = curCmd.substring(0, index);
    var params = curCmd.substring(index + 1).replace(/, /g, ",");

    var cgiData = {
        cmd : cmd,
        params : params
    };

    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        var count = $("#cmdCount").text();
        $("#cmdCount").text(++count);

        var line_html = createLine([curCmd], data, 0, 0, false);
        if ("无执行记录" == $("#displayList").text()) {
            $("#displayList").empty().append(line_html);
        } else {
            $("#displayList").append(line_html);
        }
        $("#displayList").parent().parent().scrollTop($("#displayList").prop("scrollHeight"));
        mutex = false;
    }, function failedOption() {
        mutex = false;
    }, function timeout() {
        mutex = false;
    });
}

function manualExeCmd() {
    var cmd = $("#inputCmd").val();
    cmd = cmd.replace(/\r/g, '').replace(/\n/, '');
    if (cmd.indexOf("#") != 0) {
        cmdQueue.push($.trim(cmd));
    }
}

function scriptUpload() {

    var e = document.uploadForm.fileBrowse;
    if (e.value == "") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SelectFile, "error");
        return false;
    }

    $("#uploadPath").val(top.web_path);
    $("#uploadFilename").val(encodeURI($(".filename").text()));

    document.uploadForm.action = top.CGI.requesturl.upload;
    $("#uploadForm").ajaxForm({
        success : uploadResp
    });
    $("#uploadForm").submit();
}

function uploadResp(data) {
    if (top.CGI.isSuccess(data)) {
        renameUploadFile();
    } else {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
    }
}

function renameUploadFile() {
    var cgiData = {
        cmd : "renameAdvanceFile",
        old_file : $("#uploadPath").val() + $(".filename").text(),
        new_file : $("#uploadPath").val() + "advance.setting"
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.webapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            getData();
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
        } else {
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
        }
    });

}

function setZero() {
    $("#cmdCount").text(0);
    $("#displayList").empty().append("<p>无执行记录</p>");
}
