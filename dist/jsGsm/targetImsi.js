function bindBtnEvent() {

	$('.btn-refresh').click(function(e) {
		e.preventDefault();
		IMSI_EDIT.open();
		bindBtnEvent();
		top.ModalCommon.open2(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
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

var IMSI_EDIT = ( function() {
    
    var content = "";
    var mainCatchImsi = "";
    
    function init() {
		
		var cgiData1 = {
            cmd : 'query',
            params : "[TargetList]" 
        };
		var cgiData2 = {
            cmd : 'query',
            params : "[TargetInfo]" 
        };
    
		var data1 = top.CGI.ajaxAsyn(top.CGI.requesturl.oamapp, cgiData1);
        if (top.CGI.isSuccess(data1)) {
            var result = top.CGI.getJsonContent(data1).singleData;
            for (var i = 0; i < result.length; i++) {
                for (var key in result[i]) {
                    if ("AccessList" == key) {
                        content = result[i][key];
                    } 
                }
            }
        }
		
		var data2 = top.CGI.ajaxAsyn(top.CGI.requesturl.oamapp, cgiData2);
        if (top.CGI.isSuccess(data2)) {
            var result = top.CGI.getJsonContent(data2).singleData;
            for (var i = 0; i < result.length; i++) {
                for (var key in result[i]) {
                    if ("MainCatchImsiInfo" == key) {
                        mainCatchImsi = result[i][key];
                    } 
                }
            }
        }
						
        var html = [];
        html.push('<div class="box span12">');
        html.push('    <div class="box-header">');
        html.push('        <h2><i class="halflings-icon list"></i><span class="break"></span>目标列表</h2>');
        html.push('        <div class="box-icon">');
        html.push('		   		<a href="#" class="btn-refresh" data-rel="tooltip" data-original-title=""><i class="halflings-icon refresh"></i></a>');
		html.push('		   		<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');
        html.push('        </div>');
        html.push('    </div>');
        html.push('    <div class="box-content">');
        html.push('    <nobr><div class="span3 rebeccapurple">编辑完成后请点击【保存】！</div>');
        if ("" != mainCatchImsi) {
            html.push('<div id="main_catch_imsi" class="span3">', '主目标：', mainCatchImsi, '</div></nobr>');
        } else {
            html.push('<div id="main_catch_imsi" class="span3"></div></nobr>');
        }
        html.push('    <div class="span6 push-right">');
        html.push('        <div class="control-group success inline">');
        html.push('            <div class="controls">');
        html.push('                <input type="text" class="noMargin border-light-gray" placeholder="请输入IMSI（15位数字）">');
        html.push('                <a class="btn btn-danger" href="#" data-rel="tooltip" data-original-title="新增" onclick="IMSI_EDIT.add(event, this)"> <i class="icon-plus white"></i></a>');
        html.push('            </div>');
        html.push('         </div>');
        html.push('         <button type="submit" class="btn btn-success" onclick="IMSI_EDIT.imsi_save(event)">保存</button>');
        html.push('         <button type="submit" class="btn btn-primary" onclick="IMSI_EDIT.main_imsi_config(event, this)">设置主目标</button>');
        html.push('    </div>');
        html.push('        <table class="table table-striped table-bordered bootstrap-datatable datatable color-success">');
        html.push('            <thead>');
        html.push('                <tr>');
        html.push('                    <th class="span2">序号</th>');
        html.push('                    <th class="span3">目标IMSI</th>');
        html.push('                    <th class="span2">有效状态</th>');
        html.push('                    <th class="span5">操作选项</th>');
        html.push('                </tr>');
        html.push('            </thead>');
        html.push('            <tbody id="imsi_edit_body">');
        if (content.length > 0) {
            var imsiArr = content.split(",");
            for (var i = 0; i < imsiArr.length; i++) {
                html.push('         <tr>');
                html.push('             <td class="span2 center">', i + 1, '</td>');
                html.push('             <td class="span3 center imsi">', imsiArr[i], '</td>');
                html.push('             <td class="span2 center imsi-status valid"><span class="label label-success">有效</span></td>');
                html.push('             <td class="span5 center">');
                html.push('                 <a class="btn btn-success" href="#" data-rel="tooltip" data-original-title="恢复" onclick="IMSI_EDIT.valid(event, this)"> <i class="icon-undo white"></i></a>');
                html.push('                 <a class="btn btn-info" href="#" data-rel="tooltip" data-original-title="编辑" onclick="IMSI_EDIT.edit(event, this)"> <i class="icon-edit white"></i></a>');
                html.push('                 <a class="btn btn-danger" href="#" data-rel="tooltip" data-original-title="删除" onclick="IMSI_EDIT.invalid(event, this)"> <i class="icon-trash white"></i></a>');
                html.push('                 <div class="control-group success imsi_edit" style="display: none;">');
                html.push('                     <div class="controls">');
                html.push('                         <input type="text" class="noMargin border-light-gray" placeholder="请输入要替换的IMSI（15位数字）" onblur="IMSI_EDIT.save(event, this)">');
                html.push('                     </div>');
                html.push('                 </div>');
                html.push('             </td>');
                html.push('         </tr>');
            }
        }
        html.push('            </tbody>');
        html.push('        </table>');
        html.push('    </div>');
        html.push('</div>');
        
        $("#imsi_edit").empty().append(html.join(""));
    }
    
    function open() {
        
        init();
        $("#imsi_edit").show();
        
        $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
            "placement" : "bottom",
            delay : {
                show : 400,
                hide : 200
            }
        });
        parent.autoHeight();
    }
    
    function valid(e, element) {
        e.preventDefault();
        var tr = $(element).parent().parent();
        $(".imsi", $(tr)).css("text-decoration", "");
        $(".imsi-status", $(tr)).empty().append('<span class="label label-success">有效</span>');
        $(".imsi-status", $(tr)).removeClass("invalid").addClass("valid");
    }
    
    function invalid(e, element) {
        e.preventDefault();
        var tr = $(element).parent().parent();
        $(".imsi", $(tr)).css("text-decoration", "line-through");
        $(".imsi-status", $(tr)).empty().append('<span class="label label-important">待删除</span>');
        $(".imsi-status", $(tr)).removeClass("valid").addClass("invalid");
    }
    
    function edit(e, element) {
        e.preventDefault();
        $(".imsi_edit", $(element).parent()).css({"margin-bottom": "0px", "display": "inline-block"});
        $(".imsi_edit", $(element).parent()).children().children().focus();
    }
    
    function save(e, element) {
        e.preventDefault();
        var imsi = $(element).val();
        var regex = /^(\d{15})$/;
        if ($.trim(imsi) != "") {
            $(element).parent().parent().parent().prev().prev().text(imsi);
            $(element).parent().parent().css("display", "none");
        } else {
            $(element).parent().parent().css("display", "none");
        }
    }
    
    function add(e, element) {
        
        e.preventDefault();
        if($("#imsi_edit_body").children().length >= 10) {
            top.ModalCommon.open2(top.DISPLAY.Main.Notice, "目标IMSI不能超过10个", "error");
            return false;
        }
        
        var imsi = $.trim($(element).prev().val());
        var regex = /^(\d{15})$/;
        if(!regex.test(imsi)) {
            top.ModalCommon.open2(top.DISPLAY.Main.Notice, "IMSI： " + imsi + "必须为15位数字", "error");
            return false;
        } 
        
        var html = [];
        html.push('<tr>');
        html.push(' <td class="span2 center">', $("#imsi_edit_body").children().length + 1, '</td>');
        html.push(' <td class="span3 center imsi">', imsi, '</td>');
        html.push(' <td class="span2 center imsi-status valid"><span class="label label-success">有效</span></td>');
        html.push(' <td class="span5 center">');
        html.push('     <a class="btn btn-success" href="#" data-rel="tooltip" data-original-title="恢复" onclick="IMSI_EDIT.valid(event, this)"> <i class="icon-undo white"></i></a>');
        html.push('     <a class="btn btn-info" href="#" data-rel="tooltip" data-original-title="编辑" onclick="IMSI_EDIT.edit(event, this)"> <i class="icon-edit white"></i></a>');
        html.push('     <a class="btn btn-danger" href="#" data-rel="tooltip" data-original-title="删除" onclick="IMSI_EDIT.invalid(event, this)"> <i class="icon-trash white"></i></a>');
        html.push('     <div class="control-group success imsi_edit mbottom0" style="display: none;">');
        html.push('         <div class="controls">');
        html.push('             <input type="text" class="noMargin border-light-gray" placeholder="请输入要替换的IMSI（15位数字）" onblur="IMSI_EDIT.save(event, this)">');
        html.push('         </div>');
        html.push('     </div>');
        html.push(' </td>');
        html.push('</tr>');
        
        $("#imsi_edit_body").append(html.join(""));
        
        $('[rel="tooltip"],[data-rel="tooltip"]').tooltip({
            "placement" : "bottom",
            delay : {
                show : 400,
                hide : 200
            }
        });
        parent.autoHeight();
    }
    
    function imsi_save(e) {
        
        e.preventDefault();
        var result = "", imsi = "";
        var regex = /^(\d{15})$/;
        $("tr", $("#imsi_edit_body")).each(function (index,domEle){
            if($(".imsi-status", $(this)).hasClass("valid")) {
                if(!regex.test($.trim($(".imsi", $(this)).text()))) {
                    result += ($.trim($(".imsi", $(this)).text()) + "、");
                } else {
                    imsi += ($.trim($(".imsi", $(this)).text()) + ",");
                }
            }
        });
        
        if(result != "") {
            result = result.substring(0, result.length - 1);
            top.ModalCommon.open2(top.DISPLAY.Main.Notice, "IMSI： " + result + "必须为15位数字", "error");
            return false;
        } else {
            imsi = (imsi != "" ? imsi.substring(0, imsi.length - 1) : "");
            var cgiData = {
                cmd : 'config',
                params : '[TargetList](AccessList="' + imsi + '")'
            };
        
            top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data){
                if (top.CGI.isSuccess(data)) {
                    IMSI_EDIT.open();
                    bindBtnEvent();
                    top.ModalCommon.open2(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
                } else {
                    top.ModalCommon.open2(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
                }
            });
        }
    }
    
    function main_imsi_config(e, element) {
        e.preventDefault();
        
        var regex = /^(\d{15})$/;
        var imsi = $.trim($("input", $(element).prev().prev()).val());
        
        if(imsi !="" && !regex.test(imsi)) {
            top.ModalCommon.open2(top.DISPLAY.Main.Notice, "IMSI： " + imsi + "必须为15位数字", "error");
            return false;
        } 
        
        var cgiData = {
            cmd : 'config',
            params : '[TargetInfo](MainCatchImsiInfo="' + imsi + '")'
        };
    
        top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data){
            if (top.CGI.isSuccess(data)) {
                IMSI_EDIT.open();
                bindBtnEvent();
                top.ModalCommon.open2(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationSuccess, "ok");
            } else {
                top.ModalCommon.open2(top.DISPLAY.Main.Notice, top.DISPLAY.NoticeInfo.OperationFail, "error");
            }
        });
    }
    
    return {
            open : open,
            valid : valid,
            invalid : invalid,
            edit : edit,
            save : save,
            add : add,
            imsi_save : imsi_save,
            main_imsi_config : main_imsi_config
        };
    }());
    

$(document).ready(function(){
    IMSI_EDIT.open();
    bindBtnEvent();
    bindToolTip();
});