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
			
			loadJs("jquery-migrate-1.0.0.min.js");
			loadJs("jquery.chosen.min.js");
			loadJs("bootstrap.min.js");
			loadJs("jquery.uniform.min.js");
			loadJs("pk.js", 500);
		</script>
	</head>
	<body class="bg-white">
		<div class="row-fluid sortable" id="switch_one">
			<div class="box span6">
				<div class="box-header">
					<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="pkFile"></span></h2>
					<div class="box-icon">
						<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>
					</div>
				</div>
				<div class="box-content">
					<div class="control-group">
						<div class="controls">
							<button class="btn btn-primary noty" onclick="pkFileDownload()" style="margin-left: 35%;">
								<i class="halflings-icon white bell"></i><span id="pkFileDownload"></span>
							</button>
							<button class="btn btn-primary noty" onclick="pkFileClear()">
								<i class="halflings-icon white bell"></i><span id="pkFileClear"></span>
							</button>
						</div>
					</div>
				</div>
				<form method="post" id="DownLoadForm" name="DownLoadForm" action="/goform/pkImsiFile" >
					<input type="hidden"  id="apdownfile" name="edtDownLoadCfgFileName" value="" />
					<input id="path" name="path" type="hidden" value=""/>
				</form>
			</div>
		</div>
	</body>
</html>