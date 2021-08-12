// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modify: 2021-08-12
// @Version: 1.3.0

window.oncontextmenu=function(){return false;}

$("form").keyup(function(event){
    if(event.keyCode == 13){
        $("#login-btn").trigger("click");
    }
});

$("#login-btn").click(function(){
    $("#login-btn").attr("disabled", true);
    var username = $("#username").val();
    var password = $("#password").val();
    var json = {
        "username": username,
        "password": password,
    };
    if (username == "" || password == "") {
        $("#login-tips").text("Username or password is empty")
        $("#login-btn").attr("disabled", false);
    } else {
        $.post('/api/login', JSON.stringify(json), function(result){
            if (result.status == true) {
                $("#login-tips").text("")
                $(window).attr('location','/');
            } else if (result.status == false) {
                $("#login-tips").text("Wrong credentialss")
                $("#login-btn").attr("disabled", false);
            }
        }).fail(function() {
            $("#login-tips").text("Unknown error")
            $("#login-btn").attr("disabled", false);
        });
    }

});