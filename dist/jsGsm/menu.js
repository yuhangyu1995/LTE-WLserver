var g_menutree = null;
var g_curMenu = null;

var MENU = ( function() {

        function loadMenu() {

            top.CGI.ajaxJson("../menu/menu.json", false, parseMenu);
        }

        function parseMenu(data) {

            var localmenu = data;

            for (var i = 0; i < localmenu.root.length; i++) {
                if (localmenu.root[i].authorization == $.cookie("authorization")) {
                    g_menutree = localmenu.root[i];
                }
            }

            for (var key in g_menutree) {
                if (key.indexOf("branch") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        for (var j = 0; j < g_database.classes.length; j++) {
                            if (g_database.classes[j].menu.indexOf("@") != -1) {
                                continue;
                            }
                            if (g_menutree[key][i].menuName == g_database.classes[j].menu) {
                                var menuNode = createMenuNode(g_database.classes[j].classId, g_database.classes[j].classCnName, g_database.classes[j].propertys[0].multi);
                                g_menutree[key][i].leaf.push(menuNode);
                            }
                        }
                    }
                }
            }

            var branch_index = 0, leaf_index = 0;
            for (var key in g_menutree) {
                if (key.indexOf("branch") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        g_menutree[key][i]["menuId"] = "branch_" + branch_index;
                        for (var j = 0; j < g_menutree[key][i].leaf.length; j++) {
                            g_menutree[key][i].leaf[j]["menuId"] = "leaf_" + branch_index + "." + j;
                        }
                        branch_index++;
                    }
                } else if (key.indexOf("leaf") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        g_menutree[key][i]["menuId"] = "leaf_" + leaf_index;
                        leaf_index++;
                    }
                }
            }

            listMenu();
        }

        function createMenuNode(classId, classCnName, multi) {

            var node = new Object();
            if (multi == '1') {
                node.url = "/general/pages/multiton.html";
                node.multiId = [];
                node.multiId.push(classId);
                node.leafName = classCnName;
            } else {
                node.url = "/general/pages/dynamicLoad.html";
                node.classId = classId;
                node.leafName = classCnName;
            };

            return node;
        }

        function listMenu() {
            var htm = [];
            var left_menu = $("#sidebar-left ul:first");

            for (var key in g_menutree) {
                if (key.indexOf("branch") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        htm.push('<li>');
                        htm.push('    <a class="dropmenu" href="#" onclick="$(this).parent().find(\'ul\').slideToggle(&quot;normal&quot;, MENU.switchMenu(this));">');
                        htm.push('        <i class="icon-folder-close-alt"></i>');
                        htm.push('        <span class="hidden-tablet">', g_menutree[key][i].menuName, '</span>');
                        htm.push('        <span class="label label-important">', g_menutree[key][i].leaf.length, '</span>');
                        htm.push('    </a>');
                        htm.push('    <ul style="display:block;">');

                        for (var j = 0; j < g_menutree[key][i].leaf.length; j++) {
                            htm.push('<li>');
                            htm.push('    <a class="submenu" href="#" skip="', g_menutree[key][i].leaf[j].skip, '" onclick="MENU.touchMenu(&quot;', g_menutree[key][i].leaf[j].menuId, '&quot;);MENU.initBreadcrumb(&quot;', g_menutree[key][i].menuName, '&quot;,&quot;', g_menutree[key][i].leaf[j].leafName, '&quot;);MENU.activeMenu(this);">');
                            htm.push('        <i class="icon-file-alt"></i>');
                            htm.push('        <span class="hidden-tablet">', g_menutree[key][i].leaf[j].leafName, '</span>');
                            htm.push('    </a>');
                            htm.push('</li>');

                        }
                        htm.push('    </ul>');
                        htm.push('</li>');
                    }
                } else if (key.indexOf("leaf") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        for (var i = 0; i < g_menutree[key].length; i++) {
                            htm.push('<li>');
                            htm.push('    <a href="#" skip="', g_menutree[key][i].skip, '" onclick="MENU.touchMenu(&quot;', g_menutree[key][i].menuId, '&quot;);MENU.initBreadcrumb(&quot;&quot;,&quot;', g_menutree[key][i].leafName, '&quot;);MENU.activeMenu(this);">');
                            htm.push('        <i class="icon-bar-chart"></i>&nbsp;');
                            htm.push('        <span class="hidden-tablet">', g_menutree[key][i].leafName, '</span>');
                            htm.push('    </a>');
                            htm.push('</li>');
                        }
                    }
                }
            }

            left_menu.empty().append(htm.join(""));

            setTimeout(function() {
                $(".dropmenu").parent().find("ul").hide();

                var skip_name = $.cookie("skip");
                if (skip_name != null) {
                    $("a", left_menu).each(function() {
                        if ($(this).attr("skip") == skip_name) {
                            if ($(this).hasClass("submenu")) {
                                $(this).parent().parent().prev().click();
                            }
                            $(this).click();
                        }
                    });
                    $.cookie("skip", null);
                } else {
                    if ($("#sidebar-left a:first").hasClass("dropmenu")) {
                        $("#sidebar-left a:first").click();
                        $(".submenu :first").click();
                    } else {
                        $("#sidebar-left a:first").click();
                    }
                }

            }, 50);
        }

        function touchMenu(menuId) {
            for (var key in g_menutree) {
                if (key.indexOf("branch") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        for (var j = 0; j < g_menutree[key][i].leaf.length; j++) {
                            if (g_menutree[key][i].leaf[j].menuId == menuId) {
                                g_curMenu = g_menutree[key][i].leaf[j];
                                document.getElementById('mainFrame').src = g_menutree[key][i].leaf[j].url + "?_=" + new Date().getTime();
                                return;
                            }
                        }
                    }
                } else if (key.indexOf("leaf") != -1) {
                    for (var i = 0; i < g_menutree[key].length; i++) {
                        if (g_menutree[key][i].menuId == menuId) {
                            g_curMenu = g_menutree[key][i];
                            document.getElementById('mainFrame').src = g_menutree[key][i].url + "?_=" + new Date().getTime();
                            return;
                        }
                    }
                }
            }
        }

        function initBreadcrumb(mainmenu, submenu) {
            if ("" == $.trim(mainmenu)) {
                $("#mainmenu").html(mainmenu);
                $("#submenu").html(submenu);
                $(".icon-angle-right").hide();
            } else {
                $("#mainmenu").html(mainmenu);
                $("#submenu").html(submenu);
                $(".icon-angle-right").show();
            }
        }

        function switchMenu(obj) {
            setTimeout(function() {
                $(".dropmenu.menu-visible").each(function() {
                    $(this).removeClass("menu-visible");
                });

                if ($(obj).next().is(':visible')) {
                    $(obj).addClass("menu-visible");
                } else {
                    $(obj).removeClass("menu-visible");
                }

                $(".dropmenu").each(function() {
                    if (!$(this).hasClass('menu-visible')) {
                        $(this).next().css("display", "none");
                    }
                });
            }, 200);
            setTimeout(function() {
                top.autoHeight();
            }, 1000);
        }

        function activeMenu(obj) {
            $('ul.main-menu li a').each(function() {
                $(this).parent().removeClass('active');
            });

            $('ul.main-menu li ul li a').each(function() {
                $(this).parent().removeClass('active');
            });

            $(obj).parent().addClass("active");
            $(obj).blur();
        }

        return {
            loadMenu : loadMenu,
            listMenu : listMenu,
            touchMenu : touchMenu,
            initBreadcrumb : initBreadcrumb,
            switchMenu : switchMenu,
            activeMenu : activeMenu
        }

    }())