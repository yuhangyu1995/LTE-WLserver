var g_database = null;

var g_curData = {
    mixDataDisable : false,
    property : [],
    split : [],
    mixData : [],
    classId : "",
    multiId : [],
    multiData : [],
    multiDefaultData : []
};

var Database = ( function() {

        function loadDatabase() {

            var cgiData = {
                cmd : "query_class_table",
                params : ""
            };
            top.CGI.ajaxSyn(top.CGI.requesturl.oamapp, cgiData, function(data) {
                if (top.CGI.isSuccess(data)) {
                    top.CGI.ajaxJson("../conf/database.json", false, parseDatabase);
                }
            });
        }

        function parseDatabase(data) {
            g_database = data;
            MENU.loadMenu();
        }

        return {
            loadDatabase : loadDatabase
        }

    }())