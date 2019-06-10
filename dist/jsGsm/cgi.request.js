var CGI = ( function() {
    var cgi_path = "/cgi-bin/";
    var requesturl = {
        login : cgi_path + "login.cgi",
        filemanage : cgi_path + "filemanage.cgi",
        webdir : cgi_path + "webdir.cgi",
        download : cgi_path + "download.cgi",
        upload : cgi_path + "upload.cgi",
        oamapp : cgi_path + "oamapp.cgi",
        webapp : cgi_path + "webapp.cgi"
    };

    function ajaxSyn(url, data, callback, error_callback, timeoutOption) {

        $.ajax({
            url : url,
            async : true,
            type : "post",
            // contentType : "text/html",
            cache : false,
            timeout : 15000,
            data : data,
            success : function(res) {
                callback(res);
            },
            error : function() {
                if ( typeof error_callback == 'function') {
                    error_callback();
                }
            },
            complete : function(XMLHttpRequest, status) {
                if ( typeof timeoutOption == 'function') {
                    if (status == 'timeout' || status == 'error') {
                        timeoutOption();
                    }

                }
            }
        });
    }

    function ajaxAsyn(url, data) {

        var ack = "";

        $.ajax({
            url : url,
            async : false,
            type : "post",
            // contentType : "text/html",
            cache : false,
            timeout : 15000,
            data : data,
            success : function(res) {
                ack = res;
            }
        });

        return ack;
    }

    function ajaxXml(url, async, callback) {
        var ack = "";

        $.ajax({
            url : url,
            async : async,
            type : "post",
            dataType : "xml",
            cache : false,
            timeout : 15000,
            success : function(res) {
                ack = res;
                if (callback) {
                    callback(ack);
                }
            }
        });
    }

    function ajaxJson(url, async, callback) {
        var ack = "";

        $.ajax({
            url : url,
            async : async,
            type : "get",
            dataType : "json",
            cache : false,
            timeout : 15000,
            success : function(res) {
                ack = res;
                if (callback) {
                    callback(ack);
                }
            }
        });
    }

    function ajaxJs(url, async, callback) {
        var ack = "";

        $.ajax({
            url : url,
            async : async,
            type : "post",
            dataType : "script",
            cache : false,
            timeout : 15000,
            success : function(res) {
                ack = res;
                if (callback) {
                    callback(ack);
                }
            }
        });
    }

    function ajaxText(url, async, callback) {
        var ack = "";

        $.ajax({
            url : url,
            async : async,
            type : "post",
            dataType : "text",
            cache : false,
            timeout : 15000,
            success : function(res) {
                ack = res;
                if (callback) {
                    callback(ack);
                }
            }
        });
    }

    function isSuccess(data) {
        var reg = /^\{\"resp\":\"(.*)\",\"content\":([\s\S]*),\"hint\":\"(.*)\"\}$/;
        console.log(data)
        console.log(reg.test(data))
        if (reg.test(data)) {
            return RegExp.$1 == "1";
        }

        return false;
    }

    function getContent(data) {
        var reg = /^\{\"resp\":\"(.*)\",\"content\":\"([\s\S]*)\",\"hint\":\"(.*)\"\}$/;
        if (reg.test(data)) {
            return RegExp.$2;
        }

        return "";
    }

    function getJsonContent(data) {
        var reg = /^\{\"resp\":\"(.*)\",\"content\":([\s\S]*),\"hint\":\"(.*)\"\}$/;
        if (reg.test(data)) {
            return JSON.parse(RegExp.$2);
        }

        return "";
    }

    function getHint(data) {
        var reg = /^\{\"resp\":\"(.*)\",\"content\":([\s\S]*),\"hint\":\"(.*)\"\}$/;
        if (reg.test(data)) {
            return RegExp.$3;
        }

        return "";
    }

    return {
        requesturl : requesturl,
        ajaxSyn : ajaxSyn,
        ajaxAsyn : ajaxAsyn,
        isSuccess : isSuccess,
        getContent : getContent,
        getHint : getHint,
        getJsonContent : getJsonContent,
        ajaxXml : ajaxXml,
        ajaxJson : ajaxJson,
        ajaxJs : ajaxJs,
        ajaxText : ajaxText
    };

}());