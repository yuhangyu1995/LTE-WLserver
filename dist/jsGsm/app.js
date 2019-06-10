var APP = ( function() {

        function isOnline() {
            var username = $.cookie("username"), listed = $.cookie("listed");
            if (username == null || listed == null) {
                logout();
            }
        }

        function login(username, password, authorization) {

            var usertmp = username;
            username = md5(username);
            password = md5(password);

            var cgiData = {
                username : username,
                password : password,
                authorization : authorization
            };
            top.CGI.ajaxSyn(top.CGI.requesturl.login, cgiData, function(data) {
                if (top.CGI.isSuccess(data)) {
                    $.cookie('username', usertmp);
                    $.cookie('authorization', authorization);
                    $.cookie('listed', '1');
                    gotoMain();
                } else {
                    top.ModalCommon.open(top.DISPLAY.Main.Notice, top.CGI.getContent(data), "error");
                }
            });
        }

        function logout(msg) {
            if (msg == -1) {
                return;
            };
            $.cookie('authorization', null);
            $.cookie('username', null);
            $.cookie('listed', null);
            window.top.location = 'login.html?_=' + new Date().getTime();
        }

        function getWebPath() {

            var cgiData = {
                cmd : "queryWebPath"
            };

            var data = top.CGI.ajaxAsyn(top.CGI.requesturl.webapp, cgiData);
            if (top.CGI.isSuccess(data)) {
                top.web_path = top.CGI.getContent(data);
            } else {
                top.ModalCommon.open(top.DISPLAY.Main.Notice, top.CGI.getContent(data), "error");
            }
        }
        
        function gotoMain() {
            window.top.location = 'main.html?_=' + new Date().getTime();
        }

        return {
            isOnline : isOnline,
            login : login,
            logout : logout,
            gotoMain : gotoMain,
            getWebPath : getWebPath
        }

    }())

