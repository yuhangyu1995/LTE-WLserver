$(document).ready(function () {
    translate();
    initVersions();
    bindBtnEvent();
    bindToolTip();
    parent.autoHeight()
});

function translate() {
    $("#SoftwareUpdate").text(top.T.VersionManagementPage.SoftwareUpdate);
    $("#AUTOUpload").text(top.T.VersionManagementPage.SoftwareUpdate);
    $("#VersionActive").text(top.T.VersionManagementPage.VersionActive);
    $("#CurrentVersion").text(top.T.VersionManagementPage.CurrentVersion);
    $("#BakVersion").text(top.T.VersionManagementPage.BakVersion)
}

function initVersions() {
    var a = parent.System.getSWpackageVersion();
    if (a.result != "000") {
        return
    }
    if (a.name_val.length < 1) {
        return
    }
    if (a.name_val.length > 0) {
        var b = a.name_val[0].val;
        $("#currentVersion").text((b == "" || b == null) ? "" : b)
    }
    if (a.name_val.length > 1) {
        var b = a.name_val[1].val;
        $("#backupVersion").text((b == "" || b == null) ? "" : b)
    }
}

function bindBtnEvent() {
    $(".btn-upgrade").click(function (a) {
        a.preventDefault();
        submit_upgrade()
    });
    $(".btn-activeversion").click(function (a) {
        a.preventDefault();
        doActive(1)
    });
    $(".btn-minimize").click(function (b) {
        b.preventDefault();
        var a = $(this).parent().parent().next(".box-content");
        if (a.is(":visible")) {
            $("i", $(this)).removeClass("chevron-up").addClass("chevron-down");
            $(this).removeClass("btn-minimize").addClass("btn-maximize")
        } else {
            $("i", $(this)).removeClass("chevron-down").addClass("chevron-up");
            $(this).removeClass("btn-maximize").addClass("btn-minimize")
        }
        bindToolTip();
        a.slideToggle();
        parent.autoHeight()
    });
    $("input:file").not('[data-no-uniform="true"],#uniform-is-ajax').uniform()
}

function bindToolTip() {
    $(".btn-minimize").attr("data-original-title", top.T.BTN.Minimize);
    $(".btn-maximize").attr("data-original-title", top.T.BTN.Maximize);
    $(".btn-upgrade").attr("data-original-title", top.T.VersionManagementPage.SoftwareUpdate);
    $(".btn-activeversion").attr("data-original-title", top.T.VersionManagementPage.ActiveVersion);
    $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
        "placement": "bottom",
        delay: {
            show: 400,
            hide: 200
        }
    })
}

function doActive(a) {
    var b = $("#backupVersion").text();
    if (b == null || b == "") {
        return
    }
    if (a == 2) {
        parent.ModalCommon.open(top.T.AlertInfo.VersionActivationConfirm, top.T.AlertInfo.VersionActivationConfirmInfo + g_version + "?", "choice", doActive)
    } else {
        if (a == 1) {
            top.ModalCommon.progress();
            parent.System.activePackage(b)
        }
    }
}

function submit_upgrade() {
    var a = $("#fileBrowse").val();
    if (a == "") {
        return
    }
    if (!parent.USER.hasModule(parent.Branch.Modules.TD_HOME)) {
        var b = "/req/transferReq.asp";
        var c = "webupdateUpload";
        $("#main_body").hide();
        $.post(b, {
            caseType: 0,
            data: c
        }, function (d) {
            processView()
        })
    } else {
        parent.System.downloadRequest(submit_upload_apply_callback)
    }
}

function submit_upload_apply_callback(b) {
    var a = parent.LMT.parser(b);
    if (a.result != "000") {
        parent.ModalCommon.open(top.T.Main.Notice, a.hint, "error");
        return
    } else {
        var c = "/req/transferReq.asp";
        var d = "webupdateUpload";
        $.post(c, {
            caseType: 0,
            data: d
        }, function (e) {
            processView()
        })
    }
}

function processView() {
    $("#main_body").hide();
    setTimeout(upgradeNow, 500)
}

function upgradeNow() {
    document.UpLoadConfigForm.submit();
    parent.ModalCommon.showSteps([top.T.VersionManagementPage.FirewareUploading]);
    parent.Progress.updateFeller(1)
};