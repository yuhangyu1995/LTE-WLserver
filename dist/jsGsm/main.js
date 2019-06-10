$(document).ready(function() {
    APP.isOnline();
    APP.getWebPath();
    autoHeight();
    init();
});

function init() {
    var username = $.cookie('username');
    $("#username").text(username);

    Database.loadDatabase();
}

function autoHeight() {
    var headHeight = $(".navbar").height();
    var breadcrumbHeight = $(".breadcrumb").height();
    var footerHeight = $("#footer").height();
    if (window.frames["mainFrame"].document.body.offsetHeight < (window.screen.availHeight - headHeight - breadcrumbHeight - footerHeight)) {
        document.getElementById("mainFrame").height = window.screen.availHeight - headHeight - breadcrumbHeight - footerHeight;
    } else {
        document.getElementById("mainFrame").height = window.frames["mainFrame"].document.body.offsetHeight;
    }

    var height = parseInt($(".nav-collapse.sidebar-nav").css("height")) < parseInt($("#content").css("height")) ? $("#content").css("height") : $(".nav-collapse.sidebar-nav").css("height");
    height = parseInt($("#sidebar-left").css("min-height")) < parseInt(height) ? height : $("#sidebar-left").css("min-height");
    $("#sidebar-left").css("height", height);
}
