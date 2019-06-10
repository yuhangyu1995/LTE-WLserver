$(document).ready(function() {
    translate();
    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
});

function translate() {
    $("#configImport").text(top.DISPLAY.BKPRef.ParamsImport);
    $("#backupTitle").text(top.DISPLAY.BKPRef.ParamsExportTitle);
    $("#backupBtn").text(top.DISPLAY.BTN.Backup);
}

function bindBtnEvent() {
    $('.btn-import').click(function(e) {
        e.preventDefault();
        refImport();
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
    $('.btn-import').attr("data-original-title", top.DISPLAY.BTN.Import);
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

function refImport() {
    if ($("#fileBrowse").val() != "") {

        $("#uploadPath").val(top.web_path);
        $("#uploadFilename").val($(".filename").text());

        document.uploadForm.action = top.CGI.requesturl.upload;
        
        $("#uploadForm").ajaxForm({
	        success : importResp,
	        uploadProgress: showPercent
	    });
	    $("#uploadForm").submit();
	    top.ModalCommon.progress(top.DISPLAY.Transfer.Uploading + "0%");
        
    } else {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SelectFile, "error");
        return false;
    }
}

function showPercent(event, position, total, percentComplete) {
    if(percentComplete < 100) {
        top.ModalCommon.switchMsg(top.DISPLAY.Transfer.Uploading + percentComplete + "%");
    } else {
        top.ModalCommon.switchMsg(top.DISPLAY.Transfer.Moving);
    }
}

function importResp(data) {
    if (top.CGI.isSuccess(data)) {
    	top.ModalCommon.delProgress();
        notifyOam();
    }
}

function notifyOam() {
	var cgiData = {
        cmd : "import_database",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
        } else {
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
        }
    });
}

function refExport() {
    var cgiData = {
        cmd : "export_database",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            $("#downFilepath").val(top.web_path);
            $("#downFilename").val("database.xml");
		    document.downloadForm.action = top.CGI.requesturl.download;
		    document.downloadForm.submit();
        } else {
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
        }
    });
}