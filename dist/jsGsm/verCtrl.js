$(document).ready(function() {
    translate();
    initVersions();
    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
});

function translate() {
    $("#SoftwareUpdate").text(top.DISPLAY.VerCtrl.SoftwareUpdate);
    $("#FileUpload").text(top.DISPLAY.VerCtrl.SoftwareUpdate);
    $("#VersionActive").text(top.DISPLAY.VerCtrl.VersionActive);
    $("#CurVersion").text(top.DISPLAY.VerCtrl.CurrentVersion);
    $("#BakVersion").text(top.DISPLAY.VerCtrl.BakVersion);
}

function initVersions() {
    var cgiData = {
        cmd : "query_version",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            var content = top.CGI.getJsonContent(data.replace(/\n/g, ""));
            $("#currentVersion").text(content.currentVersion);
            $("#backupVersion").text(content.backupVersion);
        }
    });
}

function bindBtnEvent() {
    $('.btn-upgrade').click(function(e) {
        e.preventDefault();
        sysUpload();
    });
    $('.btn-activeversion').click(function(e) {
        e.preventDefault();
        sysActive();
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
        bindToolTip();
        $target.slideToggle();
        parent.autoHeight();
    });
    $("input:file").not('[data-no-uniform="true"],#uniform-is-ajax').uniform();
}

function bindToolTip() {
    $('.btn-minimize').attr("data-original-title", top.DISPLAY.BTN.Minimize);
    $('.btn-maximize').attr("data-original-title", top.DISPLAY.BTN.Maximize);
    $('.btn-upgrade').attr("data-original-title", top.DISPLAY.VerCtrl.SoftwareUpdate);
    $('.btn-activeversion').attr("data-original-title", top.DISPLAY.VerCtrl.ActiveVersion);
    $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
        "placement" : "bottom",
        delay : {
            show : 400,
            hide : 200
        }
    });
}

function sysActive(msg) {

    var version = $("#backupVersion").text();
    if (version == null || version == '') {
        return;
    }

    if ( typeof (msg) == "undefined") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.VersionActivationConfirmInfo + $("#backupVersion").text() + "?", "choice", sysActive);
        return;
    } else if (window.parseInt(msg) == 1) {
        var cgiData = {
            cmd : "version_switch",
            params : ""
        };
        top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
            if (top.CGI.isSuccess(data)) {
                top.ModalCommon.progress(top.DISPLAY.VerCtrl.VersionActiving);
                setTimeout(loopActiveState, 5 * 1000);
            } else {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.CGI.getContent(data), "error");
            }
        });
    }
}

function loopActiveState() {
    var cgiData = {
        cmd : "version_switch_state",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            var hint = top.CGI.getHint(data);
            if (hint == '0') {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.GotoMainPageInfo, "choice", sysReload);
            } else if (hint == '1') {
                setTimeout(loopUpgradeState, 5 * 1000);
            } else if (hint == '2') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.ActiveStatusError2, "error");
            } else if (hint == '16') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.ActiveStatusError16, "error");
            } else if (hint == '18') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.ActiveStatusError18, "error");
            }

        } else if (data == "") {
            setTimeout(loopActiveState, 5 * 1000);

        }
    }, function failedOption() {

    }, function timeout() {
        top.ModalCommon.switchMsg(top.DISPLAY.VerCtrl.VersionActiveSuccess);
        setTimeout(loopActiveState, 5 * 1000);
    });
}

function sysUpload() {

    var e = document.uploadForm.fileBrowse;
    if (e.value == "") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SelectFile, "error");
        return false;
    }
    $("#uploadFilename").val($(".filename").text());
    $("#uploadPath").val(top.web_path);

    document.uploadForm.action = top.CGI.requesturl.upload;
    $("#uploadForm").ajaxForm({
        success : uploadResp,
        uploadProgress : showPercent
    });

    top.ModalCommon.progress(top.DISPLAY.VerCtrl.FirewareUploading + "0%");
    $("#uploadForm").submit();
}

function uploadResp(data) {
    if (top.CGI.isSuccess(data)) {
        sysUpgrade($(".filename").text());
    } else {
        top.ModalCommon.delProgress();
        top.ModalCommon.open(top.DISPLAY.NoticeInfo.OperationFail);
    }
}

function showPercent(event, position, total, percentComplete) {
    top.ModalCommon.switchMsg(top.DISPLAY.VerCtrl.FirewareUploading + percentComplete + "%");
}

function sysUpgrade(filename) {

    var cgiData = {
        cmd : "upgrade",
        params : "[" + filename + "]"
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            top.ModalCommon.switchMsg(top.DISPLAY.VerCtrl.FirewareChecking);
            setTimeout(loopUpgradeState, 5 * 1000);
        } else {
            top.ModalCommon.delProgress();
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.CmdError, "error");
        }
    });
}

function loopUpgradeState() {
    var cgiData = {
        cmd : "upgrade_state",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            var hint = top.CGI.getHint(data);
            if (hint == '0') {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.GotoMainPageInfo, "choice", sysReload);
            } else if (hint == '1') {
                setTimeout(loopUpgradeState, 5 * 1000);
            } else if (hint == '2') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError2, "error");
            } else if (hint == '3') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError3, "error");
            } else if (hint == '4') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError4, "error");
            } else if (hint == '5') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError5, "error");
            } else if (hint == '6') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError6, "error");
            } else if (hint == '7') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError7, "error");
            } else if (hint == '8') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError8, "error");
            } else if (hint == '9') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError9, "error");
            } else if (hint == '10') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError10, "error");
            } else if (hint == '11') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError11, "error");
            } else if (hint == '12') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError12, "error");
            } else if (hint == '13') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError13, "error");
            } else if (hint == '14') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError14, "error");
            } else if (hint == '15') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError15, "error");
            } else if (hint == '16') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError16, "error");
            } else if (hint == '17') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError17, "error");
            } else if (hint == '18') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.VerCtrl.UpgradeStatusError18, "error");
            }

        } else if (data == "") {
            setTimeout(loopUpgradeState, 5 * 1000);

        }
    }, function failedOption() {

    }, function timeout() {
        top.ModalCommon.switchMsg(top.DISPLAY.VerCtrl.FirewareFinished);
        setTimeout(loopUpgradeState, 5 * 1000);
    });
}

function sysReload() {
    top.ModalCommon.delProgress();
    top.APP.logout();
}