var WebTypeObj = (function () {
    var b = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 20],
        a = [10, 11, 12, 13, 14, 15, 16, 17, 8, 9, 20];

    function d() {
        return b
    }

    function c() {
        return a
    }

    function f(k) {
        var j = 0;
        k = window.parseInt(k, 10);
        for (j = 0; j < b.length; j += 1) {
            if (k === b[j]) {
                return true
            }
        }
        return false
    }

    function e(k) {
        var j = 0;
        k = window.parseInt(k, 10);
        for (j = 0; j < a.length; j += 1) {
            if (k === a[j]) {
                return true
            }
        }
        return false
    }

    function g(k) {
        var j = 0;
        k = window.parseInt(k, 10);
        for (j = 0; j < b.length; j += 1) {
            if (k === b[j]) {
                return a[j]
            }
        }
        return k
    }

    function h(k) {
        var j = 0;
        k = window.parseInt(k, 10);
        for (j = 0; j < a.length; j += 1) {
            if (k === a[j]) {
                return b[j]
            }
        }
        return k
    }
    return {
        isEnable: f,
        isDisable: e,
        toDisable: g,
        toEnable: h
    }
}());
(function () {
    var d, a;

    function e(h) {
        h = h || window.event;
        var i = (h.pageX || h.pageY) ? {
            x: h.pageX,
            y: h.pageY
        } : {
            x: h.clientX + top.document.body.scrollLeft - top.document.body.clientLeft,
            y: h.clientY + top.document.body.scrollTop - top.document.body.clientTop
        };
        return i
    }

    function c(h) {
        if (window.attachEvent) {
            a.setCapture()
        } else {
            if (window.addEventListener) {
                window.addEventListener("mouseup", g, true)
            }
        }
    }

    function b(h) {
        if (window.attachEvent) {
            a.releaseCapture();
            a.detachEvent("onlosecapture", f)
        } else {
            if (window.removeEventListener) {
                window.removeEventListener("mouseup", g, true)
            }
        }
    }

    function g() {
        a.onmousemove = null;
        b(g)
    }

    function f() {
        a.onmousemove = null
    }
    ModalControl = function (h, i) {
        d = h;
        a = i
    };
    ModalControl.prototype = {
        dragStart: function (j) {
            var i = e(j),
                h = {
                    x: i.x - (d.offsetLeft),
                    y: i.y - (d.offsetTop)
                };
            d.onmousemove = function (l) {
                var k = e(l);
                d.style.left = k.x - h.x + "px";
                d.style.top = k.y - h.y + "px"
            };
            d.onmouseup = g;
            c(g)
        }
    }
}());
var Modal = (function () {
    var g = function () {},
        k = function () {},
        j = function () {
            return true
        };

    function d() {
        $("#coverPage").show()
    }

    function e() {
        if ($("#progress").css("display") !== "block") {
            $("#coverPage").hide()
        }
    }

    function l(p) {
        d();
        $("#progress").show();
        if (p) {
            $("#progressTitle").text(p)
        }
    }

    function f() {
        $("#progress").hide();
        $("#loading").hide();
        $("#progressTitle").empty();
        e()
    }

    function n(p) {
        $("#events").empty();
        var q = 0,
            r = 0;
        for (q = 0, r = p.length; q < r; q += 1) {
            $("#events").append("<p style='margin-top:10px;margin-left:150px;'>" + p[q] + "</p>")
        }
        $("#events").show();
        $("#loading").show()
    }

    function h(s, q, r, p) {
        d();
        if (p) {
            g = p
        } else {
            g = function () {}
        }
        $("#alert_title").empty();
        $("#alert_title").append(s);
        $("#alert_context").removeClass();
        $("#modalTips_btn_reboot").hide();
        $("#modalTips_btn_submit").show();
        if (r === "choice") {
            $("#alert_context").addClass("question");
            $("#modalTips_btn_cancel").show()
        } else {
            if (r === "error") {
                $("#alert_context").addClass("error");
                $("#modalTips_btn_cancel").hide()
            } else {
                if (r === "reboot") {
                    $("#alert_context").addClass("ok");
                    $("#modalTips_btn_reboot").show();
                    $("#modalTips_btn_submit").hide();
                    $("#modalTips_btn_cancel").show()
                } else {
                    $("#alert_context").addClass(r);
                    $("#modalTips_btn_cancel").hide()
                }
            }
        }
        $("#alert_inform").empty();
        $("#alert_inform").append(q);
        $("#alert_window").show();
        m()
    }

    function m() {
        var p = $("#dialog_center").width();
        if (p < 350) {
            $("#dialog_center").css("width", "350px")
        } else {
            $("#dialog_center").css("width", "450px")
        }
    }

    function o() {
        $("#alert_window").hide();
        e();
        g(1)
    }

    function a() {
        $("#alert_window").hide();
        e();
        g(-1)
    }

    function i(t, q, p, s, r) {
        d();
        if (typeof (p) === "function") {
            k = p
        } else {
            k = function () {}
        }
        if (typeof (s) === "function") {
            j = s
        } else {
            j = function () {
                return true
            }
        }
        $("#wlan_guide_iframe").attr("style", r);
        $("#dialog_tit_div").empty();
        $("#dialog_tit_div").append(t);
        $("#wlan_guide_iframe").attr("src", q + ((q.indexOf("?") >= 0) ? "" : ("?r=" + DeviceInfo.getPartialVersion())));
        $("#wlan_guide").show()
    }

    function c(p) {
        if (typeof (p) !== "undefined") {
            AutoParamConfig.var_guide_exit_manual(p)
        } else {
            AutoParamConfig.var_guide_exit_manual(false)
        }
        if (!j()) {
            return
        }
        $("#wlan_guide").hide();
        e();
        k()
    }

    function b() {
        j = function () {
            return true
        }
    }
    return {
        progress: l,
        delProgress: f,
        showSteps: n,
        open: h,
        setWidth: m,
        submit: o,
        cancel: a,
        openPopUpPage: i,
        closePopUpPage: c,
        clearBeforeCloseFun: b,
        delCover: e,
        cover: d
    }
}());
var LMT = (function () {
    function h(o) {
        if (typeof (o) === "undefined") {
            return false
        }
        var r, q, p = "";
        if (typeof (o) === "object") {
            p = o.result || ""
        } else {
            if (typeof (o) === "string") {
                r = o.indexOf("RESULT");
                q = o.indexOf("\n", r);
                p = o.substring(r, q)
            }
        }
        if (p.indexOf("000") !== -1) {
            return true
        } else {
            return false
        }
    }

    function b(p, o) {
        var r = p.indexOf("@") == 0 ? true : false;
        if (r) {
            p = p.substring(1)
        }
        p = p.replace(/[\r\n]+$/, "");
        p = e(p);
        var t = "/req/connectCliReq.asp",
            s = 1,
            q = function () {
                $.ajax({
                    url: t,
                    async: true,
                    type: "post",
                    cache: false,
                    timeout: 360010,
                    data: {
                        cmd: p
                    },
                    success: function (v) {
                        if (-1 != p.indexOf("download complete")) {
                            try {
                                if (v.length < 1) {
                                    o("update_exception_sucess")
                                } else {
                                    o(v)
                                }
                            } catch (u) {
                                o("update_exception_sucess")
                            }
                        } else {
                            o(v)
                        }
                    },
                    error: function (x, w, v) {
                        s += 1;
                        if (!r) {
                            ModalCommon.progress(top.T.Main.PageWait)
                        }
                        if (s > 3 && !r) {
                            var u = top.T.AlertInfo.LoadingCmdErrInfo;
                            ModalCommon.delProgress();
                            ModalCommon.open(top.T.Main.Notice, u, "error")
                        } else {
                            q()
                        }
                    }
                })
            };
        q()
    }

    function a(o, q) {
        q = q || 30010;
        var r = "/req/connectCliReq.asp",
            p = null;
        o = o.replace(/[\r\n]+$/, "");
        o = e(o);
        $.ajax({
            url: r,
            type: "post",
            async: false,
            cache: false,
            timeout: q,
            data: {
                cmd: o
            },
            success: function (s) {
                p = s
            },
            error: function (u, t, s) {
                ModalCommon.open(top.T.Main.Notice, top.T.AlertInfo.ConnectFailInfo, "error");
                p = "连接中断"
            }
        });
        return p
    }

    function m(p, o) {
        var q = "/req/wlanSession.asp";
        p = e(p);
        $.ajax({
            url: q,
            async: true,
            type: "post",
            contentType: "text/html",
            cache: false,
            timeout: 360010,
            data: {
                cmd: p
            },
            success: function (r) {
                o(r)
            },
            error: function (u, t, r) {
                var s = "连接中断";
                o(s)
            }
        })
    }

    function l(o, q) {
        q = q || 30010;
        var r = "/req/wlanSession.asp",
            p = null;
        o = e(o);
        $.ajax({
            url: r,
            async: false,
            type: "post",
            contentType: "text/html",
            cache: false,
            timeout: q,
            data: {
                cmd: o
            },
            success: function (s) {
                p = s
            },
            error: function (u, t, s) {
                p = "连接中断"
            }
        });
        return p
    }

    function k(o) {
        var p = g(o);
        p.name_val = [];
        if (p.result === "000") {
            p.name_val = i(p.content)
        }
        return p
    }

    function g(p) {
        var s = {
                text: p,
                header: "",
                result: "",
                hint: "",
                content: ""
            },
            q = {
                text: "",
                header: "",
                result: "005",
                hint: "",
                content: ""
            };
        if (p.indexOf("\r\nCMD") !== 0) {
            return q
        }
        var t = p.indexOf(":");
        if (t === -1) {
            return q
        }
        var u = p.indexOf("\n", 2);
        s.header = p.substring(t + 1, u).replace(/\r/g, "").replace(/^ */, "");
        p = p.substring(u + 1, p.length);
        if (p.indexOf("RESULT") !== 0) {
            return q
        }
        t = p.indexOf(":");
        if (t === -1) {
            return q
        }
        u = p.indexOf("\n", 2);
        s.result = p.substring(t + 1, u).replace(/\r/g, "");
        s.result = s.result.match(/\d+/)[0];
        p = p.substring(u + 1, p.length);
        if (p.indexOf("HINT") !== 0) {
            return q
        }
        t = p.indexOf(":");
        if (t === -1) {
            return q
        }
        u = p.indexOf("\n", 2);
        s.hint = p.substring(t + 1, u).replace(/\r/g, "").replace(/^ */, "");
        p = p.substring(u + 1, p.length);
        if (p.indexOf("CONTENT") !== 0) {
            return q
        }
        t = p.indexOf(":");
        p_end = p.lastIndexOf("->");
        if (t === -1) {
            return q
        }
        s.protoContent = $.trim(p.substring(t + 1, p_end));
        var r = p.indexOf("\r\n", t);
        var w = $.trim(p.substring(t + 1, r));
        s.content = p.substring(r + 2, p_end);
        var v = new RegExp("(\r\n)*CONTENT *: *" + w + " *(\r\n)*", "g");
        s.content = $.trim(s.content.replace(v, "\r\n"));
        return s
    }

    function i(p) {
        var o = p.replace(/\\\\/g, "**"),
            q;
        o = o.replace(/\\\"/g, "**");
        o = d(o);
        q = f(o);
        return j(p, q)
    }

    function j(q, y) {
        var v = [],
            u = [],
            r = 0,
            t = 0,
            s = 0,
            w = {},
            z = "",
            A = null,
            x = 0;
        for (r = 0, t = y.length; r < t; r += 1) {
            v.push(q.substring(s, y[r]));
            s = y[r] + 2
        }
        for (r = 0, t = v.length; r < t; r += 1) {
            x = v[r].indexOf("=");
            w = {};
            w.name = $.trim(v[r].substring(0, x));
            w.val = $.trim(v[r].substring(x + 1, v[r].length));
            w.type = "integer";
            z = w.val.replace(/\n/g, "*");
            A = null;
            if (z.charAt(0) === "[" && z.charAt(z.length - 1) === "]") {
                w.val = w.val.substring(1, w.val.length - 1);
                w.proto = w.val;
                w.type = "array"
            } else {
                if (w.val.charAt(0) === '"' && w.val.charAt(w.val.length - 1) === '"') {
                    w.val = w.val.substring(1, w.val.length - 1);
                    w.type = "string"
                }
            }
            u.push(w)
        }
        return u
    }

    function f(o) {
        var p = 0,
            s = [],
            q = 0,
            r = o.length;
        while ((q = o.indexOf("\r\n", p)) !== -1) {
            s.push(q);
            p = q + 2
        }
        if (r) {
            s.push(r)
        }
        return s
    }

    function d(r) {
        var p = [],
            t = false,
            q = r.length,
            s = 0,
            u = 0,
            o = "";
        for (s = 0; s < q; s += 1) {
            o = r.charAt(s);
            if (o === '"') {
                t = !t;
                p[u] = o;
                u += 1;
                continue
            }
            if (t && o === "\r" && r.charAt(s + 1) === "\n") {
                p[u] = "**";
                u += 1;
                s += 1;
                continue
            }
            p[u] = o;
            u += 1
        }
        return p.join("")
    }

    function n(p) {
        var u = {},
            q = p.split("\n");
        u.name_val = [];
        u.version = q[0];
        u.result = q[1];
        if (u.result === "OK") {
            u.content = q[2];
            var s = q.slice(3, q.length),
                r = 0,
                t = {
                    name: "",
                    val: ""
                },
                v = [];
            for (r = 0; r < s.length; r += 1) {
                v = s[r].split("\t");
                t.name = v[0];
                t.val = v[1];
                if (t.name.length > 0) {
                    u.name_val.push(t)
                }
            }
        }
        return u
    }

    function e(o) {
        o = o.replace(/%/ig, "%25");
        o = o.replace(/#/ig, "%23");
        o = o.replace(/&/ig, "%D2");
        o = o.replace(/\+/ig, "%D3");
        return o
    }

    function c(o) {
        return o.replace(/%27/ig, "'")
    }
    return {
        isCmdSuccess: h,
        cmdSyn: b,
        cmdAsyn: a,
        wlanCmdSyn: m,
        wlanCmdAsyn: l,
        parser: k,
        wlanParser: n
    }
}());
var CMDBuider = (function () {
    function b(o, n, p) {
        if (o === null || o === "" || n.length <= 0 || n.length !== p.length) {
            return ""
        }
        var j = "delete " + o + " : ",
            h = "",
            l = 0,
            m = 0,
            k = 0;
        for (k = 0, l = n.length, m = l - 1; k < l; k += 1) {
            h += n[k] + '= "' + Transfer.symTransfer(p[k]) + '"' + (k === m ? ";" : ", ")
        }
        return j + h
    }

    function a(t, p, u) {
        if (t === null || t === "" || p.length <= 0 || p.length !== u.length) {
            return ""
        }
        var r = localDB.T_search("c_param", "StructName", t),
            w = top.c_paramRowIndex.WebType,
            q = top.c_paramRowIndex.MemberName,
            s = [],
            k = 0,
            n = 0,
            m = 0,
            j = "add " + t + " : ",
            h = "",
            l = 0,
            o = 0,
            v = 0;
        for (k = 0, n = r.length; k < n; k += 1) {
            if (r[k][q] === p[m]) {
                s.push(r[k]);
                m += 1
            }
        }
        for (l = 0, n = p.length, o = n - 1; l < p.length; l += 1) {
            v = window.parseInt(s[l][w], 10);
            if (v === 12 || v === 2 || v === 20) {
                h += p[l] + "=[" + Transfer.symTransfer(u[l]) + "]" + (l === o ? ";" : ", ")
            } else {
                h += p[l] + '= "' + Transfer.symTransfer(u[l]) + '"' + (l === o ? ";" : ", ")
            }
        }
        return j + h
    }

    function c(o, n, p) {
        if (o === null || o === "" || n.length <= 0 || n.length !== p.length) {
            return ""
        }
        var j = "get " + o + " : ",
            h = "",
            k = 0,
            l = 0,
            m = 0;
        for (k = 0, l = n.length, m = l - 1; k < l; k += 1) {
            h += n[k] + '="' + Transfer.symTransfer(p[k]) + '"' + (k === m ? ";" : ", ")
        }
        return j + h
    }

    function d(h) {
        if (h === null || h === "") {
            return ""
        }
        return "get " + h + ";"
    }

    function e(v, p, w) {
        if (v === null || v === "" || p.length <= 0 || p.length !== w.length) {
            return ""
        }
        var s = localDB.T_search("c_param", "StructName", v),
            y = top.c_paramRowIndex.WebType,
            q = top.c_paramRowIndex.MemberName,
            t = [],
            k = 0,
            m = 0,
            r = 0,
            u = 0;
        for (r = p.length; m < r; m += 1) {
            k = 0;
            for (u = s.length; k < u; k += 1) {
                if (s[k][q] === p[m]) {
                    t.push(s[k]);
                    break
                }
            }
        }
        var j = "set " + v + " : ",
            h = "",
            l = 0,
            n = 0,
            o = 0,
            x = 0;
        for (l = 0, n = p.length, o = n - 1; l < n; l += 1) {
            x = window.parseInt(t[l][y], 10);
            if (x === 12 || x === 2 || x === 20) {
                h += p[l] + "=[" + Transfer.symTransfer(w[l]) + "]" + (l === o ? ";" : ", ")
            } else {
                h += p[l] + '="' + Transfer.symTransfer(w[l]) + '"' + (l === o ? ";" : ", ")
            }
        }
        return j + h
    }

    function g(m, l) {
        if (m.length !== l.length) {
            return ""
        }
        var h = "lmt\017set\017",
            j = 0,
            k = 0;
        for (j = 0, k = m.length; j < k; j += 1) {
            if (m[j] === "" || m[j] === null) {
                continue
            }
            h = h + m[j] + "=" + l[j] + "\017"
        }
        return h
    }

    function f(l) {
        var h = "lmt\017get\017",
            j = 0,
            k = 0;
        for (j = 0, k = vOid.length; j < k; j += 1) {
            h = h + l[j] + "\017"
        }
        return h
    }
    return {
        lmtCreateDelReq: b,
        lmtCreateAddReq: a,
        lmtCreateGetMReq: c,
        lmtCreateGetReq: d,
        lmtCreateSetReq: e,
        wlanCreateSetReq: g,
        wlanCreateGetReq: f
    }
}());
var Struct = (function () {
    function a(u, t, v) {
        var r = CMDBuider.lmtCreateDelReq(u, t, v),
            s = LMT.cmdAsyn(r);
        return LMT.parser(s)
    }

    function l(u, t, v) {
        var r = CMDBuider.lmtCreateSetReq(u, t, v),
            s = LMT.cmdAsyn(r);
        return LMT.parser(s)
    }

    function o(u, t) {
        var r = CMDBuider.wlanCreateSetReq(u, t),
            s = LMT.wlanCmdAsyn(r);
        return LMT.wlanParser(s)
    }

    function c(A) {
        var r = {},
            C = top.c_paramRowIndex.WebType,
            y = top.c_paramRowIndex.MemberName,
            z = localDB.T_search("c_param", "StructName", A),
            t = localDB.T_getColNames("c_param"),
            s = "",
            u = "",
            B = [],
            w = 0,
            x = 0,
            v = "";
        s = CMDBuider.lmtCreateGetReq(A);
        u = LMT.cmdAsyn(s);
        r = LMT.parser(u);
        r.structType = "OAM";
        r.structDesc = z;
        r.colsNames = t;
        r.sLength = z.length;
        r = j(r);
        return r
    }

    function e(v, r) {
        var s = "",
            t = 0,
            u = 0;
        for (u = v.length; t < u; t += 1) {
            s += CMDBuider.lmtCreateGetReq(v[t]);
            if (t !== u - 1) {
                s += "\017"
            }
        }
        LMT.cmdSyn(s, function (z) {
            var A = 0,
                B = 0,
                w = {},
                x = [],
                z = z.split("\017"),
                E = top.c_paramRowIndex.WebType,
                C = top.c_paramRowIndex.MemberName,
                y = localDB.T_getColNames("c_param"),
                D = null;
            if (z.length !== v.length) {
                alert("Struct.getStructsSyn() error.")
            }
            for (B = z.length; A < B; A += 1) {
                w = {};
                w = LMT.parser(z[A]);
                w.structType = "OAM";
                D = localDB.T_search("c_param", "StructName", v[A]);
                w.structDesc = D;
                w.colsNames = y;
                w.sLength = D.length;
                w = j(w);
                x.push(top.objectClone(w))
            }
            r(x)
        })
    }

    function m(t) {
        var v = 0,
            w = 0,
            s = "",
            u = "",
            r = {};
        for (w = t.length; v < w; v += 1) {
            s += t[v];
            if (v !== w - 1) {
                s += "\017"
            }
        }
        u = LMT.cmdAsyn(s);
        u = u.split("\017");
        for (v = 0, w = u.length; v < w; v += 1) {
            r = LMT.parser(u[v]);
            if (r.result !== "000") {
                return r
            }
        }
        return LMT.parser(u[w - 1])
    }

    function d(w, r, u) {
        var v = localDB.T_search("c_param", "StructName", w),
            t = localDB.T_getColNames("c_param"),
            s = CMDBuider.lmtCreateGetReq(w);
        if (u) {
            s = "@" + s
        }
        LMT.cmdSyn(s, function (y) {
            var x = {};
            x = LMT.parser(y);
            x.structType = "OAM";
            x.structDesc = v;
            x.colsNames = t;
            x.sLength = v.length;
            x = j(x);
            r(x)
        })
    }

    function b(z, x, A) {
        var s = CMDBuider.lmtCreateGetMReq(z, x, A),
            u = LMT.cmdAsyn(s),
            r = LMT.parser(u),
            y = localDB.T_search("c_param", "StructName", z),
            t = localDB.T_getColNames("c_param"),
            v = 0,
            w = 0;
        r.structDesc = y;
        r.colsNames = t;
        r.sLength = y.length;
        r = j(r);
        return r
    }

    function g(x, v) {
        var w = [],
            r = 0,
            u = v.length,
            t = 0,
            s = x.name_val.length;
        for (r = 0; r < u; r += 1) {
            for (t = 0; t < s; t += 1) {
                if (x.name_val[t].name === v[r]) {
                    w.push(x.name_val[t].val);
                    break
                }
            }
        }
        return w
    }

    function f(y, s) {
        var x = [],
            r = 0,
            u = 0,
            v = s.length,
            t = y.structDesc.length,
            w = top.c_paramRowIndex.MemberID;
        for (r = 0; r < v; r += 1) {
            for (u = 0; u < t; u += 1) {
                if (y.structDesc[u][w] === s[r]) {
                    x.push(y.name_val[u].val);
                    break
                }
            }
        }
        return x
    }

    function n(r) {
        if (GetCookie("role") === top.Roles.Root) {
            var v = top.c_paramRowIndex.IsPrimaryKey,
                z = top.c_paramRowIndex.WebType,
                s = 0,
                w = 0,
                x = "",
                y = 0,
                u = 0;
            for (s = 0, w = r.structDesc.length; s < w; s += 1) {
                x = r.structDesc[s][z];
                y = x.toString().length;
                if (y === 2 && window.parseInt(x, 10) < 0) {
                    continue
                }
                u = window.parseInt(r.structDesc[s][v], 10);
                if (u !== 1) {
                    r.structDesc[s][z] = WebTypeObj.toEnable(x)
                }
            }
        }
        return r
    }

    function j(r) {
        if (r.name_val) {
            var s = 0,
                t = 0;
            if ("OAM" === r.structType) {
                for (s = 0, t = r.name_val.length; s < t; s += 1) {
                    r.name_val[s].changed = false;
                    r.name_val[s].val = Transfer.deSymTransfer(r.name_val[s].val)
                }
            } else {
                for (s = 0, t = r.name_val.length; s < t; s += 1) {
                    r.name_val[s].changed = false
                }
            }
        }
        return r
    }

    function h(u) {
        var r = 0,
            t = u.length,
            s = top.c_paramRowIndex.IsPrimaryKey;
        for (r = 0; r < t; r += 1) {
            if (window.parseInt(u[r][s], 10) === 1) {
                return true
            }
        }
        return false
    }

    function i(s) {
        var u = top.c_paramRowIndex.Level,
            t = "",
            r = "";
        if (s.length > 0) {
            t = Transfer.parseDbColumn(s[0][u]);
            t = t.split("|");
            if (t.length >= 5) {
                r = $.trim(t[4]);
                if (r === "") {
                    return true
                } else {
                    return window.parseInt(r, 10) === 1
                }
            } else {
                return true
            }
        }
        return false
    }

    function k(u, t) {
        var v = top.c_paramRowIndex.StructID,
            r = 0,
            s = 0;
        for (r = 0, s = u.length; r < s; r += 1) {
            u[r][v] = t
        }
    }

    function p(w) {
        var x = top.c_paramRowIndex.WebType,
            v = -1,
            r = 0,
            u = 0,
            s = 0,
            t = 0;
        for (r = 0, u = w.length; r < u; r += 1) {
            v = window.parseInt(w[r][x], 10);
            w[r][x] = WebTypeObj.toEnable(v)
        }
    }

    function q(w) {
        var x = top.c_paramRowIndex.WebType,
            v = -1,
            r = 0,
            u = 0,
            s = 0,
            t = 0;
        for (r = 0, u = w.length; r < u; r += 1) {
            v = window.parseInt(w[r][x], 10);
            w[r][x] = WebTypeObj.toDisable(v)
        }
    }
    return {
        getStructByName: c,
        getStructByNameSyn: d,
        getStructsSyn: e,
        getMStructByName: b,
        getValuesByMemberName: g,
        getValuesByMemberID: f,
        setStruct: l,
        setStructsByCmds: m,
        deleteStruct: a,
        isMulInstance: h,
        isMulInstanceAddable: i,
        setNewStructID: k,
        toChangeable: p,
        toUnchangeable: q,
        setWlanParams: o
    }
}());
var MatrixStruct = (function () {
    var h = {},
        l = [];

    function f(o) {
        var r = [],
            p = [],
            m = 0,
            n = 0,
            q = [];
        for (m = 0, n = o.length; m < n; m += 1) {
            q = o[m].split(".");
            r.push(q[0]);
            p.push(q[1])
        }
        return [r, p]
    }

    function g(r) {
        var q = [],
            u = [],
            p = [],
            t = top.c_paramRowIndex.StructName,
            o = top.c_paramRowIndex.MemberName,
            m = 0,
            n = 0,
            s = "";
        for (m = 0, n = r[0].length; m < n; m += 1) {
            s = localDB.T_search("c_param", "MemberID", r[1][m]);
            if (s.length > 0) {
                u.push(s[0][t]);
                p.push(s[0][o]);
                q.push(s[0][t] + "_" + s[0][o])
            } else {
                u.push(r[0][m]);
                p.push(r[1][m]);
                q.push(r[0][m] + "_" + r[1][m])
            }
        }
        return [u, p, q]
    }

    function d(p, o) {
        var m = 0,
            n = 0;
        for (m = 0, n = o.length; m < n; m += 1) {
            if (p === o[m]) {
                return true
            }
        }
        return false
    }

    function a(m) {
        var p = [],
            n = 0,
            o = 0,
            q;
        for (n = 0, o = m.length; n < o; n += 1) {
            q = m[n];
            if (!d(q, p)) {
                p.push(q)
            }
        }
        return p
    }

    function e(r) {
        var m = [],
            q = r[1],
            n = 0,
            o = 0,
            p = "";
        for (n = 0, o = q.length; n < o; n += 1) {
            p = localDB.T_search("c_param", "MemberID", q[n]);
            if (p.length > 0) {
                m.push(p[0])
            }
        }
        return m
    }

    function c(m) {
        Struct.getStructsSyn(h.structNameList, function (w) {
            l = w;
            var n = 0,
                o = 0,
                u = "",
                t = "",
                p = 0,
                q = 0,
                r = 0,
                s, v = top.c_paramRowIndex.StructName;
            for (n = 0, o = h.name_val.length; n < o; n += 1) {
                u = h.name_val[n].structName;
                t = h.name_val[n].memberName;
                for (p = 0, q = w.length; p < q; p += 1) {
                    if (u === w[p].structDesc[0][v]) {
                        for (r = 0, s = w[p].name_val.length; r < s; r += 1) {
                            if (t === w[p].name_val[r].name) {
                                h.name_val[n].val = w[p].name_val[r].val;
                                h.structDesc[n][top.c_paramRowIndex.WebType] = w[p].structDesc[r][top.c_paramRowIndex.WebType];
                                break
                            }
                        }
                        break
                    }
                }
            }
            CurrentStructs.push(h);
            if (typeof (m) === "function") {
                m(h)
            }
        })
    }

    function k(p) {
        var q = top.c_paramRowIndex.StructID,
            o = top.c_paramRowIndex.MemberName,
            m = 0,
            n = 0;
        for (m = 0, n = h.structDesc.length; m < n; m += 1) {
            h.structDesc[m][q] = p;
            h.structDesc[m][o] = h.sms[m]
        }
    }

    function i(o) {
        var q = f(o),
            r = g(q),
            p = {
                structName: "",
                memberName: "",
                name: "?",
                changed: false,
                type: "string",
                val: "?"
            },
            m = 0,
            n = 0;
        h.structDesc = e(q);
        h.colsNames = localDB.T_getColNames("c_param");
        h.structNames = r[0];
        h.memberNames = r[1];
        h.sms = r[2];
        h.memberIds = q[1];
        h.structIds = q[0];
        h.structNameList = a(r[0]);
        h.name_val = [];
        for (m = 0, n = h.structDesc.length; m < n; m += 1) {
            p.name = h.sms[m];
            p.structName = h.structNames[m];
            p.memberName = h.memberNames[m];
            h.name_val.push(objectClone(p))
        }
    }

    function b(n, m, o) {
        h = {};
        h.matrix = true;
        l = [];
        i(n);
        k(o);
        c(m)
    }

    function j() {
        var n = 0,
            o = 0,
            u = "",
            t = "",
            p = 0,
            q = 0,
            r = 0,
            v = top.c_paramRowIndex.StructName,
            s = top.c_paramRowIndex.MemberName,
            w = {},
            m = top.objectClone(l);
        for (n = 0, o = h.name_val.length; n < o; n += 1) {
            u = h.name_val[n].structName;
            t = h.name_val[n].memberName;
            if (h.name_val[n].changed) {
                for (p = 0, q = m.length; p < q; p += 1) {
                    w = m[p];
                    if (w.structDesc[0][v] === u) {
                        for (r = 0; r < w.name_val.length; r += 1) {
                            if (w.name_val[r].name === t) {
                                w.name_val[r].val = h.name_val[n].val;
                                w.name_val[r].changed = true;
                                break
                            }
                        }
                        break
                    }
                }
            }
        }
        return m
    }
    return {
        get: b,
        set: j
    }
}());
var DB = (function () {
    function a() {
        var c = {};
        c.Tables = [];
        c.TableNames = [];
        c.CreateTable = function (e, d) {
            var f = new b(d.length);
            c.TableNames.push(e);
            f.T_setColName(d);
            c.Tables.push(f)
        };
        c.DeleteTable = function (g) {
            var e = -1,
                d = 0,
                f = 0;
            for (d = 0, f = c.TableNames.length; d < f; d += 1) {
                if (g === c.TableNames[d]) {
                    e = d;
                    break
                }
            }
            if (e === -1 || e >= f) {
                return false
            } else {
                c.TableNames.splice(e, 1);
                c.Tables.splice(e, 1);
                return true
            }
        };
        c.T_getColNames = function (f) {
            var d = 0,
                e = 0;
            for (d = 0, e = c.TableNames.length; d < e; d += 1) {
                if (c.TableNames[d] === f) {
                    break
                }
            }
            if (d < e) {
                return c.Tables[d].T_getColNames()
            } else {
                throw "ERR: Not such table."
            }
        };
        c.emptyCol = function (g, d) {
            var e = 0,
                f = 0;
            for (e = 0, f = c.TableNames.length; e < f; e += 1) {
                if (c.TableNames[e] === g) {
                    break
                }
            }
            if (e < f) {
                return c.Tables[e].emptyCol(d)
            } else {
                throw "ERR: Not such table."
            }
        };
        c.T_getIndexByName = function (g, d) {
            var e = 0,
                f = 0;
            for (e = 0, f = c.TableNames.length; e < f; e += 1) {
                if (c.TableNames[e] === g) {
                    break
                }
            }
            if (e < f) {
                return c.Tables[e].T_getColIndexByName(d)
            } else {
                throw "ERR: Not such table."
            }
        };
        c.T_import = function (g, d) {
            var e = 0,
                f = 0;
            for (e = 0, f = c.TableNames.length; e < f; e += 1) {
                if (c.TableNames[e] === g) {
                    break
                }
            }
            if (e < f) {
                c.Tables[e].T_import(d)
            } else {
                throw "ERR: No such table."
            }
        };
        c.T_search = function (g, d, h) {
            var e = 0,
                f = 0;
            for (e = 0, f = c.TableNames.length; e < f; e += 1) {
                if (c.TableNames[e] === g) {
                    break
                }
            }
            if (e < f) {
                return c.Tables[e].T_search(d, h)
            } else {
                throw "ERR: Not such table."
            }
        };
        c.T_getColumn = function (g, d) {
            var e = 0,
                f = 0;
            for (e = 0, f = c.TableNames.length; e < f; e += 1) {
                if (c.TableNames[e] === g) {
                    break
                }
            }
            if (e < f) {
                return c.Tables[e].T_getColumn(d)
            } else {
                throw "ERR: No such table."
            }
        };
        c.T_reverse = function (f) {
            var d = 0,
                e = 0;
            for (d = 0, e = c.TableNames.length; d < e; d += 1) {
                if (c.TableNames[d] === f) {
                    break
                }
            }
            if (d < e) {
                return c.Tables[d].T_reverse()
            } else {
                throw "ERR: No such table."
            }
        };
        return c
    }

    function b(c) {
        var d = {};
        d.vector = [];
        d.totalCols = 0;
        d.colNames = [];
        d.T_setColName = function (e) {
            if (e.length === d.totalCols) {
                d.colNames = d.colNames.concat(d.colNames, e)
            } else {
                throw "ERR: Table names length not equal table column length."
            }
        };
        d.T_getColNames = function () {
            return d.colNames
        };
        d.T_getColIndexByName = function (e) {
            var f = -1,
                g = 0,
                h = 0;
            for (g = 0, h = d.colNames.length; g < h; g += 1) {
                if (e === d.colNames[g]) {
                    f = g;
                    break
                }
            }
            if (f >= h || f === -1) {
                return -1
            } else {
                return f
            }
        };
        d.T_import = function (e) {
            if ((e.length % d.totalCols) !== 0) {
                throw "ERR: Table total size not fit the cols."
            }
            d.vector = d.vector.concat(d.vector, e)
        };
        d.T_search = function (f, q) {
            var l = -1,
                m = 0,
                n = 0;
            for (m = 0, n = d.colNames.length; m < n; m += 1) {
                if (f === d.colNames[m]) {
                    l = m;
                    break
                }
            }
            if (l >= n || l === -1) {
                throw "ERR: No such column."
            }
            var o = [],
                k = false;
            var h = l,
                p = "",
                e = 0,
                r = [],
                g;
            for (h = l, n = d.vector.length; h < n; h += d.totalCols) {
                p = d.vector[h];
                if (p == q) {
                    e = h - l;
                    r = d.vector.slice(e, e + d.totalCols);
                    o.push(r)
                }
            }
            return o
        };
        d.emptyCol = function (e) {
            var f = -1,
                g = 0,
                h = 0;
            for (g = 0, h = d.colNames.length; g < h; g += 1) {
                if (e === d.colNames[g]) {
                    f = g;
                    break
                }
            }
            if (f >= h || f === -1) {
                throw "ERR:No such column."
            }
            for (g = f, h = d.vector.length; g < h; g += d.totalCols) {
                d.vector[g] = ""
            }
            return true
        };
        d.T_getColumn = function (e) {
            var g = -1,
                h = 0,
                i = 0,
                f = [];
            for (h = 0, i = d.colNames.length; h < i; h += 1) {
                if (e === d.colNames[h]) {
                    g = h;
                    break
                }
            }
            if (g >= i || g === -1) {
                throw "ERR:No such column."
            }
            for (h = g, i = d.vector.length; h < i; h += d.totalCols) {
                f.push(d.vector[h])
            }
            return f
        };
        d.T_reverse = function () {
            var h = [],
                e = d.vector.length,
                f = d.colNames.length,
                g = 0;
            for (g = e; g > -1; g -= f) {
                h = h.concat(d.vector.slice(g - f, g))
            }
            d.vector = h;
            return true
        };
        d.init = function (e) {
            d.totalCols = e
        };
        d.init(c);
        return d
    }
    return {
        createDB: a
    }
}());
var Transfer = (function () {
    function c(e) {
        return e.toString().replace(/\\/g, "\\\\").replace(/\"/g, '\\"')
    }

    function a(e) {
        return e.toString().replace(/\\\\/g, "\\").replace(/\\\"/g, '"')
    }

    function b(e) {
        return e
    }

    function d(p, k) {
        var f = 0,
            n = "",
            e = [],
            l = [],
            m = [],
            o = p.replace(/\\\\,/g, "\017"),
            h = 0,
            g = o.split("||");
        if (g.length === 2) {
            if (g[0] === "" || top.Language === "EN") {
                h = 1
            }
            if (g[1] === "" && top.Language === "EN") {
                h = 0
            }
        }
        items = g[h].split(",");
        for (f = 0; f < items.length; f += 1) {
            items[f] = items[f].replace(/\017/g, ",");
            n = items[f].replace(/\\\\=/g, "\017");
            e = n.split("=");
            if (e.length === 2) {
                e[0] = e[0].replace(/\017/g, "=");
                e[0] = $.trim(e[0]);
                e[1] = e[1].replace(/\017/g, "=");
                e[1] = $.trim(e[1]);
                if (e[0] == "" || e[1] == "") {
                    alert("成员" + k + "的TypeDesc字段,错误配成x=''或''=y的形式了,应配成X=Y形式!");
                    break
                } else {
                    l.push(e[0]);
                    m.push(e[1])
                }
            } else {
                alert("成员" + k + "的TypeDesc字段" + p + "应配成X=Y形式!");
                break
            }
        }
        return {
            selectNames: l,
            selectVals: m
        }
    }
    return {
        symTransfer: c,
        deSymTransfer: a,
        parseDbColumn: b,
        wtListParser: d
    }
}());
var DeviceInfo = (function () {
    var e, d;

    function b() {
        if (!e) {
            a()
        }
        return e
    }

    function c() {
        if (!d) {
            a()
        }
        return d
    }

    function a() {
        var g = LMT.cmdAsyn("get DeviceInfo;"),
            f = null,
            h = "Description";
        if (USER.hasModule(Branch.Modules.LTE_EN)) {
            h = "ProductClass"
        }
        if (LMT.isCmdSuccess(g)) {
            f = LMT.parser(g);
            $(f.name_val).each(function (j) {
                if (this.name === h) {
                    d = this.val
                }
            });
            $(f.name_val).each(function (j) {
                if (this.name === "SoftwareVersion") {
                    e = this.val.substring(this.val.indexOf(".") + 1)
                }
            })
        }
    }
    return {
        getPartialVersion: b,
        getProductClass: c
    }
}());