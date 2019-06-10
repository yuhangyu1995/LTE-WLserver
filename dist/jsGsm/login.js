var Authorization = {
    Admin : "0",
    Guest : "-1"
};

$(document).ready(function() {
    init();
});

function init() {
    $(document).keypress(function(event) {
        if (event.keyCode == 13) {
            login();
        }
    });

    $("#username").val($.cookie("username") == null ? "" : $.cookie("username"));
    $("#password").val("");

    var username = $("#username").val();
    if (username == "") {
        $("#username").focus();
    } else {
        $("#password").focus();
    }
}

function login() {
    var authorization = "";
    var username = $("#username").val();
    var password = $("#password").val();
    if ((username === "user")) {
        authorization = top.Authorization.Admin;
    } else if ((username === "admin")){
        authorization = top.Authorization.Guest;
    }

    APP.login(username, password, authorization);
}
