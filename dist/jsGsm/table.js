function TABLE(attribute) {

    var table = {

        dataBlock : [],
        rows : 0,
        cols : 0,

        pageNo : 0,
        pageSize : 10,

        tableId : "",
        tableName : "",
        tableHeader : "",

        skipBtnStart : 0,
        skipBtnEnd : 0,
        width : 1024
    };

    table.init = function() {
        table.cols = attribute.tableHeader.length;
        table.tableId = attribute.tableId;
        table.tableName = attribute.tableName;
        table.tableHeader = attribute.tableHeader;
    };

    table.empty = function() {
        table.dataBlock = [];
    };

    table.loadData = function(data) {
        if ( data instanceof Array) {
            table.dataBlock = table.dataBlock.concat(data);
            table.rows = Math.ceil(table.dataBlock.length / table.cols);
        }
    };

    table.skip = function(pageNum) {
        var total = Math.ceil(table.rows / table.pageSize);
        table.pageNo = (pageNum == 'last' || pageNum >= total) ? (total - 1) : pageNum;

        var parent = $("#" + table.tableId).parent().parent();
        parent.empty().append(table.listTable());
    };

    table.goto2 = function(node) {
        var pageNum = $(node).next().val();
        var reg = /^\d+$/;
        if (reg.test(pageNum.toString())) {
            pageNum = pageNum - 1;
            if (pageNum >= 0) {
                table.skip(pageNum);
            }
        }
    };

    table.listTable = function() {

        var html = [];

        html.push('<div>');
        html.push(' <table id="', table.tableId, '" class="table table-striped table-bordered bootstrap-datatable datatable">');
        html.push(table.loadHeader());
        html.push(table.loadBody());
        html.push(' </table>');
        html.push('</div>');

        if (table.rows > 0) {
            html.push('<div class="pagination pagination-centered">');
            html.push(table.loadFooter());
            html.push('</div>');
        }

        return html.join("");
    };

    table.loadHeader = function() {

        var html = [];
        html.push('<thead>');
        html.push(' <tr>');
        for (var i = 0; i < table.tableHeader.length; i++) {
            var percent = (table.tableHeader[i].width.split("%")[0]) / 100;
            var width = parseInt(percent * table.width);
            html.push('<th nowrap="nowrap" width="', width, '" class="cursor-pointer">', table.tableHeader[i].name, '</th>');
        }
        html.push(' </tr>');
        html.push('</thead>');

        return html.join("");
    };

    table.loadBody = function() {

        var html = [];
        html.push('<tbody>');

        var startIndex = table.pageNo * table.pageSize;

        for (var i = 0; i < table.pageSize; i++) {
            if ((startIndex + i) <= table.rows) {
                html.push('<tr id="', table.tableName, '_', i, '">');
                for (var j = 0; j < table.cols; j++) {
                    var content = table.dataBlock[(startIndex + i) * table.cols + j];
                    if (null != content && content != '') {
                        html.push('<td>', content, '</td>');
                    } else {
                        html.push('<td>&nbsp;</td>');
                    }
                }
                html.push('</tr>');

            } else {
                html.push('<tr>');
                for (var k = 0; k < table.cols; k++) {
                    html.push('<td>&nbsp;</td>');
                }
                html.push('</tr>');
            };
        }

        html.push('</tbody>');

        return html.join("");
    }

    table.loadFooter = function() {

        var html = [];
        var totalPage = Math.ceil(table.rows / table.pageSize);
        if (table.pageNo < 3) {
            table.skipBtnStart = 0;
            table.skipBtnEnd = totalPage > 6 ? 6 : totalPage;
        } else {
            if (table.pageNo >= (totalPage - 3)) {
                table.skipBtnEnd = totalPage;
                table.skipBtnStart = (totalPage - 6) > 0 ? (totalPage - 6) : 0;
            } else {
                table.skipBtnStart = table.pageNo - 3;
                table.skipBtnEnd = (table.pageNo + 3) > totalPage - 1 ? (totalPage - 1) : (table.pageNo + 3);
            }
        }

        html.push('<ul>');

        /* load Fst Btn*/
        if (table.rows > 0) {
            html.push('<li>');
            html.push(' <a href="#" onclick="', table.tableName, '.skip(0)">首页</a>');
            html.push('</li>');
        }

        /* load Skip Btn*/
        for (var i = table.skipBtnStart; i < table.skipBtnEnd; i++) {
            if (table.pageNo == i) {
                html.push('<li class="active">');
                html.push(' <a href="#" onclick="', table.tableName, '.skip(', i, ')">', (i + 1), '</a>');
                html.push('</li>');
            } else {
                html.push('<li>');
                html.push(' <a href="#" onclick="', table.tableName, '.skip(', i, ')">', (i + 1), '</a>');
                html.push('</li>');
            }
        }

        /* load Lst Btn*/
        if (table.rows > 0) {
            html.push('<li>');
            html.push(' <a href="#" onclick="', table.tableName, '.skip(&quot;last&quot;)">末页</a>');
            html.push('</li>');
        }

        /* load Goto Btn*/
        if (table.rows > 100) {
            html.push('<button class="btn btn-primary btn-round btn-goto" onclick="', table.tableName, '.goto2(this)">跳转</button>');
            html.push('<input id="" type="text" class="width24" value=""/>');
        }

        html.push('</ul>');

        return html.join("");
    }

    table.init();

    return table;

} 