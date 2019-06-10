<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
﻿
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta http-equiv="Pragma" content="no-cache" />
		<meta http-equiv="Cache-Control" content="no-cache" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<title></title>
		<style type="text/css"></style>
		<!-- start: CSS -->
		<section id="cssArea"></section>
		<!-- end: CSS -->

		<!-- The HTML5 shim, for IE6-8 support of HTML5 elements -->
		<!--[if lt IE 9]>
		<script src="../../js/html5shiv.js"></script>
		<link id="ie-style" href="../../css/ie.css" rel="stylesheet">
		<![endif]-->

		<!--[if IE 9]>
		<link id="ie9style" href="../../css/ie9.css" rel="stylesheet">
		<![endif]-->
		<script src="../../js/jquery-1.9.1.min.js" ></script>
		<script src="../../js/js-css-loader.js"></script>
		<script defer="defer" type="text/javascript">
			
            loadCss("bootstrap.min.css");
            loadCss("bootstrap-responsive.min.css");
            loadCss("style.css");
            loadCss("style-responsive.css");
			
			loadJs("networkDiagnosis.js", 500);
			
		</script>
	</head>
	<body class="bg-white">
		<div class="row-fluid sortable">
			<div class="span12">
				<h1>
					<span id="netDiagTitle"></span>
					<div class="control-group success" style="float: right">
						<label id="ipFilterLabel" class="control-label" style="display: inline-block;"></label>
						<div class="controls" style="display: inline-block;">
							<input id="ipFilter" type="text" style="margin-bottom: 0px;" value="">
							<button class="btn btn-primary" id="startDiagBtn"></button>
						</div>
					</div>
				</h1>
				<div class="priority medium" style="display: inline-block;width: 100%;">
					<span id="diagResultLabel" style="cursor: pointer;"></span>
					<span id="diagResult" style="display: none;"></span>
				</div>
				<div id="diagDetail">
					<div class="task medium last">
						<div class="desc">
					     <div class="title"></div>
					     <div class="diagnosisContent">没有记录</div>
					    </div>
					    <div class="time">
					     <div class="date"></div>
					     <div class="diagnosisTime"></div>
					    </div>
					</div>
				</div>
			<div>
		<div>
	</body>
</html>