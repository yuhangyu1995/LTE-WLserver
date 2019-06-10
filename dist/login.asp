<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

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

	<style type="text/css">
		body {
			background: #EEE !important;
		}
	</style>
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

			loadMainJs("common.js");
			loadMainJs("md5.js");
			loadMainJs("application.js");
			loadMainJs("display.js");
			loadMainJs("bootstrap.min.js");
			loadMainJs("bootstrapmodal.js");
			loadMainJs("login.js", 500);
		});
	</script>
</head>

<body>
	<div class="container-fluid-full">
		<div class="row-fluid">
			<div class="row-fluid">
				<div class="login-box">
					<div class="icons"></div>
					<h2>Login to your account</h2>
					<form class="form-horizontal" method="post">
						<div class="input-prepend" title="Username">
							<span class="add-on"><i class="halflings-icon user"></i></span>
							<input class="input-large span10" name="username" id="username" type="text" placeholder="type username" />
						</div>
						<div class="clearfix"></div>

						<div class="input-prepend" title="Password">
							<span class="add-on"><i class="halflings-icon lock"></i></span>
							<input class="input-large span10" name="password" id="password" type="password"
								placeholder="type password" />
						</div>
						<div class="clearfix"></div>

						<div class="button-login">
							<button type="button" class="btn btn-primary" onclick="checkLogin();">
								Login
							</button>
						</div>
						<div class="clearfix"></div>
					</form>
				</div>
			</div>
		</div>
	</div>

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
</body>

</html>