var alarmTable;

$(document).ready(function() {
    
    init();
    showAlarm();
    translate();
    bindBtnEvent();
    bindToolTip();
});

function translate() {
    $("#historyAlarm").text(top.DISPLAY.Alarm.HistoryAlarm);
}

function init() {
    var attribute = {
        tableId : "tablelist",
        tableName : "alarmTable",
        tableHeader : [{
            "name" : top.DISPLAY.Alarm.Name,
            "width" : "15%"
        }, {
            "name" : top.DISPLAY.Alarm.ActiveTime,
            "width" : "15%"
        }, {
            "name" : top.DISPLAY.Alarm.ClearTime,
            "width" : "15%"
        }, {
            "name" : top.DISPLAY.Alarm.Reason,
            "width" : "25%"
        }, {
            "name" : top.DISPLAY.Alarm.ClearReason,
            "width" : "15%"
        }, {
            "name" : top.DISPLAY.Alarm.State,
            "width" : "10%"
        }]
    }
    alarmTable = new TABLE(attribute);
}

function showAlarm() {

    var dataArr = getAlarm();
    loadAlram(dataArr);
    $("#mainBox").empty().append(alarmTable.listTable());
    parent.autoHeight();
}

function getAlarm() {
    
    var cgiData = {
        cmd : 'multi_query',
        params : "[HistoryAlarmInfo]"
    };
    var record = top.CGI.ajaxAsyn(top.CGI.requesturl.oamapp, cgiData);
    if (top.CGI.isSuccess(record)) {
        return parseAlarm(top.CGI.getJsonContent(record));
    }

    return [];
}

function parseAlarm(record) {

    var dataArr = [];

    for (var i = record.multiData.length; i > 0; i--) {
        var state = '<span class="label label-success">已清除</span>'
        dataArr.push([record.multiData[i - 1].AlarmCnName, record.multiData[i - 1].AlarmRaiseTime, record.multiData[i - 1].AlarmClearTime, record.multiData[i - 1].AlarmRaiseCause, record.multiData[i - 1].AlarmClearCause, state]);
    }

    return dataArr;
}

function loadAlram(dataArr) {
    
    alarmTable.empty();

    if (dataArr.length == 0) {
        alarmTable.loadData([]);
    } else {
        var tempArr = [];
        for (var i = 0, len = dataArr.length; i < len; i++) {
            for (var j = 0; j < dataArr[i].length; j++) {
                tempArr.push(dataArr[i][j]);
            }
        }
        alarmTable.loadData(tempArr);
    }
}

function bindBtnEvent() {

    $('.btn-refresh').click(function(e) {
        e.preventDefault();
        reflesh();
        top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
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
    $('.btn-refresh').attr("data-original-title", top.DISPLAY.BTN.Refresh);
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

function reflesh() {
    showAlarm();
}
