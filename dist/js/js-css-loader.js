var indexCss = 1;

function loadJs(a, b) {
    if (b) {
        setTimeout(loadingJs(a), b)
    } else {
        setTimeout(loadingJs(a), 50)
    }
}

function loadingJs(a) {
    return function () {
        $.ajax({
            url: "../../js/" + a + "?v=" + top.versionNo,
            async: false,
            dataType: "script"
        })
    }
}

function loadMainJs(a, b) {
    if (b) {
        setTimeout(loadingMainJs(a), b)
    } else {
        setTimeout(loadingMainJs(a), 50)
    }
}

function loadingMainJs(a) {
    return function () {
        $.ajax({
            url: "js/" + a + "?v=" + top.versionNo,
            async: false,
            dataType: "script"
        })
    }
}

function loadCss(a) {
    $("#cssArea").append('<link id="css-' + indexCss + '" rel="stylesheet">');
    $("#css-" + indexCss).attr("href", "../../css/" + a + "?v=" + top.versionNo);
    indexCss++
}

function loadMainCss(a) {
    $("#cssArea").append('<link id="css-' + indexCss + '" rel="stylesheet">');
    $("#css-" + indexCss).attr("href", "css/" + a + "?v=" + top.versionNo);
    indexCss++
};