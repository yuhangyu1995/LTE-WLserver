var ModalCommon = ( function() {

        var modal_callback = function() {
        };

        function cover() {
            if ($("#backdrop1").length <= 0) {
                $('<div id = "backdrop1" class="modal-backdrop fade in"></div>').appendTo("body");
            }
            $("#backdrop1").show();
        }

        function delCover() {
            $("#backdrop1").remove();
        }

        function progress(msg) {
            cover();
            if (msg) {
                $(".bar", $("#progress")).html(msg);
            }
            $("#progress").show();
        }

        function switchMsg(msg) {
            if (msg) {
                $(".bar", $("#progress")).html(msg);
            }
        }

        function delProgress() {
            delCover();
            $("#progress").hide();
            $(".bar", $("#progress")).html("加载中...");
        }

        function open(title, inform, msgType, callback) {
            if (callback) {
                modal_callback = callback;
            } else {
                modal_callback = function() {
                };
            }

            $(".modal-title", top.$(".modal-header")).html("" == $.trim(title) ? top.DISPLAY.Main.Notice : title);
            $("p", top.$(".modal-body")).html(inform);

            $("#confirmBtn", top.$(".modal-footer")).html("确定");

            if (msgType === "choice") {
                $("#closeBtn").show();
                $("#confirmBtn").show();
            } else if (msgType === "error") {
                $("#closeBtn").hide();
                $("#confirmBtn").show();
            } else {
                $("#closeBtn").hide();
                $("#confirmBtn").show();
            }
            $("#myModal").modal({
                backdrop : 'static'
            });
        }

        function open2(title, inform, msgType, callback, confirmBtnText, closeBtnText) {
            if (callback) {
                modal_callback = callback;
            } else {
                modal_callback = function() {
                };
            }

            $('button[class="close"]', $('#myModal')).click();
            $(".modal-backdrop").remove();
            setTimeout(function() {
                $(".modal-title", $(".modal-header")).html("" == $.trim(title) ? top.DISPLAY.Main.Notice : title);
                $("p", $(".modal-body")).html(inform);

                $("#confirmBtn", $(".modal-footer")).html("" == $.trim(confirmBtnText) ? "确定" : confirmBtnText);
                $("#closeBtn", $(".modal-footer")).html("" == $.trim(closeBtnText) ? "取消" : closeBtnText);

                if (msgType === "choice") {
                    $("#closeBtn").show();
                    $("#confirmBtn").show();
                } else if (msgType === "error") {
                    $("#closeBtn").hide();
                    $("#confirmBtn").show();
                } else {
                    $("#closeBtn").hide();
                    $("#confirmBtn").show();
                }
                $("#myModal").modal({
                    backdrop : 'static'
                });
            }, 500);
        }

        function submit() {
            $("#myModal").modal("hide");
            modal_callback(1);
        }

        function cancel() {
            $("#myModal").modal("hide");
            modal_callback(-1);
        }

        return {
            progress : progress,
            delProgress : delProgress,
            open : open,
            open2 : open2,
            submit : submit,
            cancel : cancel,
            delCover : delCover,
            cover : cover,
            switchMsg : switchMsg
        };
    }());
