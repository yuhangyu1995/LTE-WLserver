var UI = ( function() {

        function loadCurData() {
            clearMix();
            g_curData.property = g_curMenu.property;
            var params = "";
            for (var i = 0; i < g_curData.property.length; i++) {
                params += getNameById(g_curData.property[i]);
            }
            
            var cgiData = {
                cmd : "mix_query",
                params : params
            };
            
            var data = top.CGI.ajaxAsyn(top.CGI.requesturl.oamapp, cgiData);
            if (top.CGI.isSuccess(data)) {
                parseCurData(top.CGI.getJsonContent(data));
            } else {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
            }
        }
        
        function dynamicLoadCurData() {
            
            var classEnName = "";
            clearMix();
            g_curData.classId = g_curMenu.classId;
            
            for (var i = 0; i < g_database.classes.length; i++) {
                if (g_database.classes[i].classId == g_curMenu.classId) {
                    classEnName = g_database.classes[i].classEnName;
                    for (var j = 0; j < g_database.classes[i].propertys.length; j++) {
                        g_curData.property.push(g_curData.classId + "." + g_database.classes[i].propertys[j].propertyId);
                    }
                    break;
                }
            }
            
            var cgiData = {
                cmd : "query",
                params : "[" + classEnName + "]"
            };
            
            var data = top.CGI.ajaxAsyn(top.CGI.requesturl.oamapp, cgiData);
            if (top.CGI.isSuccess(data)) {
                parseDynamicCurData(top.CGI.getJsonContent(data));
            } else {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
            }
        }
        
        function clearMix() {
            g_curData.property = [];
            g_curData.split = [];
            g_curData.mixData = [];
        }
        
        function getNameById(id) {
            var result = "";
            var classId = id.split(".")[0];
            var propertyId = id.split(".")[1];
            for (var i = 0; i < g_database.classes.length; i++) {
                if (g_database.classes[i].classId == classId) {
                    result = result + "[" + g_database.classes[i].classEnName + "]";
                    for (var j = 0; j < g_database.classes[i].propertys.length; j++) {
                        if (g_database.classes[i].propertys[j].propertyId == propertyId) {
                            result = result + "(" + g_database.classes[i].propertys[j].propertyEnName + ")";
                            break;
                        }
                    }
                    break;
                }
            }
            
            return result;
        }
        
        function parseDynamicCurData(data) {
            composeMixData(data.singleData);
            g_curData.split = g_curMenu.split;
        }

        function parseCurData(data) {
            composeMixData(data.mixData);
            g_curData.split = g_curMenu.split;
        }
        
        function composeMixData(data) {
            
            for(var i = 0; i < data.length; i++) {
                var obj = {};
                var classId = g_curData.property[i].split(".")[0];
                var propertyId = g_curData.property[i].split(".")[1];
                
                for(var j = 0; j < g_database.classes.length; j++) {
                    if(g_database.classes[j].classId == classId) {
                        for(var k = 0; k < g_database.classes[j].propertys.length; k++) {
                            if(g_database.classes[j].propertys[k].propertyId == propertyId) {
                                obj.classId = g_database.classes[j].classId;
                                obj.classEnName = g_database.classes[j].classEnName;
                                obj.classCnName = g_database.classes[j].classCnName;
                                obj.init = g_database.classes[j].init;
                                obj.menu = g_database.classes[j].menu;
                                obj.propertyId = g_database.classes[j].propertys[k].propertyId;
                                obj.propertyEnName = g_database.classes[j].propertys[k].propertyEnName;
                                obj.propertyCnName = g_database.classes[j].propertys[k].propertyCnName;
                                obj.propertyValue = data[i][obj.propertyEnName];
                                obj.propertyType = g_database.classes[j].propertys[k].propertyType;
                                obj.propertyOption = g_database.classes[j].propertys[k].propertyOption;
                                obj.multi = g_database.classes[j].propertys[k].multi;
                                obj["export"] = g_database.classes[j].propertys[k]["export"];
                                obj.validation = g_database.classes[j].propertys[k].validation;
                                obj.tips = g_database.classes[j].propertys[k].tips;
                                obj.setting = g_database.classes[j].propertys[k].setting;
                                obj.change = "false";
                                g_curData.mixData.push(obj);
                                break;
                            }
                        }
                        break;
                    }
                }
            }
        }

        function initStaticBox() {

            var linenum = parseInt(g_curData.split.length / 2) + g_curData.split.length % 2;

            var html = [];
            for (var i = 0, j = 0; j < linenum; i += 2, j++) {
                html.push('<div class="row-fluid sortable">');
                html.push(constructStaticBoxHeader(g_curData.mixData, g_curData.split[i]));
                if (g_curData.split[i + 1]) {
                    html.push(constructStaticBoxHeader(g_curData.mixData, g_curData.split[i + 1]));
                }
                html.push('</div>');
            }

            return html.join("");
        }

        function constructStaticBoxHeader(mixData, split) {

            var html = [];
            html.push('<div class="box span6">');
            html.push(' <div class="box-header">');
            html.push('     <h2><i class="halflings-icon align-justify"></i><span class="break"></span>', split.name, '</h2>');
            html.push('     <div class="box-icon">');
            if (!top.g_curData.mixDataDisable) {
            	html.push('     <a href="#" class="btn-save" data-rel="tooltip" data-original-title=""><i class="halflings-icon ok"></i></a>');
            }
            html.push('         <a href="#" class="btn-refresh" data-rel="tooltip" data-original-title=""><i class="halflings-icon refresh"></i></a>');
            html.push('         <a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');
            html.push('     </div>');
            html.push(' </div>');
            html.push(' <div class="box-content">');
            html.push('     <form class="form-horizontal">');
            html.push('         <fieldset>');
            html.push(constructBoxContent(mixData, split));
            html.push('         </fieldset>');
            html.push('     </form>');
            html.push(' </div>');
            html.push('</div>');

            return html.join("");
        }

        function constructBoxContent(mixData, split) {

            var html = [];

            for (var i = parseInt(split.start); i <= parseInt(split.end); i++) {
                html.push('<div class="control-group success">');
                html.push(' <label for="' + mixData[i].propertyId + '" class="control-label width-35percent">', mixData[i].propertyCnName != "" ? mixData[i].propertyCnName : mixData[i].propertyEnName, '</label>');
                html.push(' <div class="controls mleft-40percent">');
                html.push(constructElement(mixData[i]));
                html.push(' </div>');
                html.push('</div>');
            }

            return html.join("");
        }

        function constructElement(mixElement) {

            var html = [];
            var propertyType = mixElement.propertyType;
            if (g_curData.mixDataDisable) {
                if (propertyType == 'A1') {
                    propertyType = 'A2'
                } else if (propertyType == 'B1') {
                    propertyType = 'B2'
                }
            }

            switch(propertyType) {
                case 'A1': {
                    html.push('<input');
                    html.push(' id="', mixElement.propertyId, '"');
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' validation="', mixElement.validation, '"');
                    html.push(' tips="', mixElement.tips, '"');
                    html.push(' class="border-light-gray"');
                    html.push(' onkeyup="top.UI.markData(this)"');
                    html.push(' onmouseout="top.UI.markData(this)"');
                    html.push('>');

                    break;
                }
                case 'B1': {
                    var options = mixElement.propertyOption.split(",");

                    html.push('<select');
                    html.push(' data-rel="chosen"');
                    html.push(' id="', mixElement.propertyId, '"');
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' onchange="top.UI.markData(this)"');
                    html.push('>');

                    for ( i = 0; i < options.length; i++) {
                        var key_value = options[i].split("=");
                        html.push('<option value="', key_value[1], '">', key_value[0], '</option>');
                    }

                    html.push('</select>');

                    break;
                }
                case 'A2': {
                    html.push('<span');
                    html.push(' class="uneditable-input"');
                    html.push(' id="', mixElement.propertyId, '"');
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push('>');

                    break;
                }
                case 'B2': {
                    var options = mixElement.propertyOption.split(",");

                    html.push('<select');
                    html.push(' id="', mixElement.propertyId, '"');
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' style="display:none"');
                    html.push('>');

                    for ( i = 0; i < options.length; i++) {
                        var key_value = options[i].split("=");
                        html.push('<option value="', key_value[1], '">', key_value[0], '</option>');
                    }

                    html.push('</select>');
                    html.push('<span class="uneditable-input"></span>');

                    break;
                }
                default:
                    break;
            }

            return html.join("");
        }

        function initDynamicBox() {

            var linenum = g_curData.mixData.length;

            var html = [];
            html.push('<div class="row-fluid sortable">');
            html.push(' <div class="box span12">');
            html.push('     <div class="box-header">');
            html.push('         <h2><i class="halflings-icon align-justify"></i><span class="break"></span>', g_curData.mixData[0].classCnName != "" ? g_curData.mixData[0].classCnName : g_curData.mixData[0].classEnName, '</h2>');
            html.push('         <div class="box-icon">');
            html.push('             <a href="#" class="btn-save" data-rel="tooltip" data-original-title=""><i class="halflings-icon ok"></i></a>');
            html.push('             <a href="#" class="btn-refresh" data-rel="tooltip" data-original-title=""><i class="halflings-icon refresh"></i></a>');
            html.push('             <a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');
            html.push('         </div>');
            html.push('     </div>');
            html.push('     <div class="box-content">');
            html.push('         <form class="form-horizontal">');
            html.push('             <fieldset>');

            for (var i = 0; i < linenum; i++) {
                if (g_curData.mixData[i].tips.indexOf("#") != -1) {
                    continue;
                }
                html.push('<div class="control-group success">');
                html.push(' <label for="' + g_curData.mixData[i].propertyId + '" class="control-label width-35percent">', g_curData.mixData[i].propertyCnName != "" ? g_curData.mixData[i].propertyCnName : g_curData.mixData[i].propertyEnName, '</label>');
                html.push(' <div class="controls mleft-40percent">');
                html.push(constructElement(g_curData.mixData[i]));
                html.push(' </div>');
                html.push('</div>');
            }

            html.push('             </fieldset>');
            html.push('         </form>');
            html.push('     </div>');
            html.push(' </div>');
            html.push('</div>');

            return html.join("");
        }

        function markData(element) {

            var classEnName = $(element).attr("name");
            var propertyId = $(element).attr("propertyId");
            var propertyValue = $.trim($(element).val());

            for (var i = 0; i < g_curData.mixData.length; i++) {
                if (g_curData.mixData[i].classEnName == classEnName && g_curData.mixData[i].propertyId == propertyId) {
                    if (g_curData.mixData[i].propertyValue != propertyValue) {
                        g_curData.mixData[i].propertyValue = propertyValue;
                        g_curData.mixData[i].change = "true";
                    }
                    break;
                }
            }
        }

        function saveData(callback) {

            if (!validateData()) {
                return;
            }

            var params = getSaveParam();

            if (params != [] && params.length > 0) {
                doSave(params, 0, callback);
            } else {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SaveSuccess, "ok")
            }
        }
        
        function doSave(params, index, callback) {
            
            if (index == params.length) {
                if ( typeof callback == 'function') {
                    callback();
                }
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SaveSuccess, "ok");
                return;
            }
            
            var cgiData = {
                cmd : "config",
                params : params[index]
            }
            
            top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
                if (top.CGI.isSuccess(data)) {
                    doSave(params, ++index, callback);
                } else {
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error")
                }
            });
        }

        function validateData() {
            for (var i = 0; i < g_curData.mixData.length; i++) {
                if (g_curData.mixData[i].change == "true") {
                    var regex = g_curData.mixData[i].validation;
                    if (regex != "") {
                        if (regex.indexOf("Range:") != -1) {
                            var range = regex.substring(regex.indexOf("Range:") + 6).split("|");
                            if (!rangeValidation(g_curData.mixData[i].propertyValue, range)) {
                                var result = g_curData.mixData[i].tips.indexOf("#") != -1 ? g_curData.mixData[i].tips.substring(g_curData.mixData[i].tips.indexOf("#") + 1) : g_curData.mixData[i].tips;
                                top.ModalCommon.open(top.DISPLAY.Main.Notice, (g_curData.mixData[i].propertyCnName != "" ? g_curData.mixData[i].propertyCnName : g_curData.mixData[i].propertyEnName) + ":" + result, "error");
                                return false;
                            }

                        } else if (regex.indexOf("Regex:") != -1) {
                            if (!eval(regex.substring(regex.indexOf("Regex:") + 6)).test(g_curData.mixData[i].propertyValue)) {
                                var result = g_curData.mixData[i].tips.indexOf("#") != -1 ? g_curData.mixData[i].tips.substring(g_curData.mixData[i].tips.indexOf("#") + 1) : g_curData.mixData[i].tips;
                                top.ModalCommon.open(top.DISPLAY.Main.Notice, (g_curData.mixData[i].propertyCnName != "" ? g_curData.mixData[i].propertyCnName : g_curData.mixData[i].propertyEnName) + ":" + result, "error");
                                return false;
                            }
                        }
                    }
                }
            }

            return true;
        }

        function rangeValidation(value, range) {

            value = window.parseInt(value);
            var reg = /^\(([-]{0,1}\d+),([-]{0,1}\d+)\)$/;

            for (var i = 0; i < range.length; i++) {
                if (reg.test(range[i])) {
                    if (value < window.parseInt(RegExp.$1) || value > window.parseInt(RegExp.$2)) {
                        continue;
                    }
                    return true;
                }
            }

            return false;
        }

        function getSaveParam() {

            var res = [], params = [];
            for (var i = 0; i < g_curData.mixData.length; i++) {
                if (g_curData.mixData[i].change == "true") {
                    res.push(g_curData.mixData[i]);
                }
            }

            if (res.length > 0) {
                res.sort(function(a, b) {
                    return a.classId - b.classId
                });

                var start = 0;
                var lastClassEnName = res[0].classEnName;
                for (var i = 0; i < res.length; i++) {
                    if (i == (res.length - 1)) {
                        if (res[i].classEnName != lastClassEnName) {
                            params.push(joinParam(res, start, i - 1));
                            params.push(joinParam(res, i, i));
                        } else {
                            params.push(joinParam(res, start, i));
                        }
                    } else if (res[i].classEnName != lastClassEnName) {
                        params.push(joinParam(res, start, i - 1));
                        lastClassEnName = res[i].classEnName;
                        start = i;
                    }
                }
            }

            return params;
        }

        function joinParam(arr, start, end) {

            var param = "[" + arr[start].classEnName + "](";
            for (var i = start; i <= end; i++) {
                param = param + arr[i].propertyEnName + "=\"" + arr[i].propertyValue + "\",";
            }
            param = param.substring(0, param.length - 1);
            param = param + ")";

            return param;
        }

        function assignValue($) {
            for (var i = 0; i < g_curData.mixData.length; i++) {
                var element = $("#" + g_curData.mixData[i].propertyId);
                var propertyType = g_curData.mixData[i].propertyType;
                if (g_curData.mixDataDisable) {
                    if (propertyType == 'A1') {
                        propertyType = 'A2'
                    } else if (propertyType == 'B1') {
                        propertyType = 'B2'
                    }
                }

                switch(propertyType) {
                    case 'A1': {
                        $(element).val(g_curData.mixData[i].propertyValue);
                        break;
                    }
                    case 'B1': {
                        $(element).val(g_curData.mixData[i].propertyValue);
                        break;
                    }
                    case 'A2': {
                        $(element).text(g_curData.mixData[i].propertyValue);
                        $(element).attr("title", g_curData.mixData[i].propertyValue);
                        break;
                    }
                    case 'B2': {
                        $(element).val(g_curData.mixData[i].propertyValue);

                        var options = $(element).children('option');
                        for (var j = 0; j < options.length; j++) {
                            if ($(options[j]).attr("value") == g_curData.mixData[i].propertyValue) {
                                $(element).next().text($(options[j]).html());
                                $(element).next().attr("title", $(options[j]).html());
                                break;
                            }
                        }
                        break;
                    }
                    default:
                        break;
                }
            }
        }

        return {
            loadCurData : loadCurData,
            initStaticBox : initStaticBox,
            initDynamicBox : initDynamicBox,
            markData : markData,
            saveData : saveData,
            assignValue : assignValue,
            dynamicLoadCurData : dynamicLoadCurData
        }

    }());

var MULTI_UI = ( function() {

        var containerId = "";

        function initial($, id) {
            containerId = id;
            clearMulti();
            
            if(!g_curMenu.multiId){
                return;
            }
            
            loadCurData();
            $("#" + id).empty().append(initMultiBox());
            assignMultiValue($);
            $("#" + id).show();
            initMainBoxContentHeight($);
        }
        
        function clearMulti() {
            g_curData.multiId = [];
            g_curData.multiData = [];
            g_curData.multiDefaultData = [];
        }

        function loadCurData() {
            g_curData.multiId = g_curMenu.multiId;
            
            var params = "";
            for (var i = 0; i < g_curData.multiId.length; i++) {
                params = getNameById(g_curData.multiId[i]);
            
            
                var cgiData = {
                    cmd : "multi_query",
                    params : params
                };
                
                var data = top.CGI.ajaxAsyn(top.CGI.requesturl.oamapp, cgiData);
                if (top.CGI.isSuccess(data)) {
                    parseCurData(top.CGI.getJsonContent(data), g_curData.multiId[i]);
                } else {
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
                }
            }
        }
        
        function getNameById(id){
            var result = "";
            var classId = id;
            for (var i = 0; i < g_database.classes.length; i++) {
                if (g_database.classes[i].classId == classId) {
                    result = "[" + g_database.classes[i].classEnName + "]";
                    break;
                }
            }
            
            return result;
        }

        function parseCurData(data, classId) {
            composeMultiData(data.multiData, classId);
            composeMultiDefaultData(classId);
        }
        
        function composeMultiData(data, classId) {
            
            for(var i = 0; i < g_database.classes.length; i++) {
                if(g_database.classes[i].classId == classId) {
                    var record = {};
                    for(var j = 0; j < data.length; j++) {
                        var key = classId + "_" + j;
                        record[key] = [];
                        for(var k = 0; k < g_database.classes[i].propertys.length; k++) {
                            var obj = {};
                            obj.classId = g_database.classes[i].classId;
                            obj.classEnName = g_database.classes[i].classEnName;
                            obj.classCnName = g_database.classes[i].classCnName;
                            obj.init = g_database.classes[i].init;
                            obj.menu = g_database.classes[i].menu;
                            obj.propertyId = g_database.classes[i].propertys[k].propertyId;
                            obj.propertyEnName = g_database.classes[i].propertys[k].propertyEnName;
                            obj.propertyCnName = g_database.classes[i].propertys[k].propertyCnName;
                            obj.propertyValue = data[j][obj.propertyEnName];
                            obj.propertyType = g_database.classes[i].propertys[k].propertyType;
                            obj.propertyOption = g_database.classes[i].propertys[k].propertyOption;
                            obj.multi = g_database.classes[i].propertys[k].multi;
                            obj["export"] = g_database.classes[i].propertys[k]["export"];
                            obj.validation = g_database.classes[i].propertys[k].validation;
                            obj.tips = g_database.classes[i].propertys[k].tips;
                            obj.setting = g_database.classes[i].propertys[k].setting;
                            obj.change = "false";
                            record[key].push(obj);
                        }
                    }
                    if(record && data.length > 0) {
                        g_curData.multiData.push(record);
                    }
                    break;
                }
            }
        }
        
        function composeMultiDefaultData(classId) {
            
            for(var i = 0; i < g_database.classes.length; i++) {
                if(g_database.classes[i].classId == classId) {
                    var record = {};
                    record[classId] = [];
                    for(var k = 0; k < g_database.classes[i].propertys.length; k++) {
                        var obj = {};
                        obj.classId = g_database.classes[i].classId;
                        obj.classEnName = g_database.classes[i].classEnName;
                        obj.classCnName = g_database.classes[i].classCnName;
                        obj.init = g_database.classes[i].init;
                        obj.menu = g_database.classes[i].menu;
                        obj.propertyId = g_database.classes[i].propertys[k].propertyId;
                        obj.propertyEnName = g_database.classes[i].propertys[k].propertyEnName;
                        obj.propertyCnName = g_database.classes[i].propertys[k].propertyCnName;
                        obj.propertyValue = g_database.classes[i].propertys[k].propertyValue;
                        obj.propertyType = g_database.classes[i].propertys[k].propertyType;
                        obj.propertyOption = g_database.classes[i].propertys[k].propertyOption;
                        obj.multi = g_database.classes[i].propertys[k].multi;
                        obj["export"] = g_database.classes[i].propertys[k]["export"];
                        obj.validation = g_database.classes[i].propertys[k].validation;
                        obj.tips = g_database.classes[i].propertys[k].tips;
                        obj.setting = g_database.classes[i].propertys[k].setting;
                        obj.change = "false";
                        record[classId].push(obj);
                    }
                    g_curData.multiDefaultData.push(record);
                    break;
                }
            }
        }

        function initMultiBox() {

            var html = [];

            for (var i = 0; i < g_curData.multiId.length; i++) {
                var classId = "", classEnName = "", classCnName = "";
                for (var j = 0; j < g_database.classes.length; j++) {
                    if (g_database.classes[j].classId == g_curData.multiId[i]) {
                        classId = g_database.classes[j].classId;
                        classEnName = g_database.classes[j].classEnName;
                        classCnName = g_database.classes[j].classCnName;
                        break;
                    }
                }

                html.push('<div class="row-fluid sortable">');
                html.push(' <div class="box span12">');
                html.push('     <div class="box-header" data-original-title>');
                html.push('         <h2><i class="halflings-icon align-justify"></i><span class="break"></span>', classCnName != "" ? classCnName : classEnName, '</h2>');
                html.push('         <div class="box-icon">');
                html.push('             <a href="#" class="btn-multisave" classId="', classId, '" name="' + classEnName + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon ok"></i></a>');
                html.push('             <a href="#" class="btn-multiadd" classId="', classId, '" name="' + classEnName + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon plus"></i></a>');
                html.push('             <a href="#" class="btn-multiminus" classId="', classId, '" name="' + classEnName + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon minus"></i></a>');
                html.push('             <a href="#" class="btn-multirefresh" classId="', classId, '" name="' + classEnName + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon refresh"></i></a>');
                html.push('             <a href="#" class="btn-minimize" classId="', classId, '" name="' + classEnName + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');
                html.push('         </div>');
                html.push('     </div>');
                html.push('     <div class="box-content xy-scroll">');
                html.push('         <table class="table table-bordered" id="instanceBox_', classId, '">');
                html.push(constructMultiContent(g_curData.multiData, classId));
                html.push('         </table>');
                html.push('     </div>');
                html.push(' </div>');
                html.push('</div>');

                html.push(constructAddBoxHeader(classId));
            }

            return html.join("");
        }

        function constructMultiContent(multiData, classId) {

            var index = 0;

            var html = [];
            html.push('<tr>');
            html.push('<th nowrap="nowrap"><input type="checkbox" classId="', classId, '" onclick="top.MULTI_UI.selectAll($, this);"/></th>');

            for (var i = 0; i < g_database.classes.length; i++) {
                if (g_database.classes[i].classId == classId) {
                    for (var j = 0; j < g_database.classes[i].propertys.length; j++) {
                        if (g_database.classes[i].propertys[j].tips.indexOf("#") != -1) {
                            continue;
                        }
                        html.push("<th nowrap='nowrap'><font size='2px'>", g_database.classes[i].propertys[j].propertyCnName != "" ? g_database.classes[i].propertys[j].propertyCnName : g_database.classes[i].propertys[j].propertyEnName, "</font></th>");
                    }
                }
            }
            
            html.push('</tr>');

			for (var i = 0; i < multiData.length; i++) {
				for (var key in multiData[i]) {
					if (key.split("_")[0] != classId) {
						continue;
					}
					html.push('<tr>');
					html.push(' <td>');
					html.push('     <input type="checkbox" class="', classId, '_checklist" index="', index, '">');
					html.push(' </td>');
					for (var j = 0; j < multiData[i][key].length; j++) {
						if (multiData[i][key][j].tips.indexOf("#") != -1) {
							continue;
						}

						html.push('<td>');
						html.push(constructMultiElement(multiData[i][key][j], index));
						html.push('</td>');
					}
					html.push('</tr>');
					index++;
				}
			}
            
            return html.join("");
        }

        function constructAddBoxHeader(classId) {

            var className = "";
            for (var i = 0; i < g_database.classes.length; i++) {
                if (g_database.classes[i].classId == classId) {
                    className = g_database.classes[i].classEnName;
                    break;
                }
            }

            var html = [];
            html.push('<div class="row-fluid sortable" style="display:none;" id="newMultiContainer_', classId, '">');
            html.push(' <div class="box span12">');
            html.push('     <div class="box-header" data-original-title>');
            html.push('         <h2><i class="halflings-icon align-justify"></i><span class="break"></span>', top.DISPLAY.BTN.Add, '</h2>');
            html.push('         <div class="box-icon">');
            html.push('             <a href="#" class="btn-comfirm" classId="', classId, '" name="' + className + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon ok"></i></a>');
            html.push('             <a href="#" class="btn-cancel" classId="', classId, '" name="' + className + '" data-rel="tooltip" data-original-title=""><i class="halflings-icon remove"></i></a>');
            html.push('         </div>');
            html.push('     </div>');
            html.push('     <div class="box-content xy-scroll">');
            html.push('         <table class="table table-bordered" id="newMultiContent_', classId, '">');
            html.push('         </table>');
            html.push('     </div>');
            html.push(' </div>');
            html.push('</div>');

            return html.join("");
        }

        function constructMultiElement(mixElement, index) {

            var html = [];

            switch(mixElement.propertyType) {
                case 'A1': {
                    html.push('<input');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' validation="', mixElement.validation, '"');
                    html.push(' tips="', mixElement.tips, '"');
                    html.push(' class="border-light-gray"');
                    html.push(' onkeyup="top.MULTI_UI.markMultiData(this)"');
                    html.push(' onmouseout="top.MULTI_UI.markMultiData(this)"');
                    html.push(' onblur="editCallback(this);"');
                    html.push('>');

                    break;
                }
                case 'B1': {
                    var options = mixElement.propertyOption.split(",");

                    html.push('<select');
                    html.push(' data-rel="chosen"');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' onchange="top.MULTI_UI.markMultiData(this);editCallback(this);"');
                    html.push('>');

                    for ( i = 0; i < options.length; i++) {
                        var key_value = options[i].split("=");
                        html.push('<option value="', key_value[1], '">', key_value[0], '</option>');
                    }

                    html.push('</select>');

                    break;
                }
                case 'A2': {
                    html.push('<span');
                    html.push(' class="uneditable-input"');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push('>');

                    break;
                }
                case 'B2': {
                    var options = mixElement.propertyOption.split(",");

                    html.push('<select');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' style="display:none"');
                    html.push('>');

                    for ( i = 0; i < options.length; i++) {
                        var key_value = options[i].split("=");
                        html.push('<option value="', key_value[1], '">', key_value[0], '</option>');
                    }

                    html.push('</select>');
                    html.push('<span class="uneditable-input"></span>');

                    break;
                }
                default:
                    break;
            }

            return html.join("");
        }

        function markMultiData(element) {

            var index = $(element).attr("index");
            var classId = $(element).attr("classId");
            var classEnName = $(element).attr("name");
            var propertyId = $(element).attr("propertyId");
            var propertyValue = $.trim($(element).val());

            for (var i = 0; i < g_curData.multiData.length; i++) {
                for (var key in g_curData.multiData[i]) {
                    if (key == classId + "_" + index) {
                        var obj = g_curData.multiData[i][key];
                        for (var j = 0; j < obj.length; j++) {
                            if (obj[j].classEnName == classEnName && obj[j].propertyId == propertyId) {
                                if (obj[j].propertyValue != propertyValue) {
                                    obj[j].propertyValue = propertyValue;
                                    obj[j].change = "true";
                                }
                                break;
                            }
                        }
                    }
                }
            }
        }

        function assignMultiValue($) {
            for (var i = 0; i < g_curData.multiData.length; i++) {
                var index = 0;
                for (var key in g_curData.multiData[i]) {
                    for (var j = 0; j < g_curData.multiData[i][key].length; j++) {
                        var element = $("#" + g_curData.multiData[i][key][j].propertyId + "_" + index);
                        var propertyType = g_curData.multiData[i][key][j].propertyType;
                        var propertyValue = g_curData.multiData[i][key][j].propertyValue;

                        switch(propertyType) {
                            case 'A1': {
                                $(element).val(propertyValue);
                                break;
                            }
                            case 'B1': {
                                $(element).val(propertyValue);
                                break;
                            }
                            case 'A2': {
                                $(element).text(propertyValue);
                                $(element).attr("title", propertyValue);
                                break;
                            }
                            case 'B2': {
                                $(element).val(propertyValue);

                                var options = $(element).children('option');
                                for (var t = 0; t < options.length; t++) {
                                    if ($(options[t]).attr("value") == propertyValue) {
                                        $(element).next().text($(options[t]).html());
                                        $(element).next().attr("title", $(options[t]).html());
                                        break;
                                    }
                                }
                                break;
                            }
                            default:
                                break;
                        }
                    }
                    index++;
                }
            }
        }

        function assignAddMultiValue($, element) {

            var classId = $(element).attr('classId');
            for (var i = 0; i < g_curData.multiDefaultData.length; i++) {
                for (var key in g_curData.multiDefaultData[i]) {
                    if (key == classId) {
                        for (var j = 0; j < g_curData.multiDefaultData[i][key].length; j++) {
                            var element = $("#" + g_curData.multiDefaultData[i][key][j].propertyId + "_new");
                            var propertyType = g_curData.multiDefaultData[i][key][j].propertyType;
                            var propertyValue = g_curData.multiDefaultData[i][key][j].propertyValue;

                            switch(propertyType) {
                                case 'A1': {
                                    $(element).val(propertyValue);
                                    break;
                                }
                                case 'B1': {
                                    $(element).val(propertyValue);
                                    break;
                                }
                                case 'A2': {
                                    $(element).text(propertyValue);
                                    $(element).attr("title", propertyValue);
                                    break;
                                }
                                case 'B2': {
                                    $(element).val(propertyValue);

                                    var options = $(element).children('option');
                                    for (var t = 0; t < options.length; t++) {
                                        if ($(options[t]).attr("value") == propertyValue) {
                                            $(element).next().text($(options[t]).html());
                                            $(element).next().attr("title", $(options[t]).html());
                                            break;
                                        }
                                    }
                                    break;
                                }
                                default:
                                    break;
                            }
                        }
                    }
                }
            }
        }

        function newMultiData($, element) {

            var classId = $(element).attr('classId');
            var defaultData = null;

            for (var i = 0; i < g_curData.multiDefaultData.length; i++) {
                for (var key in g_curData.multiDefaultData[i]) {
                    if (key == classId) {
                        defaultData = g_curData.multiDefaultData[i][key];
                        break;
                    }
                }
            }

            var html = [];
            html.push('<tr>');
            for (var i = 0; i < defaultData.length; i++) {
                if (defaultData[i].tips.indexOf("#") != -1) {
                    continue;
                }
                html.push("<th nowrap='nowrap'><font size='2px'>", defaultData[i].propertyCnName != "" ? defaultData[i].propertyCnName : defaultData[i].propertyEnName, "</font></th>");
            }
            html.push('</tr>');

            html.push('<tr>');
            for (var i = 0; i < defaultData.length; i++) {
                if (defaultData[i].tips.indexOf("#") != -1) {
                    continue;
                }
                if (defaultData[i].propertyType == "A2") {
                    defaultData[i].propertyType = "A1";
                } else if (defaultData[i].propertyType == "B2") {
                    defaultData[i].propertyType = "B1";
                }
                
                html.push(' <td>');
                html.push(constructAddMultiElement(defaultData[i], "new"));
                html.push(' </td>');
            }
            html.push('</tr>');

            $("#newMultiContent_" + classId).empty().append(html.join(""));
            assignAddMultiValue($, element);
            $("#newMultiContainer_" + classId).show();
            initAddBoxContentHeight($, element);
            parent.autoHeight();
        }

        function constructAddMultiElement(mixElement, index) {

            var html = [];

            switch(mixElement.propertyType) {
                case 'A1': {
                    html.push('<input');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' validation="', mixElement.validation, '"');
                    html.push(' tips="', mixElement.tips, '"');
                    html.push(' class="border-light-gray"');
                    html.push(' onblur="editCallback(this);"');
                    html.push('>');

                    break;
                }
                case 'B1': {
                    var options = mixElement.propertyOption.split(",");

                    html.push('<select');
                    html.push(' data-rel="chosen"');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' onchange="editCallback(this);"');
                    html.push('>');

                    for ( i = 0; i < options.length; i++) {
                        var key_value = options[i].split("=");
                        html.push('<option value="', key_value[1], '">', key_value[0], '</option>');
                    }

                    html.push('</select>');

                    break;
                }
                case 'A2': {
                    html.push('<span');
                    html.push(' class="uneditable-input"');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push('>');

                    break;
                }
                case 'B2': {
                    var options = mixElement.propertyOption.split(",");

                    html.push('<select');
                    html.push(' id="', mixElement.propertyId, '_', index, '"');
                    html.push(' index=', index);
                    html.push(' classId="', mixElement.classId, '"');
                    html.push(' name="', mixElement.classEnName, '"');
                    html.push(' propertyId="', mixElement.propertyId, '"');
                    html.push(' propertyName="', mixElement.propertyEnName, '"');
                    html.push(' propertyType="', mixElement.propertyType, '"');
                    html.push(' type="text"');
                    html.push(' style="display:none"');
                    html.push('>');

                    for ( i = 0; i < options.length; i++) {
                        var key_value = options[i].split("=");
                        html.push('<option value="', key_value[1], '">', key_value[0], '</option>');
                    }

                    html.push('</select>');
                    html.push('<span class="uneditable-input"></span>');

                    break;
                }
                default:
                    break;
            }

            return html.join("");
        }

        function addMultiData($, element, callback) {

            if (!validateMultiDefaultData($, g_curData.multiDefaultData)) {
                return;
            }

            var classId = $(element).attr("classId");
            var className = $(element).attr("name");
            var params = "";
            for (var i = 0; i < g_curData.multiDefaultData.length; i++) {
                if (g_curData.multiDefaultData[i][classId]) {
                    for (var j = 0; j < g_curData.multiDefaultData[i][classId].length; j++) {
                        if (g_curData.multiDefaultData[i][classId][j].tips.indexOf("#") != -1) {
                            continue;
                        }
                        var propertyId = g_curData.multiDefaultData[i][classId][j].propertyId;
                        var propertyEnName = g_curData.multiDefaultData[i][classId][j].propertyEnName;
                        params = params + propertyEnName + "=\"" + $("#" + propertyId + "_new").val() + "\",";
                    }
                }
            }
            params = "[" + className + "](" + params.substring(0, params.length - 1) + ")";

            var cgiData = {
                cmd : "add",
                params : params
            }

            top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
                if (top.CGI.isSuccess(data)) {
                    if ( typeof callback == 'function') {
                        callback();
                    }
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SaveSuccess, "ok");
                } else {
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
                }
            });
        }

        function saveMultiData($, element, callback) {

            if (!validateMultiData($, g_curData.multiData)) {
                return;
            }

            var classId = $(element).attr("classId");
            var className = $(element).attr("name");

            var params = [];
            for (var i = 0; i < g_curData.multiData.length; i++) {
                for (var key in g_curData.multiData[i]) {
                    var temp = "";
                    var pk = "";
                    for (var j = 0; j < g_curData.multiData[i][key].length; j++) {
                        if (g_curData.multiData[i][key][0].classId == classId) {
                            if(g_curData.multiData[i][key][j].multi == "1") {
                                if(pk == ""){
                                    pk = g_curData.multiData[i][key][j].propertyEnName + "=\"" + g_curData.multiData[i][key][j].propertyValue + "\"";
                                }else{
                                    pk = pk + "," + g_curData.multiData[i][key][j].propertyEnName + "=\"" + g_curData.multiData[i][key][j].propertyValue + "\"";
                                }
                            }
                            if (g_curData.multiData[i][key][j].change == "true") {
                                temp = temp + g_curData.multiData[i][key][j].propertyEnName + "=\"" + g_curData.multiData[i][key][j].propertyValue + "\",";
                            }
                        }
                    }
                    if (temp != "") {
                        params.push("[" + className + "](" + temp + pk + ")");
                    }
                }
            }

            if (params == []) {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SaveSuccess, "ok");
                return false;
            }

            doMultiSave(params, 0, callback);
        }
        
        function doMultiSave(params, index, callback){
            
            if (index == params.length) {
                if ( typeof callback == 'function') {
                    callback();
                }
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.SaveSuccess, "ok");
                return;
            }
            
            var cgiData = {
                cmd : "config",
                params : params[index]
            }
            
            top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
                if (top.CGI.isSuccess(data)) {
                    doMultiSave(params, ++index, callback);
                } else {
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
                }
            });
        }

        function validateMultiData($, multiData) {
            for (var i = 0; i < multiData.length; i++) {
                for (var key in multiData[i]) {
                    for (var j = 0; j < multiData[i][key].length; j++) {
                        if (multiData[i][key][j].change == "true") {
                            var regex = multiData[i][key][j].validation;
                            var multiObj = multiData[i][key][j];
                            if (regex != "") {
                                if (regex.indexOf("Range:") != -1) {
                                    var range = regex.substring(regex.indexOf("Range:") + 6).split("|");
                                    if (!rangeValidation(multiObj.propertyValue, range)) {
                                        top.ModalCommon.open(top.DISPLAY.Main.Notice, (multiObj.propertyCnName != "" ? multiObj.propertyCnName : multiObj.propertyEnName) + ":" + multiObj.tips, "error");
                                        return false;
                                    }

                                } else if (regex.indexOf("Regex:") != -1) {
                                    if (!eval(regex.substring(regex.indexOf("Regex:") + 6)).test(multiObj.propertyValue)) {
                                        top.ModalCommon.open(top.DISPLAY.Main.Notice, (multiObj.propertyCnName != "" ? multiObj.propertyCnName : multiObj.propertyEnName) + ":" + multiObj.tips, "error")
                                        return false;
                                    }
                                }
                            }
                        }
                    }
                }
            }

            return true;
        }

        function validateMultiDefaultData($, multiData) {
            
            for (var i = 0; i < multiData.length; i++) {
                for (var key in multiData[i]) {
                    for (var j = 0; j < multiData[i][key].length; j++) {
                        if (multiData[i][key][j].tips.indexOf("#") == -1) {
                            var regex = multiData[i][key][j].validation;
                            var multiObj = multiData[i][key][j];
                            if (regex != "") {
                                if (regex.indexOf("Range:") != -1) {
                                    var range = regex.substring(regex.indexOf("Range:") + 6).split("|");
                                    if (!rangeValidation($("#" + multiObj.propertyId + "_new").val(), range)) {
                                        top.ModalCommon.open(top.DISPLAY.Main.Notice, (multiObj.propertyCnName != "" ? multiObj.propertyCnName : multiObj.propertyEnName) + ":" + multiObj.tips, "error");
                                        return false;
                                    }

                                } else if (regex.indexOf("Regex:") != -1) {
                                    if (!eval(regex.substring(regex.indexOf("Regex:") + 6)).test($("#" + multiObj.propertyId + "_new").val())) {
                                        top.ModalCommon.open(top.DISPLAY.Main.Notice, (multiObj.propertyCnName != "" ? multiObj.propertyCnName : multiObj.propertyEnName) + ":" + multiObj.tips, "error")
                                        return false;
                                    }
                                }
                            }
                        }
                    }
                }
            }

            return true;
        }

        function rangeValidation(value, range) {

            value = window.parseInt(value);
            var reg = /^\(([-]{0,1}\d+),([-]{0,1}\d+)\)$/;

            for (var i = 0; i < range.length; i++) {
                if (reg.test(range[i])) {
                    if (value < window.parseInt(RegExp.$1) || value > window.parseInt(RegExp.$2)) {
                        continue;
                    }
                    return true;
                }
            }

            return false;
        }

        function change(multiData) {
            for (var i = 0; i < multiData.length; i++) {
                for (var key in multiData[i]) {
                    for (var j = 0; j < multiData[i][key].length; j++) {
                        multiData[i][key][j].change = "true";
                    }
                }
            }
        }

        function deleteMultiData($, element, callback) {
            var classId = $(element).attr("classId");
            var checklist = $("." + classId + "_checklist");
            var className = $(element).attr("name");

            var params = "";
            
            for (var k = 0; k < checklist.length; k++) {
                if ($(checklist[k]).parent().hasClass("checked")) {
                    for (var i = 0; i < g_curData.multiData.length; i++) {
                        for (var key in g_curData.multiData[i]) {
                            if (key.split("_")[0] == classId && key.split("_")[1] == $(checklist[k]).attr("index")) {
                                for (var j = 0; j < g_curData.multiData[i][key].length; j++) {
                                    if(g_curData.multiData[i][key][j].multi == "1") {
                                        params = params + g_curData.multiData[i][key][j].propertyEnName + "=\"" + g_curData.multiData[i][key][j].propertyValue + "\",";
                                    }
                                }
                            }
                        }
                    }
                }
            }
            
            if (params == "") {
                return false;
            }

            params = params.substring(0, params.length - 1);
            params = "[" + className + "](" + params + ")";

            var cgiData = {
                cmd : "delete",
                params : params
            }

            top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
                if (top.CGI.isSuccess(data)) {
                    if ( typeof callback == 'function') {
                        callback();
                    }
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
                } else {
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
                }
            });
        }

        function refreshMultiData($, callback) {
            initial($, containerId);
            if ( typeof callback == 'function') {
                callback();
            }
            top.ModalCommon.open(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
        }

        function cancelMultiData($, element) {
            var classId = $(element).attr("classId");
            $("#newMultiContainer_" + classId).hide();
            $("#newMultiContent_" + classId).empty();
        }

        function selectAll($, element) {
            var classId = $(element).attr('classId');
            var checklist = $("." + classId + "_checklist");
            var len = checklist.length;
            if ($(element).is(":checked")) {
                for (var i = 0; i < len; i++) {
                    $(checklist[i]).attr("checked", "checked");
                    $(checklist[i]).parent().addClass("checked");
                }
            } else {
                for ( i = 0; i < len; i++) {
                    $(checklist[i]).removeAttr("checked");
                    $(checklist[i]).parent().removeClass("checked");
                }
            }
        }
        
        function initMainBoxContentHeight($) {
		    $("table[id^=instanceBox_]").each(function() {
		        var len = $("tr", $(this)).length;
		        if(len > 1) {
		            $(this).parent().height(133 + (len - 2) * 57 + 170);
		        } else {
		            $(this).parent().height(58);
		        }
		    });
		}
		
		function initAddBoxContentHeight($, elem) {
		    var id = $(elem).attr('classId');
		    var elem = $(".box-content", $("#newMultiContainer_" + id));
		    elem.height(153 + 170);
		}


        return {
            initial : initial,
            loadCurData : loadCurData,
            initMultiBox : initMultiBox,
            markMultiData : markMultiData,
            assignMultiValue : assignMultiValue,
            assignAddMultiValue : assignAddMultiValue,
            newMultiData : newMultiData,
            addMultiData : addMultiData,
            saveMultiData : saveMultiData,
            deleteMultiData : deleteMultiData,
            refreshMultiData : refreshMultiData,
            cancelMultiData : cancelMultiData,
            selectAll : selectAll,
            initMainBoxContentHeight : initMainBoxContentHeight
        }

    }());

