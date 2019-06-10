<!DOCTYPE html>
<html lang="en">

<head>
	<!-- start: Meta -->
	<meta charset="utf-8">
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<meta http-equiv="expires" content="0" />
	<meta http-equiv="progma" content="no-cache" />
	<meta http-equiv="cache-control" content="no-cache" />
	<meta http-equiv="Expires" content="0" />
	<title></title>
	<!-- end: Meta -->

	<!-- start: Mobile Specific -->
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<!-- end: Mobile Specific -->

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

	<!-- start: Favicon -->
	<link rel="shortcut icon" href="img/favicon.ico">
	<!-- end: Favicon -->
	<script src="js/jquery-1.9.1.min.js"></script>
	<script src="js/js-css-loader.js"></script>
	<script defer="defer" type="text/javascript">
		var versionNo = "";
		$.get("conf/web-version.txt", null, function (data) {
			versionNo = data;

			loadMainCss("bootstrap.min.css");
			loadMainCss("bootstrap-responsive.min.css");
			loadMainCss("style.css");
			loadMainCss("style-responsive.css");

			loadMainJs("application.js");
			loadMainJs("display.js");
			loadMainJs("bootstrap.min.js");
			loadMainJs("bootstrapmodal.js");
			loadMainJs("main.js", 500);
		});
	</script>
</head>

<body class="bg-white">
	<!-- start: Header -->
	<div class="navbar">
		<div class="navbar-inner">
			<div class="container-fluid">
				<a class="btn btn-navbar" data-toggle="collapse" data-target=".top-nav.nav-collapse,.sidebar-nav.nav-collapse">
					<span class="icon-bar"></span> <span class="icon-bar"></span> <span class="icon-bar"></span> </a>
				<a class="brand" href="#"><span></span></a>

				<!-- start: Header Menu -->
				<div class="nav-no-collapse header-nav">
					<ul class="nav pull-right">
						<!-- start: User Dropdown -->
						<li class="dropdown">
							<a class="btn dropdown-toggle" data-toggle="dropdown" href="#"> <i
									class="halflings-icon white user"></i>&nbsp;<span id="username"></span> <span class="caret"></span>
							</a>
							<ul class="dropdown-menu">
								<li>
									<a href="#" onclick="System.logout(1)"><i class="halflings-icon off"></i> Logout</a>
								</li>
							</ul>
						</li>
						<!-- end: User Dropdown -->
					</ul>
				</div>
				<!-- end: Header Menu -->

			</div>
		</div>
	</div>
	<!-- start: Header -->

	<div class="container-fluid-full">
		<div class="row-fluid">

			<!-- start: Main Menu -->
			<div id="sidebar-left" class="span2">
				<div class="nav-collapse sidebar-nav">
					<ul class="nav nav-tabs nav-stacked main-menu"></ul>
				</div>
			</div>
			<!-- end: Main Menu -->

			<!-- start: Content -->
			<div id="content" class="span10 mleft0">
				<ul id="breadcrumb" class="breadcrumb">
					<li></li>
					<li class="right-link">
						<i class="icon-share-alt"></i><a href="#" onclick="backToMain()">返回工作主页</a>
					</li>
				</ul>
				<iframe id="contentMainFrame" name="mainFrame" allowtransparency="true" width="100%" height="100%"
					frameborder="0" scrolling="no"></iframe>
			</div>
			<!--/.fluid-container-->
			<!-- end: Content -->
		</div>
		<!--/#content.span10-->
	</div>
	<!--/fluid-row-->

	<div class="modal hide fade" id="myModal">
		<div class="modal-header">
			<button type="button" class="close" data-dismiss="modal">
				×
			</button>
			<h3></h3>
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

	<div id="progress" class="progress progress-warning progress-striped active" style="display: none;">
		<div class="bar">Loading...</div>
	</div>
</body>

</html>