<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<meta http-equiv="expires" content="0" />
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
		<script type="text/javascript">
		
			loadCss("bootstrap.min.css");
            loadCss("bootstrap-responsive.min.css");
            loadCss("style.css");
            loadCss("style-responsive.css");
			
			loadJs("common.js");
			loadJs("table.js");
			loadJs("jquery-migrate-1.0.0.min.js");
			loadJs("jquery.chosen.min.js");
			loadJs("bootstrap.min.js");
			loadJs("diagnosticsInfo.js", 500);
		</script>
	</head>
	<body class="bg-white">
		<div id="y_scroll" style="width:100%; height:100%; overflow-x:visible; overflow-y:visible; visibility:visible;">
			<div class="row-fluid sortable">
				<div class="box span12">
					<div class="box-header" data-original-title>
						<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="Diagnosticslog"></span></h2>
						<div class="box-icon">
							<a href="#" class="btn-stoprefresh" data-rel="tooltip" data-original-title=""><i class="halflings-icon off"></i></a>
							<a href="#" class="btn-refresh" data-rel="tooltip" data-original-title=""><i class="halflings-icon refresh"></i></a>
							<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>
						</div>
					</div>
					<div id="container" class="box-content" style="overflow: auto;"></div>
				</div>
			</div>
		</div>
	</body>
</html>
