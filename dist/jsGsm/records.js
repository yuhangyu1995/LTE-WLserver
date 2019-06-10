var recordTable, timer;

$(document).ready(function() {
    init();
    printRecord();
});

function init() {
    var attribute = {
        tableId : "tablelist",
        tableName : "recordTable",
        tableHeader : [{
            "name" : top.DISPLAY.Records.Index,
            "width" : "10%"
        }, {
            "name" : top.DISPLAY.Records.IMSIInfo,
            "width" : "20%"
        }, {
            "name" : top.DISPLAY.Records.IMEIInfo,
            "width" : "20%"
        }, {
            "name" : top.DISPLAY.Records.TIMEInfo,
            "width" : "26%"
        }]
    }
    recordTable = new TABLE(attribute);

    if ( typeof top.isRefresh === "undefined") {
        top.isRefresh = true;
    }
    if ( typeof top.isMonitor === "undefined") {
        top.isMonitor = false;
    }
    if ( typeof top.filter === "undefined") {
        top.filter = [];
    }

    var value = "";
    for ( i = 0, j = top.filter.length; i < j; i++) {
        value += top.filter[i] + ",";
    }
    $("#IMEI_IMSI").val(value);

    translate();
    bindBtnEvent();

    if (!top.isRefresh) {
        $('i', $('.btn-stoprefresh')).removeClass('off').addClass('retweet');
        $('.btn-stoprefresh').removeClass('btn-stoprefresh').addClass('btn-autorefresh');
    }

    if (top.isMonitor) {
        $('i', $('.btn-monitor-off')).removeClass('play').addClass('stop');
        $('.btn-monitor-off').removeClass('btn-monitor-off').addClass('btn-monitor-on');
    }

    bindToolTip();
}

function printRecord() {
    if (top.isMonitor) {
        showFilterRecord();
        timer = setInterval(function() {
            showFilterRecord();
        }, 10 * 1000);
    } else if (top.isRefresh) {
        showRecord();
        timer = setInterval(function() {
            showRecord();
        }, 10 * 1000);
    } else {
        showRecord();
    }
}

function translate() {
    $("#ImsiLog").text(top.DISPLAY.Records.ImsiLog);
}

function showRecord() {

    var dataArr = getRecord();
    loadRecord(dataArr);

    var html = recordTable.listTable();
    $("#container").empty().append(html);
    recordTable.skip("last");

    parent.autoHeight();
}

function showFilterRecord() {

    var dataArr = getRecord();
    loadFilterRecord(dataArr);

    var html = recordTable.listTable();
    $("#container").empty().append(html);
    recordTable.skip("last");

    parent.autoHeight();
}

function getRecord() {
    var cgiData = {
        cmd : 'getFile',
        filename : '../IMSI.imsi',
        lineNum : '0'
    };
    var record = top.CGI.ajaxAsyn(top.CGI.requesturl.filemanage, cgiData);
    if (top.CGI.isSuccess(record)) {
        return parseRecord(top.CGI.getContent(record));
    } else {
        return [];
    }
}

function parseRecord(record) {
    var dataArr = [];
    var reg = /^ID:(.*) IMSI:(.*) IMEI:(.*) TIME:(\d{14})$/;

    if (record.length > 0) {
        record = record.split("\r\n");
        for (var i = 0, len = record.length; i < len; i++) {
            if (reg.test(record[i])) {
                dataArr.push([RegExp.$1, RegExp.$2, RegExp.$3, timeFormat(RegExp.$4)]);
            }
        }
    }

    return dataArr;
}

function timeFormat(time) {
    return time.substring(0, 4) + "-" + time.substring(4, 6) + "-" + time.substring(6, 8) + " " + //
    time.substring(8, 10) + ":" + time.substring(10, 12) + ":" + time.substring(12);
}

function loadRecord(dataArr) {
    recordTable.empty();
    if (dataArr.length == 0) {
        recordTable.loadData([]);
    } else {
        var tempArr = [];
        for (var i = 0, len = dataArr.length; i < len; i++) {
            for (var j = 0; j < dataArr[i].length; j++) {
                tempArr.push(dataArr[i][j]);
            }
        }
        recordTable.loadData(tempArr);
    }
}

function loadFilterRecord(array) {
    recordTable.empty();
    var dataArr = filterRecord(array);
    if (dataArr.length == 0) {
        recordTable.loadData([]);
    } else {
        var tempArr = [];
        for (var i = 0, len = dataArr.length; i < len; i++) {
            for (var j = 0; j < dataArr[i].length; j++) {
                tempArr.push(dataArr[i][j]);
            }
        }
        recordTable.loadData(tempArr);
    }
}

function filterRecord(dataArr) {
    var array = [];

    var filterStr = $("#IMEI_IMSI").val().replace(/，/g, ",").replace(/\s+/g, "");
    var filterArr = filterStr.split(",");
    var reg = /^\d+$/;
    var index = 1;
    top.filter = [];
    for (var i = 0; i < filterArr.length; i++) {
        if (reg.test(filterArr[i])) {
            top.filter.push(filterArr[i]);
        }
    }

    if (top.filter.length > 0) {
        for (var i = 0; i < dataArr.length; i++) {
            for (var j = 0; j < top.filter.length; j++) {
                if (dataArr[i][1].indexOf(top.filter[j]) != -1 || dataArr[i][2].indexOf(top.filter[j]) != -1) {
                    dataArr[i][0] = index++;
                    array.push(dataArr[i]);
                    break;
                }
            }
        }

    } else {
        array = dataArr;
    }

    return array;
}

function bindBtnEvent() {

    $('.btn-stoprefresh').click(function(e) {
        e.preventDefault();
        if (top.isMonitor) {
            return;
        }

        if (top.isRefresh) {
            top.isRefresh = false;
            clearTimeout(timer);

        } else {
            top.isRefresh = true;
            printRecord();
        }

        if (top.isRefresh) {
            $('i', $(this)).removeClass('retweet').addClass('off');
            $(this).removeClass('btn-autorefresh').addClass('btn-stoprefresh');
        } else {
            $('i', $(this)).removeClass('off').addClass('retweet');
            $(this).removeClass('btn-stoprefresh').addClass('btn-autorefresh');
        }

        bindToolTip();
    });

    $('.btn-refresh').click(function(e) {
        e.preventDefault();
        clearTimeout(timer);
        printRecord();
    });

    $('.btn-trash').click(function(e) {
        e.preventDefault();
        if (top.isMonitor) {
            return;
        }
        
        var cgiData = {
            cmd : "clear_imsi",
            params : ""
        };
    
        top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
            if (top.CGI.isSuccess(data)) {
                clearTimeout(timer);
                printRecord();
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
            }
        });
    });

    $('.btn-monitor-off').click(function(e) {
        e.preventDefault();

        if (top.isMonitor) {
            top.isMonitor = false;
        } else {
            top.isRefresh = false;
            top.isMonitor = true;
        }
        clearTimeout(timer);
        printRecord();

        if (!top.isRefresh) {
            $('i', $('.btn-stoprefresh')).removeClass('off').addClass('retweet');
            $('.btn-stoprefresh').removeClass('btn-stoprefresh').addClass('btn-autorefresh');
        }

        if (top.isMonitor) {
            $('i', $(this)).removeClass('play').addClass('stop');
            $(this).removeClass('btn-monitor-off').addClass('btn-monitor-on');
        } else {
            $('i', $(this)).removeClass('stop').addClass('play');
            $(this).removeClass('btn-monitor-on').addClass('btn-monitor-off');
        }

        bindToolTip();
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
    $("#IMEI_IMSI").keydown(function(e) {
        if (e.keyCode == "13") {
            var btn = $('.btn-monitor-off')[0] ? $('.btn-monitor-off') : $('.btn-monitor-on');
            btn.click();
        }
    });
}

function bindToolTip() {
    $('.btn-autorefresh').attr("data-original-title", top.DISPLAY.Records.AutoRefreshOn);
    $('.btn-stoprefresh').attr("data-original-title", top.DISPLAY.Records.AutoRefreshOff);
    $('.btn-refresh').attr("data-original-title", top.DISPLAY.BTN.Refresh);
    $('.btn-trash').attr("data-original-title", top.DISPLAY.Records.ClearImsi);
    $('.btn-monitor-on').attr("data-original-title", top.DISPLAY.Records.MonitorOff);
    $('.btn-monitor-off').attr("data-original-title", top.DISPLAY.Records.MonitorOn);
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
