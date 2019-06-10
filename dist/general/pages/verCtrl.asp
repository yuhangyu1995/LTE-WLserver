<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta http-equiv="Pragma" content="no-cache" />
		<meta http-equiv="Cache-Control" content="no-cache" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<title></title>
		<style type="text/css"></style>
		<!-- start: CSS -->
		<link id="bootstrap-style" href="../../css/bootstrap.min.css" rel="stylesheet">
		<link href="../../css/bootstrap-responsive.min.css" rel="stylesheet">
		<link id="base-style" href="../../css/style.css" rel="stylesheet">
		<link id="base-style-responsive" href="../../css/style-responsive.css" rel="stylesheet">
		<!-- end: CSS -->

		<!-- The HTML5 shim, for IE6-8 support of HTML5 elements -->
		<!--[if lt IE 9]>
		<script src="../../js/html5shiv.js"></script>
		<link id="ie-style" href="../../css/ie.css" rel="stylesheet">
		<![endif]-->

		<!--[if IE 9]>
		<link id="ie9style" href="../../css/ie9.css" rel="stylesheet">
		<![endif]-->
		<script type="text/javascript" src="../../js/jquery-1.9.1.min.js" ></script>
		<script defer="defer" type="text/javascript"></script>
	</head>
	<body id="main_body" style="background: #FFFFFF">
		<div class="row-fluid">
			<div class="span12">
    				<div class="well" style="color: red;">
					固件升级注意事项：<br>
					1、固件升级过程中，请确保设备不会断电。<br>
					2、根据固件升级需求，优先执行各制式固件升级，确定各制式固件升级成功后，执行主控固件升级。
				</div>
			</div>
		</div>
		<div class="row-fluid sortable">
			<div class="box span6">
				<div class="box-header">
					<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="SoftwareUpdate"></span></h2>
					<div class="box-icon">
						<a href="#" class="btn-upgrade" data-rel="tooltip" data-original-title=""><i class="halflings-icon upload"></i></a>
						<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>
					</div>
				</div>
				<div class="box-content" style="height: 94px">
					<form class="form-horizontal" method="post" name="UpLoadConfigForm" id="UpLoadConfigForm" action="/goform/uploadConfigFile" enctype="multipart/form-data" onSubmit="sendMsg()">
						<fieldset>
							<div class="control-group">
								<label id="AUTOUpload" class="control-label"></label>
								<div class="controls">
									<div class="uploader">
										<input type="file" id="fileBrowse" name="edtUpLoadFilePath_auto" onchange="document.getElementById('textfield').value=this.value">
										<span class="filename" style="-webkit-user-select: none;">No file selected</span><span class="action" style="-webkit-user-select: none;">Choose File</span>
									</div>
								</div>
							</div>
							<input type='hidden' name='textfield' id='textfield'/>
							<input id="Hidden1" name="type" type="hidden" value="1" />
						</fieldset>
					</form>
				</div>
			</div>

			<div class="box span6">
				<div class="box-header">
					<h2><i class="halflings-icon align-justify"></i><span class="break"></span><span id="VersionActive"></span></h2>
					<div class="box-icon">
						<a href="#" class="btn-activeversion" data-rel="tooltip" data-original-title=""><i class="halflings-icon random"></i></a>
						<a href="#" class="btn-minimize" data-rel="tooltip" data-original-title=""><i class="halflings-icon chevron-up"></i></a>
					</div>
				</div>
				<div class="box-content">
					<table class="table">
						<tr>
							<td id="CurrentVersion"></td>
							<td id="currentVersion"></td>
						</tr>
						<tr>
							<td id="BakVersion"></td>
							<td id="backupVersion"></td>
						</tr>
					</table>
				</div>
			</div>
		</div>
	</body>
	<script src="../../js/verCtrl.js"></script>
	<script src="../../js/jquery-migrate-1.0.0.min.js"></script>
	<script src="../../js/jquery.chosen.min.js"></script>
	<script src="../../js/bootstrap.min.js"></script>
	<script src="../../js/jquery.uniform.min.js"></script>
</html>