var structMembers=eval("("+parent.G_MIX_PARAM.mixMembers+")");var mixStructID=top.G_MIX_PARAM.mixStructID||"mixStructID";var titleInsert=null;try{titleInsert=eval("("+parent.G_MIX_PARAM.mixTitles+")");if(typeof titleInsert=="undefined"||titleInsert.length==0){titleInsert=null}}catch(e){titleInsert=null}var g_structs=null;try{g_structs=eval("("+parent.G_MIX_MUL_INTANCES+")");if(typeof g_structs=="undefined"||g_structs.length==0){g_structs=null}}catch(e){}function sendSPILMT(a){top.LMT.cmdSyn(a,function(c){var b=top.LMT.parser(c);if(b.result==="000"){top.ModalCommon.open(top.T.Main.Notice,top.T.AlertInfo.SendSuccess,"ok")}else{top.ModalCommon.open(top.T.Main.Notice,top.T.AlertInfo.SendFail,"error")}})}$(document).ready(function(){parent.CurrentStructs.init();initPageSyn();parent.autoHeight()});function initPageSyn(){parent.MatrixStruct.get(structMembers,webStructCallback,mixStructID)}function webStructCallback(c){var b=parent.currentRelate.update(c);var a=parent.HTML.vTable(b,true,true);if(titleInsert===null){a=top.HTML.vTableTitle()+a}$("#paramValueTable").empty().append(a);insertBlocks(titleInsert);parent.CurrentStructs.evaluation($,mixStructID);bindBtnEvent();bindToolTip();triggerRelateDelayEvent();$("#paramValueTable").show();parent.autoHeight()}function triggerRelateDelayEvent(){try{parent.currentRelate.delayEvent(mixStructID,$)}catch(a){}}function bindBtnEvent(){$('[data-rel="chosen"],[rel="chosen"]').chosen();$(".btn-save").click(function(a){a.preventDefault();TableSave();parent.autoHeight()});$(".btn-refresh").click(function(a){a.preventDefault();fresh();parent.autoHeight()});$(".btn-minimize").click(function(b){b.preventDefault();var a=$(this).parent().parent().next(".box-content");if(a.is(":visible")){$("i",$(this)).removeClass("chevron-up").addClass("chevron-down");$(this).removeClass("btn-minimize").addClass("btn-maximize")}else{$("i",$(this)).removeClass("chevron-down").addClass("chevron-up");$(this).removeClass("btn-maximize").addClass("btn-minimize")}bindToolTip();a.slideToggle();parent.autoHeight()})}function bindToolTip(){$(".btn-minimize").attr("data-original-title",top.T.BTN.Minimize);$(".btn-maximize").attr("data-original-title",top.T.BTN.Maximize);$(".btn-save").attr("data-original-title",top.T.BTN.Save);$(".btn-refresh").attr("data-original-title",top.T.BTN.Refresh);$('[rel="tooltip"],[data-rel="tooltip"]').tooltip({"placement":"bottom",delay:{show:400,hide:200}})}function updateCallback(a){}function TableSave(){parent.CurrentStructs.save()}function fresh(){initPageSyn();top.ModalCommon.open(top.T.Main.Notice,top.T.AlertInfo.OperationSuccess,"ok")}function insertBlocks(m){var l=m;var g=$("#paramValueTable").children().length;var f=0,d=0,h=0,k=0,j="";var a=[];for(var c=0;c<g;c++){a[c]=$($("#paramValueTable").children()[c]).prop("outerHTML")}var b=[];for(var c=0;c<Math.floor(l.length/2);c++,h+=2){b.push('<div class="row-fluid sortable">');if((h+1)<l.length){f=l[h].p;d=l[h+1].p;if(parent.Language=="EN"&&l[k].en.length>0){j=l[k++].en}else{j=l[k++].n}b.push(createDivBox(j,f,d,a))}if((h+2)<l.length){f=l[h+1].p;d=l[h+2].p;if(parent.Language=="EN"&&l[k].en.length>0){j=l[k++].en}else{j=l[k++].n}b.push(createDivBox(j,f,d,a))}else{if((h+2)==l.length){f=l[h+1].p;d=g;if(parent.Language=="EN"&&l[k].en.length>0){j=l[k++].en}else{j=l[k++].n}b.push(createDivBox(j,f,d,a))}}b.push("</div>")}if(l.length%2==1){f=l[l.length-1].p;d=g;if(parent.Language=="EN"&&l[k].en.length>0){j=l[k++].en}else{j=l[k++].n}b.push('<div class="row-fluid sortable">');b.push(createDivBox(j,f,d,a));b.push("</div>")}$("#paramValueTable").empty().append(b.join(""))}function createDivBox(g,d,c,a){var b=[];b.push('<div class="box span6">');b.push('<div class="box-header">');b.push('<h2><i class="halflings-icon align-justify"></i><span class="break"></span>',g,"</h2>");b.push('<div class="box-icon">');b.push('<a href="#" class="btn-save" data-rel="tooltip" data-original-title=""><i class="halflings-icon ok"></i></a>');b.push('<a href="#" class="btn-refresh" data-rel="tooltip" data-original-title=""><i class="halflings-icon refresh"></i></a>');b.push('<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>');b.push("</div>");b.push("</div>");b.push('<div class="box-content">','<form class="form-horizontal">',"<fieldset>");for(var f=d;f<c;f++){b.push(a[f])}b.push("</fieldset>","</form>","</div>","</div>");return b.join("")};