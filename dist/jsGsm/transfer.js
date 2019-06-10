var soureBrowser = true;
var upLoadPath = "";

$(document).ready(function() {
    // if (top.ValidatePwd) {
    $("#password_field").hide();
    showFileOptField();
    // } else {
    // $("#password_field").show();
    // }
    translate();
    bindBtnEvent1();
    bindToolTip();
    parent.autoHeight();
});

function showFileOptField() {

    var html = [];
    html.push('<div class="row-fluid sortable">');
    html.push('	<div class="box span6" id="file_upload_field">');
    html.push('		<div class="box-header">');
    html.push('			<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="FileUpload"></span></h2>');
    html.push('			<div class="box-icon">');
    html.push('				<a href="#" class="btn-browser1" data-rel="tooltip" data-original-title=""><i class="halflings-icon search"></i></a>');
    html.push('				<a href="#" class="btn-upload" data-rel="tooltip" data-original-title=""><i class="halflings-icon upload"></i></a>');
    html.push('				<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');
    html.push('			</div>');
    html.push('		</div>');
    html.push('		<div class="box-content">');
    html.push('			<form class="form-horizontal mbottom0" method="post" name="uploadForm" id="uploadForm" action="" target="hiddenFrame" enctype="multipart/form-data">');
    html.push('				<fieldset>');
    html.push('					<div class="control-group success">');
    html.push('						<label id="UploadFileSelect" class="control-label width-30percent"></label>');
    html.push('						<div class="controls mleft-35percent">');
    html.push('							<div class="uploader">');
    html.push('								<input type="file" id="fileBrowse" name="fileBrowse" onchange="">');
    html.push('							</div>');
    html.push('						</div>');
    html.push('					</div>');
    html.push('					<div class="control-group success">');
    html.push('						<label id="To" class="control-label width-30percent"></label>');
    html.push('						<div class="controls mleft-35percent">');
    html.push('							<span class="uneditable-input" id="uploadFilepath" type="text" title=""></span>');
    html.push('						</div>');
    html.push('					</div>');
    html.push('					<input type="hidden" id="uploadPath" name="uploadPath"/>');
    html.push('					<input type="hidden" id="uploadFilename" name="uploadFilename"/>');
    html.push('				</fieldset>');
    html.push('			</form>');
    html.push('		</div>');
    html.push('	</div>');
    html.push('	<div class="box span6" id="file_download_field">');
    html.push('		<div class="box-header">');
    html.push('			<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="FileDownload"></span></h2>');
    html.push('			<div class="box-icon">');
    html.push('				<a href="#" class="btn-browser2" data-rel="tooltip" data-original-title=""><i class="halflings-icon search"></i></a>');
    html.push('				<a href="#" class="btn-download" data-rel="tooltip" data-original-title=""><i class="halflings-icon download"></i></a>');
    html.push('				<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');
    html.push('			</div>');
    html.push('		</div>');
    html.push('		<div class="box-content height104">');
    html.push('			<form class="form-horizontal" method="post" name="downloadForm" id="downloadForm" action="">');
    html.push('				<fieldset>');
    html.push('					<div class="control-group success">');
    html.push('						<label id="DownloadFileSelect" class="control-label width-30percent"></label>');
    html.push('						<div class="controls mleft-35percent">');
    html.push('							<span class="uneditable-input" id="downFileFullname" type="text"></span>');
    html.push('						</div>');
    html.push('					</div>');
    html.push('					<input id="downFilename" name="downFilename" type="hidden" value="" />');
    html.push('					<input id="downFilepath" name="downFilepath" type="hidden" value="/"/>');
    html.push('				</fieldset>');
    html.push('			</form>');
    html.push('		</div>');
    html.push('	</div>');
    html.push('</div>');
    html.push('<div class="row-fluid sortable">');
    html.push('	<div id="position_first" class="box span6" style="display: none">');
    html.push('		<div class="box-header">');
    html.push('			<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="SelectDir"></span></h2>');
    html.push('			<div class="box-icon">');
    html.push('				<a href="#" class="btn-lastdir" data-rel="tooltip" data-original-title=""><i class="halflings-icon share-alt"></i></a>');
    html.push('				<a href="#" class="btn-selectdir" data-rel="tooltip" data-original-title=""><i class="halflings-icon ok"></i></a>');
    html.push('				<a href="#" class="btn-cancel" data-rel="tooltip" data-original-title=""><i class="halflings-icon remove"></i></a>');
    html.push('			</div>');
    html.push('		</div>');
    html.push('		<div class="box-content x-scroll"></div>');
    html.push('	</div>');
    html.push('	<div id="position_last" class="box span6 pull-right" style="display: none;">');
    html.push('		<div class="box-header">');
    html.push('			<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="SelectFile"></span></h2>');
    html.push('			<div class="box-icon">');
    html.push('				<a href="#" class="btn-lastdir" data-rel="tooltip" data-original-title=""><i class="halflings-icon share-alt"></i></a>');
    html.push('				<a href="#" class="btn-cancel" data-rel="tooltip" data-original-title=""><i class="halflings-icon remove"></i></a>');
    html.push('			</div>');
    html.push('		</div>');
    html.push('		<div class="box-content xy-scroll"></div>');
    html.push('	</div>');
    html.push('</div>');
    html.push('<div id="position_share" class="row-fluid sortable" style="display: none">');
    html.push('	<div class="box span6">');
    html.push('		<div class="box-header">');
    html.push('			<h2><i class="halflings-icon align-justify"></i><span class="break"></span></h2>');
    html.push('			<div class="box-icon"></div>');
    html.push('		</div>');
    html.push('		<div class="box-content x-scroll">');
    html.push('			<table id="folderList" class="table table-striped"></table>');
    html.push('			<input id="curpath" name="curpath" type="hidden" value=""/>');
    html.push('		</div>');
    html.push('	</div>');
    html.push('</div>');

    $("#file_opt_field").empty().append(html.join(""));

    getURL();
    translate();
    bindToolTip();
    bindBtnEvent2();
}

function translate() {
    $("#PasswordValidation").text(top.DISPLAY.Transfer.PasswordValidation);
    $("#InputPassword").text(top.DISPLAY.Transfer.InputPassword);
    $("#FileUpload").text(top.DISPLAY.Transfer.FileUpload);
    $("#UploadFileSelect").text(top.DISPLAY.Transfer.UploadFileSelect);
    $("#To").text(top.DISPLAY.Transfer.To);
    $("#FileDownload").text(top.DISPLAY.Transfer.FileDownload);
    $("#DownloadFileSelect").text(top.DISPLAY.Transfer.DownloadFileSelect);
    $("#SelectDir").text(top.DISPLAY.Transfer.SelectDir);
    $("#SelectFile").text(top.DISPLAY.Transfer.SelectFile);
}

function getURL() {
    var cgiData = {
        cmd : "query_work_dir",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
        	$("#uploadFilepath").text(top.CGI.getContent(data) + "/");
        	$("#uploadFilepath").attr("title", top.CGI.getContent(data) + "/");
        }
    });
}

function bindBtnEvent1() {
    $('.btn-submit').click(function(e) {
        e.preventDefault();
        validatePwd();
        parent.autoHeight();
    });

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
        parent.autoHeight();
    });
}

function bindBtnEvent2() {
    $('.btn-browser1').click(function(e) {
        e.preventDefault();
        showDirect();
        parent.autoHeight();
    });
    $('.btn-upload').click(function(e) {
        e.preventDefault();
        doUpload();
        parent.autoHeight();
    });
    $('.btn-browser2').click(function(e) {
        e.preventDefault();
        browser();
        parent.autoHeight();
    });
    $('.btn-download').click(function(e) {
        e.preventDefault();
        doDownload();
        parent.autoHeight();
    });
    $('.btn-lastdir').click(function(e) {
        e.preventDefault();
        goback();
        parent.autoHeight();
    });
    $('.btn-selectdir').click(function(e) {
        e.preventDefault();
        confirmDirect();
        parent.autoHeight();
    });
    $('.btn-cancel').click(function(e) {
        e.preventDefault();
        cancel();
        parent.autoHeight();
    });
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
        parent.autoHeight();
    });
    $("input:file").not('[data-no-uniform="true"],#uniform-is-ajax').uniform();
}

function bindToolTip() {
    $('.btn-submit').attr("data-original-title", top.DISPLAY.BTN.Submit);
    $('.btn-browser1').attr("data-original-title", top.DISPLAY.BTN.Browser);
    $('.btn-upload').attr("data-original-title", top.DISPLAY.BTN.Upload);
    $('.btn-browser2').attr("data-original-title", top.DISPLAY.BTN.Browser);
    $('.btn-download').attr("data-original-title", top.DISPLAY.BTN.Download);
    $('.btn-lastdir').attr("data-original-title", top.DISPLAY.BTN.Back);
    $('.btn-selectdir').attr("data-original-title", top.DISPLAY.BTN.Confirm);
    $('.btn-cancel').attr("data-original-title", top.DISPLAY.BTN.Cancel);
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

function cancel() {
    $("#position_first").hide();
    $('.box-content', $("#position_first")).empty();
    $("#position_last").hide();
    $('.box-content', $("#position_last")).empty();
}

function browser() {
    getDirList("/");
    var lasthtml = $('.box-content', $("#position_share")).html();
    $('.box-content', $("#position_last")).empty();
    $('.box-content', $("#position_last")).html(lasthtml);
    $("#position_last").show();
    $('.box-content', $("#position_first")).empty();
    $("#position_first").hide();
    soureBrowser = false;
}

function confirmDirect() {
    $("#uploadFilepath").text(upLoadPath);
    $("#uploadFilepath").attr("title", upLoadPath);
    $('.box-content', $("#position_first")).empty();
    $("#position_first").hide();
}

function showDirect() {
    getDirList("/");
    var lasthtml = $('.box-content', $("#position_share")).html();
    $('.box-content', $("#position_first")).empty();
    $('.box-content', $("#position_first")).append(lasthtml);
    $("#position_first").show();
    $('.box-content', $("#position_last")).empty();
    $("#position_last").hide();
    soureBrowser = true;
}

function getDirList(dir) {

    if (dir == null) {
        dir = "/";
    } else {
        dir = dir.replace("\n", "").replace(" ", "");
    }

    var cgiData = {
        cmd : "getDirList",
        dir : dir
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.webdir, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            $("#documentList").empty();
            $("#folderList").empty();
            var list = top.CGI.getContent(data).split('\n');
            list.pop();

            var folders = [];
            var files = [];
            folders.push("/..");

            for (var i = 0; i < list.length; i++) {
                list[i] = list[i].replace(/\?/g, "");
                list[i] = list[i].replace(/(^\s*)|(\s*$)/g, "");
                if (list[i].indexOf("/") != -1) {
                    folders.push(list[i].replace(/\//g, ""));
                } else {
                    files.push(list[i].replace(/\*|@|\||=/, ""));
                }
            }

            if (folders.length > 0) {
                var temp = folders[folders.length - 1];
                if (temp.length >= 1) {
                    var pos = folders[folders.length - 1].indexOf("\n");
                    if (pos != -1 && (pos > (temp.length / 2))) {
                        folders[folders.length - 1] = folders[folders.length - 1].substring(0, pos);
                    }

                    var foldersHtml = [];

                    for (var i = 1; i <= Math.floor(folders.length / 3); i++) {
                        foldersHtml.push('<tr>');
                        for (var j = 1; j <= 3; j++) {
                            foldersHtml.push('<td onclick="selectDirect(this)" class="cursor-pointer">');
                            foldersHtml.push('  <a href="#">');
                            foldersHtml.push('      <i name="dir" class="halflings-icon folder-open"></i>', folders[(i - 1) * 3 + j - 1]);
                            foldersHtml.push('  </a>');
                            foldersHtml.push('</td>');
                        }
                        foldersHtml.push('</tr>');
                    }

                    if (folders.length % 3 != 0) {
                        foldersHtml.push('<tr>');
                        for (var i = 1; i <= (folders.length % 3); i++) {
                            foldersHtml.push('<td onclick="selectDirect(this)" class="cursor-pointer">');
                            foldersHtml.push('  <a href="#">');
                            foldersHtml.push('      <i name="dir" class="halflings-icon folder-open"></i>', folders[folders.length - folders.length % 3 + i - 1]);
                            foldersHtml.push('  </a>');
                            foldersHtml.push('</td>');
                        }
                        for (var i = 1; i <= (3 - folders.length % 3); i++) {
                            foldersHtml.push('<td>&nbsp;</td>');
                        }
                        foldersHtml.push('</tr>');
                    }

                    $("#folderList").append(foldersHtml.join(""));
                }
            }

            if (files.length > 0) {
                var temp = files[files.length - 1];
                if (temp.length >= 1) {
                    var index = files[files.length - 1].indexOf("\n");

                    if (index != -1 && (index > (temp.length / 2))) {
                        files[files.length - 1] = files[files.length - 1].substring(0, index);
                    }

                    var filesHtml = [];

                    for (var i = 1; i <= Math.floor(files.length / 3); i++) {
                        filesHtml.push('<tr>');
                        for (var j = 1; j <= 3; j++) {
                            filesHtml.push('<td onclick="selectDirect(this)" class="cursor-pointer">');
                            filesHtml.push('  <a href="#">');
                            filesHtml.push('      <i name="file" class="icon-file"></i>', files[(i - 1) * 3 + j - 1]);
                            filesHtml.push('  </a>');
                            filesHtml.push('</td>');
                        }
                        filesHtml.push('</tr>');
                    }

                    if (files.length % 3 != 0) {
                        filesHtml.push('<tr>');
                        for (var i = 1; i <= (files.length % 3); i++) {
                            filesHtml.push('<td onclick="selectDirect(this)" class="cursor-pointer">');
                            filesHtml.push('  <a href="#">');
                            filesHtml.push('      <i name="file" class="icon-file"></i>', files[files.length - files.length % 3 + i - 1]);
                            filesHtml.push('  </a>');
                            filesHtml.push('</td>');
                        }
                        for (var i = 1; i <= (3 - files.length % 3); i++) {
                            filesHtml.push('<td>&nbsp;</td>');
                        }
                        filesHtml.push('</tr>');
                    }

                    $("#folderList").append(filesHtml.join(""));
                }
            }
            getCurDir();
            parent.autoHeight();
        }
    });
}

function selectDirect(node) {
    var type = $(node).children().children().attr("name");
    var children = $(node).children();

    if (type == "file") {
        changeFile(children);
    }
    if (type == "dir") {
        changeDir(children);
    }
}

function goback() {
    getDirList("/..");
}

function changeFile(node) {
    if (!soureBrowser) {
        var filename = $.trim($(node).text());
        filename = filename.replace(/\s/g, "");
        filename = filename.replace(/\//g, "");
        $("#downFilename").val(filename);

        var cgiData = {
            cmd : "getCurDir"
        };
        top.CGI.ajaxSyn(top.CGI.requesturl.webdir, cgiData, function(data) {
            if (top.CGI.isSuccess(data)) {
                data = top.CGI.getContent(data);
                if (data == "/") {
                    $("#downFileFullname").text(data + filename);
                    $("#downFileFullname").attr("title", data + filename);
                } else {
                    $("#downFileFullname").text(data + "/" + filename);
                    $("#downFileFullname").attr("title", data + "/" + filename);
                }
            }
        });
        cancel();
    }
}

function changeDir(node) {
    var dir = $.trim($(node).text());
    getDirList(dir);
    $("#downFilename").val("");
}

var filename = "";
function doUpload() {
    if ($("#uploadFilepath").text() == "") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SelectDir, "error");
        return false;
    }

    var e = document.uploadForm.fileBrowse;
    if (e.value == "") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SelectFile, "error");
        return false;
    }

    $("#uploadPath").val($("#uploadFilepath").text());
    $("#uploadFilename").val($(".filename").text());

    document.uploadForm.action = top.CGI.requesturl.upload;
    $("#uploadForm").ajaxForm({
        success : uploadResp,
        uploadProgress: showPercent
    });
    $("#uploadForm").submit();
    top.ModalCommon.progress(top.DISPLAY.Transfer.Uploading + "0%");
}

function uploadResp(data) {
    if (top.CGI.isSuccess(data)) {
        top.ModalCommon.delProgress();
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SoftwareUploadSuccess, "ok");
    }
}

function showPercent(event, position, total, percentComplete) {
    if(percentComplete < 100) {
        top.ModalCommon.switchMsg(top.DISPLAY.Transfer.Uploading + percentComplete + "%");
    } else {
        top.ModalCommon.switchMsg(top.DISPLAY.Transfer.Moving);
    }
}

function doDownload() {
    var filename = $("#downFilename").val();
    if (filename == "") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SelectFile, "error");
    } else {
        document.downloadForm.action = top.CGI.requesturl.download;
        document.downloadForm.submit();
    }
}

function getCurDir() {
    upLoadPath = "";
    var cgiData = {
        cmd : "getCurDir"
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.webdir, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            data = top.CGI.getContent(data);
            $("#curpath").val(data);
            $("#downFilepath").val(data);
            upLoadPath = data + "/";
        }
    });
}

function validatePwd() {

    var password = $("#password").val();
    if (password == "") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.Transfer.PasswordError, "error");
        return;
    }
    var cgiData = {
        cmd : 'validatePassword',
        params : 'password=' + md5(password)
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            $("#password_field").hide();
            showFileOptField();
            translate();
            top.ValidatePwd = true;
            parent.autoHeight();
        } else {
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.Transfer.PasswordError, "error");
        }
    });
}