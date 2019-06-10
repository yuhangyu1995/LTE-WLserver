$(document).ready(function() {
    top.MULTI_UI.initial($, "multiMainBox");
    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
});

function bindBtnEvent() {
    $('[data-rel="chosen"],[rel="chosen"]').chosen();
    $('.btn-multisave').click(function(e) {
        e.preventDefault();
        top.MULTI_UI.saveMultiData($, this, saveCallback);
        parent.autoHeight();
    });
    $('.btn-multiadd').click(function(e) {
        e.preventDefault();
        top.MULTI_UI.newMultiData($, this);
        bindBtnEvent2($(this).attr("classId"));
        parent.autoHeight();
    });
    $('.btn-multiminus').click(function(e) {
        e.preventDefault();
        top.MULTI_UI.deleteMultiData($, this, deleteCallback);
        parent.autoHeight();
    });
    $('.btn-multirefresh').click(function(e) {
        e.preventDefault();
        top.MULTI_UI.refreshMultiData($, refreshCallback);
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
        bindToolTip();
        $target.slideToggle();
        parent.autoHeight();
    });
}

function bindBtnEvent2(classId) {
    $('[data-rel="chosen"],[rel="chosen"]').chosen();
    $('.btn-comfirm', $("#newMultiContainer_" + classId)).click(function(e) {
        e.preventDefault();
        top.MULTI_UI.addMultiData($, this, addCallback);
        bindToolTip();
        parent.autoHeight();
    });
    $('.btn-cancel', $("#newMultiContainer_" + classId)).click(function(e) {
        e.preventDefault();
        top.MULTI_UI.cancelMultiData($, this);
        parent.autoHeight();
    });
}

function bindToolTip() {
    $('.btn-minimize').attr("data-original-title", top.DISPLAY.BTN.Minimize);
    $('.btn-maximize').attr("data-original-title", top.DISPLAY.BTN.Maximize);
    $('.btn-multisave').attr("data-original-title", top.DISPLAY.BTN.Save);
    $('.btn-multiadd').attr("data-original-title", top.DISPLAY.BTN.Add);
    $('.btn-multiminus').attr("data-original-title", top.DISPLAY.BTN.Delete);
    $('.btn-multirefresh').attr("data-original-title", top.DISPLAY.BTN.Refresh);
    $('.btn-comfirm').attr("data-original-title", top.DISPLAY.BTN.Confirm);
    $('.btn-cancel').attr("data-original-title", top.DISPLAY.BTN.Cancel);
    $("input:checkbox, input:radio, input:file").not('[data-no-uniform="true"],#uniform-is-ajax').uniform();
    $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
        "placement" : "bottom",
        delay : {
            show : 400,
            hide : 200
        }
    });
}

function refreshCallback() {
    bindBtnEvent();
    bindToolTip();
    parent.autoHeight();
}

function saveCallback() {
    top.MULTI_UI.refreshMultiData($, refreshCallback);
}

function deleteCallback() {
    top.MULTI_UI.refreshMultiData($, refreshCallback);
}

function addCallback() {
    top.MULTI_UI.refreshMultiData($, refreshCallback);
}

function editCallback() {
    //none
}
