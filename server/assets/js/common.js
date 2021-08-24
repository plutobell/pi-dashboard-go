// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-24
// @Version: 1.4.1

window.oncontextmenu=function(){return false;}
window.onkeydown = window.onkeyup = window.onkeypress = function (event) {
    if (event.keyCode === 123) {
        event.preventDefault();
        window.event.returnValue = false;
    }
}
window.addEventListener('keydown', function (event) {
    if (event.ctrlKey) {
        event.preventDefault();
    }
})


function getCookie(name) {
    var cookieValue = null;
    if (document.cookie && document.cookie !== '') {
        var cookies = document.cookie.split(';');
        for (var i = 0; i < cookies.length; i++) {
            var cookie = jQuery.trim(cookies[i]);
            // Does this cookie string begin with the name we want?
            if (cookie.substring(0, name.length + 1) === (name + '=')) {
                cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                break;
            }
        }
    }
    return cookieValue;
}

function csrfSafeMethod(method) {
    // 这些HTTP方法不要求携带CSRF令牌。test()是js正则表达式方法，若模板匹配成功，则返回true
    return (/^(GET|HEAD|OPTIONS|TRACE)$/.test(method));
}

function csrfAddToAjaxHeader() {
    var csrftoken = getCookie('cf_sid');

    return {
        beforeSend: function(xhr, settings) {
            if (!csrfSafeMethod(settings.type) && !this.crossDomain) {
                xhr.setRequestHeader("X-XSRF-TOKEN", csrftoken);
            }
        }
    }
}