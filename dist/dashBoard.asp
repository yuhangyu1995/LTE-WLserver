<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">

<html>
    <head>
        <!-- start: Meta -->
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <meta http-equiv="Pragma" content="no-cache" />
        <meta http-equiv="Cache-Control" content="no-cache" />
        <meta http-equiv="Expires" content="0" />
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
        <title></title>
        <!-- end: Meta -->

        <!-- start: Mobile Specific -->
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <!-- end: Mobile Specific -->

        <style type="text/css"></style>
        <!-- start: CSS -->
        <section id="cssArea"></section>
        <!-- end: CSS -->

        <!-- The HTML5 shim, for IE6-8 support of HTML5 elements -->
        <!--[if lt IE 9]>
        <script src="js/html5shiv.js"></script>
        <link id="ie-style" href="css/ie.css" rel="stylesheet">
        <![endif]-->

        <!--[if IE 9]>
        <link id="ie9style" href="css/ie9.css" rel="stylesheet">
        <![endif]-->

        <script src="js/jquery-1.9.1.min.js"></script>
        <script src="js/js-css-loader.js"></script>
        <script defer="defer" type="text/javascript">
            var T = null;
            var versionNo = "";
            $.get("conf/web-version.txt", null, function(data) {
                versionNo = data;

                loadMainCss("bootstrap.min.css");
                loadMainCss("bootstrap-responsive.min.css");
                loadMainCss("style.css");
                loadMainCss("style-responsive.css");
                loadMainCss("animate.min.css");
                loadMainCss("dashBoard.css");

                loadMainJs("common.js");
                loadMainJs("jquery.flot.js");
                loadMainJs("assist1.js");
                loadMainJs("sysAssist.js");
                loadMainJs("jquery-migrate-1.0.0.min.js");
                loadMainJs("bootstrap.min.js");
                loadMainJs("jquery.uniform.min.js");
                loadMainJs("jquery.chosen.min.js");
                loadMainJs("jquery.knob.modified.js");
                loadMainJs("jquery.sparkline-2.1.2.js");
                loadMainJs("counter.js");
                loadMainJs("display.js");
                loadMainJs("bootstrapmodal.js");
                loadMainJs("excanvas.min.js");
                loadMainJs("md5.js");
                loadMainJs("dashBoard.js", 500);

                setTimeout(function() {
                    $(document).ready(function() {
                        T = T_ZH;
                        initUser();
                        onlineorNo();

                        MainTemplate.init();
                        MainTemplate.isMoreInfoShow();
                        MainTemplate.showPaInfo();
                        MainTemplate.FINAL_PA_INFO();
                        MainTemplate.getTotalInfo();

                        bindBtnEvent();
                        initModuleSwitch();
                        initSysFlashImsiMgmt();
                        getImsiLog();
                        getFlowData();

                        setTimeout(function() {
                            MainTemplate.showPaInfo();
                            getImsiLog();
                            setTimeout(arguments.callee, 10 * 1000);
                        }, 10 * 1000);

                        setTimeout(function() {
                            MainTemplate.getTotalInfo();
                            setTimeout(arguments.callee, 60 * 1000);
                        }, 60 * 1000);

                        setTimeout(function() {
                            getFlowData();
                            setTimeout(arguments.callee, 30 * 60 * 1000);
                        }, 30 * 60 * 1000);
                    });
                }, 2000);
            });
        </script>
    </head>
    <body>
        <!-- start: Header -->
        <div class="navbar">
            <div class="navbar-inner">
                <div class="container-fluid">
                    <a class="btn btn-navbar" data-toggle="collapse" data-target=".top-nav.nav-collapse,.sidebar-nav.nav-collapse"> 
                        <span class="icon-bar"></span> 
                        <span class="icon-bar"></span> 
                        <span class="icon-bar"></span> 
                    </a>
                    <a class="brand" href="#"><span></span></a>

                    <!-- start: Header Menu -->
                    <div class="nav-no-collapse header-nav">
                        <ul class="nav pull-left mleft-minus-2percent">
                            <li id="serial_number"></li>
                            <li id="gps_position"></li>
                        </ul>
                        <ul class="nav pull-right">
                            <!-- start: User Dropdown -->
                            <li class="dropdown">
                                <a class="btn dropdown-toggle" data-toggle="dropdown" href="#"> 
                                    <i class="halflings-icon white user"></i>&nbsp; 
                                    <span id="username"></span> 
                                    <span class="caret"></span> 
                                </a>
                                <ul class="dropdown-menu">
                                    <li>
                                        <a href="#" onclick="System.logout(1)"><i class="halflings-icon off"></i>&nbsp; 登出</a>
                                    </li>
                                </ul>
                            </li>
                            <!-- end: User Dropdown -->
                        </ul>
                    </div>
                </div>
            </div>
        </div>
        <!-- start: Header -->

        <div class="container-fluid-full">
            <div class="row-fluid">
                <!-- start: Content -->
                <div id="content" class="span10">

                    <ul class="breadcrumb">
                        <li>
                            <i class="icon-home"></i>
                            <a href="#">工作主页</a>
                        </li>
                    </ul>

                    <div id="pa_module" class="row-fluid"></div>

                    <div class="row-fluid cursor-menu">
                        <div id="reboot_module" class="span3 widget orange" onTablet="span3" onDesktop="span3">
                            <h2 class="h2-no-shadow"><span class="glyphicons leaf"><i></i></span>Reboot Choice</h2>
                            <hr class="hr-margin-small">

                            <div class="control-group">
                                <div class="controls">
                                    <label class="radio inline">
                                        <input type="radio" name="scanSelfOptDeviceSwitch" value="1">
                                        Enable </label>
                                    <label class="radio inline">
                                        <input type="radio" name="scanSelfOptDeviceSwitch" value="0">
                                        Disable </label>
                                </div>
                            </div>

                            <div class="control-group">
                                <nobr>
                                    <div class="controls">
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="GSM-M">
                                            GSM-M </label>
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="GSM-U">
                                            GSM-U </label>
                                    </div>
                                </nobr>
                            </div>

                            <div class="control-group">
                                <nobr>
                                    <div class="controls">
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="WCDMA">
                                            WCDMA </label>
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="L-TDD-1900">
                                            L-TDD-1900 </label>
                                    </div>
                                </nobr>
                            </div>

                            <div class="control-group">
                                <nobr>
                                    <div class="controls">
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="L-TDD-2300">
                                            L-TDD-2300 </label>
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="L-TDD-2600">
                                            L-TDD-2600 </label>
                                    </div>
                                </nobr>
                            </div>

                            <div class="control-group">
                                <nobr>
                                    <div class="controls">
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="L-FDD-U">
                                            L-FDD-U </label>
                                        <label class="checkbox inline">
                                            <input type="checkbox" name="scanSelfOptDeviceInd" value="L-FDD-T">
                                            L-FDD-T </label>
                                    </div>
                                </nobr>
                            </div>

                            <div class="control-group">
                                <label class="control-label" for="scanSelfOptPeriod">采集检测周期(分钟)：</label>
                                <div class="controls inline">
                                    <input type="text" id="scanSelfOptPeriod"></div>
                                    <div>
                                        <button type="submit" class="btn btn-primary" onclick="saveScanSelfOptPeriod()">
                                           <i class="halflings-icon ok white"></i>
                                        </button>
                                    </div>
                            </div>
                            <span class="footer-comment">注：以上选中模块不上号时整机下电重启，配合采集检测周期使用</span>
                        </div>
                       

                        <div id="record_module" class="span5_5 widget blue" onTablet="span4" onDesktop="span5_5">
                            <div>
                                <h2 class="h2-no-shadow cursor-pointer imsi_log"><span class="glyphicons show_big_thumbnails"><i></i></span>Data Collection </h2>
                                <hr class="hr-margin-small">
                                <ul class="messagesList"></ul>
                            </div>
                        </div>

                        <div id="version_module" class="span3_5 widget green sparkLineStats pull-right" onTablet="span5" onDesktop="span3_5">
                            <h2 class="h2-no-shadow push-left"><span class="glyphicons justify"><i></i></span>Version Manager</h2>
                            <hr class="hr-margin-small">
                            <ul id="version_value"></ul>
                        </div>
                    </div>

                    <div id="imsi_count" class="row-fluid hideInIE8 circleStats"></div>

                </div><!--/.fluid-container-->
                <!-- end: Content -->
            </div><!--/#content.span10-->
        </div><!--/fluid-row-->

        <div class="modal hide fade" id="myModal">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal">
                    ×
                </button>
                <h3 class="modal-title"></h3>
            </div>
            <div class="modal-body">
                <p></p>
            </div>
            <div class="modal-footer">
                <a href="#" id="closeBtn" class="btn btn-success" data-dismiss="modal" onclick="ModalCommon.cancel();">Close</a>
                <a href="#" id="confirmBtn" class="btn btn-primary" onclick="ModalCommon.submit();">Confirm</a>
            </div>
        </div>

        <div class="clearfix"></div>

        <footer id="footer">
            <p>
                <span></span>
            </p>
        </footer>
    </body>
</html>