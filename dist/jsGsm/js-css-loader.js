var indexCss = 1;

// js
function loadJs(name, time) {
    setTimeout(loadingJs(name), time);
}

function loadingJs(name) {
    return function() {
        $.ajax({
            url : "../../js/" + name + "?v=" + top.versionNo,
            async : false,
            dataType : "script"
        });
    }
}

function loadMainJs(name, time) {
    setTimeout(loadingMainJs(name), time);
}

function loadingMainJs(name) {
    return function() {
        $.ajax({
            url : "js/" + name + "?v=" + top.versionNo,
            async : false,
            dataType : "script"
        });
    }
}

// css
function loadCss(name) {
    $('#cssArea').append('<link id="css-' + indexCss + '" rel="stylesheet">');
    $("#css-" + indexCss).attr("href", "../../css/" + name + "?v=" + top.versionNo);
    indexCss++;
}

function loadMainCss(name) {
    $('#cssArea').append('<link id="css-' + indexCss + '" rel="stylesheet">');
    $("#css-" + indexCss).attr("href", "css/" + name + "?v=" + top.versionNo);
    indexCss++;
}
