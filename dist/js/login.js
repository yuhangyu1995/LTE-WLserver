var T = null;
$(document).ready(function () {
    $("#username").focus(function () {
        $(this).parent(".input-prepend").addClass("input-prepend-focus")
    });
    $("#username").focusout(function () {
        $(this).parent(".input-prepend").removeClass("input-prepend-focus")
    });
    $("#password").focus(function () {
        $(this).parent(".input-prepend").addClass("input-prepend-focus")
    });
    $("#password").focusout(function () {
        $(this).parent(".input-prepend").removeClass("input-prepend-focus")
    });
    lmtDataAsyn("get WebConfig", ["WebLanguage"], null, function (a) {
        initValue(a[0][1])
    }, function (a) {
        initValue("ZH")
    })
});

function lmtDataAsyn(b, a, h, e, d, c) {
    var j = b.indexOf("get") != -1 ? "get" : "set",
        f, g, k;
    if (j === "set" && a.length > 0) {
        b += ": ";
        for (f = 0; f < a.length; f++) {
            b += a[f][0] + "=" + a[f][1] + ","
        }
        b = b.substring(0, b.length - 1)
    }
    b += ";";
    if (typeof (e) !== "function") {
        e = function (i) {}
    }
    if (typeof (d) !== "function") {
        d = function () {}
    }
    if (typeof (c) !== "function") {
        c = function () {}
    }
    g = [];
    k = "/req/connectCliReq.asp";
    $.ajax({
        type: "POST",
        async: true,
        timeout: h,
        url: k,
        data: "caseType=2&cmd=" + b,
        error: function (i) {
            d()
        },
        complete: function (i) {
            c()
        },
        success: function (l) {
            var r = l.indexOf("RESULT"),
                q = l.indexOf("\n", r),
                p = l.substring(r, q),
                n, o, s, m;
            if (j === "get") {
                if (p.indexOf("000") === -1) {
                    d();
                    return
                }
                n = 0;
                o = 0;
                n = l.indexOf("CONTENT");
                n = l.indexOf("\n", n);
                l = l.substring(n, l.length);
                s = "";
                for (m = 0; m < a.length; m += 1) {
                    n = l.indexOf(a[m]);
                    o = l.indexOf("\n", n);
                    s = $.trim(l.substring(n + a[m].length, o));
                    n = s.indexOf("=");
                    s = $.trim(s.substring(n + 1, s.length));
                    n = s.indexOf('"');
                    o = s.indexOf('"', n + 1);
                    if (o === -1) {
                        o = s.length
                    }
                    s = s.substring(n + 1, o);
                    g.push([a[m], s])
                }
                e(g)
            } else {
                if (j === "set") {
                    if (p.indexOf("000") !== -1) {
                        e()
                    } else {
                        d()
                    }
                }
            }
        }
    })
}

function initValue(c) {
    var a = "";
    var b = "";
    if (c == "EN") {
        T = T_EN;
        SetCookie("language", "EN")
    } else {
        T = T_ZH;
        SetCookie("language", "ZH")
    }
    translate();
    document.title = "Login";
    $("#password").keypress(function (e) {
        var g;
        g = e.keyCode ? e.keyCode : e.which ? e.which : e.charCode;
        var f = String.fromCharCode(g);
        if (g == 13) {
            checkLogin()
        }
    });
    $("#username").val(GetCookie("username"));
    $("#password").val("");
    var d = $("#username").val();
    if (d == "") {
        $("#username").focus()
    } else {
        $("#password").focus()
    }
    $(window).resize()
}

function translate() {
    $("#userRole").text(top.T.Home.UserType);
    var a = USER.getAccounts(),
        c = "";
    for (var b = 0; b < a.length; b += 1) {
        c += "<option value='" + a[b][0] + "'>" + a[b][1] + "</option>"
    }
    $("#role").append(c)
}

function checkLogin() {
    var c = $("#role").val();
    var d = $("#username").val();
    var a = $("#password").val();
    var b = top.Branch.CProduct;
    if (d === "debug") {
        c = top.Roles.Market
    } else {
        if (d === "admin") {
            c = top.Roles.Market
        } else {
            if ((d === "user") && (b == 16)) {
                c = top.Roles.Root
            } else {
                c = "noRoot"
            }
        }
    }
    a = md5(a);
    USER.login(d, a, c)
}
$(window).resize(function () {
    action_resize()
});

function action_resize() {
    var a = $(window).height();
    var b = $(window).width()
}

function setErrMSG(a) {
    ModalCommon.open(top.T.Main.Notice, a, "ok")
};