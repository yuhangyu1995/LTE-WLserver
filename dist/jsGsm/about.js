$(document).ready(function() {

    top.UI.loadCurData();
    top.g_curData.mixDataDisable = true;

    var html = top.UI.initStaticBox();
    $("#mainBox").empty().append(html);
    top.UI.assignValue($);
    $("#mainBox").show();

    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
});

function bindBtnEvent() {
    $('[data-rel="chosen"],[rel="chosen"]').chosen();
    $('.btn-refresh').click(function(e) {
        e.preventDefault();
        fresh();
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
}

function bindToolTip() {
    $('.btn-minimize').attr("data-original-title", top.DISPLAY.BTN.Minimize);
    $('.btn-maximize').attr("data-original-title", top.DISPLAY.BTN.Maximize);
    $('.btn-save').attr("data-original-title", top.DISPLAY.BTN.Save);
    $('.btn-refresh').attr("data-original-title", top.DISPLAY.BTN.Refresh);
    $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
        "placement" : "bottom",
        delay : {
            show : 400,
            hide : 200
        }
    });
}

function fresh() {
    top.UI.loadCurData();
    top.g_curData.mixDataDisable = true;

    var html = top.UI.initStaticBox();
    $("#mainBox").empty().append(html);
    top.UI.assignValue($);
    $("#mainBox").show();

    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
}
