$(document).ready(function() {
    parent.autoHeight();
    sysReboot();
});

function sysReboot(msg) {
    if ( typeof (msg) == "undefined") {
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.ResetSystemConfirm, "choice", sysReboot);
        return;
    } else if (window.parseInt(msg) == 1) {
        var cgiData = {
            cmd : "reboot",
            params : ""
        };
        top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
            if (top.CGI.isSuccess(data)) {
                top.ModalCommon.progress(top.DISPLAY.NoticeInfo.Rebooting);
                setTimeout(loopRebootState, 5 * 1000);
            } else {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.Reboot.RebootInvalid, "error");
            }
        });
    }
}

function loopRebootState() {
    var cgiData = {
        cmd : "reboot_state",
        params : ""
    };
    top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
        if (top.CGI.isSuccess(data)) {
            var hint = top.CGI.getHint(data);
            if (hint == '0') {
                top.ModalCommon.delProgress();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.GotoMainPageInfo, "choice", sysReload);
            } else {
                setTimeout(loopRebootState, 5 * 1000);
            }

        } else {
            setTimeout(loopRebootState, 5 * 1000);
        }
    }, function() {
        setTimeout(loopRebootState, 5 * 1000);
    });
}

function sysReload() {
    top.location.reload();
}
