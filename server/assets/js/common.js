// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2023-04-05
// @Version: 1.7.0

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


const themeVarLight = `
:root {
    --label-color: #979d9e;
    --navbar-color: #555555;
    --cache-color: #D9E4DD;
    --cpu-memory-title-color: #CDC9C3;
    --temperature-color: #FBF7F0;
    --ip-color: #D9E4DD;
    --time-color: #CDC9C3;
    --uptime-color: #FBF7F0;
    --cpu-color: #D9E4DD;
    --memory-color: #D9E4DD;
    --real-memory-color: #D9E4DD;
    --swap-color: #D9E4DD;
    --disk-color: #CDC9C3;
    --box-bg-color: #eaebe9; /* #E8EAE6 */
    --backdrop-color: #363636;

    --guage-font-color: black;
    --guage-stops-color-1: #D9E4DD;
    --guage-stops-color-5: #CDC9C3;
    --guage-stops-color-9: #919191;

    --net-in-color: #D9E4DD;
    --net-out-color: #CDC9C3;
    --net-grid-line-color: #e6e6e6;
    --net-line-color: #ccd6eb;

    --box-radius: 15px;
}
`;
const themeVarDark = `
::-webkit-scrollbar-track {
    background: #3f3f3f;
    -webkit-box-shadow: inset 0 0 5px #3f3f3f;
}
::-webkit-scrollbar-thumb {
    background: #7c7c7c;
    -webkit-box-shadow: inset 0 0 10px #7c7c7c;
}

body{
    background-color: #2d2d2d;
    color: #c9d1d9;
}
input {
    background-color: #3b3b3b !important;
    border-color: #3b3b3b !important;
    color: #c9d1d9 !important;
}
input:focus {
    border: 1px solid #6b6b6b !important;
}

:root {
    --label-color: #959c9c;
    --navbar-color: #474747;
    --cache-color: #414141;
    --cpu-memory-title-color: #5e5e5e;
    --temperature-color: #4b4b4b;
    --ip-color: #414141;
    --time-color: #5e5e5e;
    --uptime-color: #4b4b4b;
    --cpu-color: #414141;
    --memory-color: #414141;
    --real-memory-color: #414141;
    --swap-color: #414141;
    --disk-color: #5e5e5e;
    --box-bg-color: #373938;
    --backdrop-color: #363636;

    --guage-font-color: #b9bebe;
    --guage-stops-color-1: #626464;
    --guage-stops-color-5: #878b8b;
    --guage-stops-color-9: #b6b6b6;

    --net-in-color: #414141;
    --net-out-color: #5e5e5e;
    --net-grid-line-color: #3d3d3d;
    --net-line-color: #575757;

    --box-radius: 15px;
}

.modal-content {
    background-color: #252525;
    color: #c9d1d9;
}
.dark-bg {
    background-color: #505050 !important;
}
.spinner > div {
    background-color: #c9d1d9;
}
.navbar-toggler {
    color: #505050 !important;
}

#login-tips {
    color: #c9d1d9;
}
#pimodel {
    color: #c9d1d9;
}
#command-btns li img:hover{
    border: 1px solid #606060 !important;
}
`;
var theme = $("meta[name='theme']").attr('content');
$(document).ready(function() {
    if (theme == "dark" || window.matchMedia('(prefers-color-scheme: dark)').matches) {
        $("#theme-var").text(themeVarDark);
        $("#modal-close-btn").addClass("btn-close-white");
        $("footer").eq(0).addClass("border-secondary");
        $("meta[name='theme-color']").attr('content', '#474747');
        if ($("#favicon").text() == "linux.ico") {
            $("#device-photo").addClass("inverted");
            $("#icon").attr("href", "favicons/linux_light.ico");
            $("#shortcut-icon").attr("href", "favicons/linux_light.ico");
        }
        $.getScript('js/index.js', function() {});
    } else if (theme == "light" || window.matchMedia('(prefers-color-scheme: light)').matches) {
        $("#theme-var").text(themeVarLight);
        $("#modal-close-btn").removeClass("btn-close-white");
        $("footer").eq(0).removeClass("border-secondary");
        $("meta[name='theme-color']").attr('content', '#555555');
        $("#device-photo").removeClass("inverted");
        if ($("#favicon").text() == "linux.ico") {
            $("#icon").attr("href", "favicons/linux.ico");
            $("#shortcut-icon").attr("href", "favicons/linux.ico");
        } else {
            $("#icon").attr("href", "favicons/raspberrypi.ico");
            $("#shortcut-icon").attr("href", "favicons/raspberrypi.ico");
        }
        $.getScript('js/index.js', function() {});
    }
});

let media = window.matchMedia('(prefers-color-scheme: dark)');
let callback = (e) => {
    let prefersDarkMode = e.matches;
    if (prefersDarkMode) {
        $("#theme-var").text(themeVarDark);
        $("#modal-close-btn").addClass("btn-close-white");
        $("footer").eq(0).addClass("border-secondary");
        $("meta[name='theme-color']").attr('content', '#474747');
        if ($("#favicon").text() == "linux.ico") {
            $("#device-photo").addClass("inverted");
            $("#icon").attr("href", "favicons/linux_light.ico");
            $("#shortcut-icon").attr("href", "favicons/linux_light.ico");
        }
        $.getScript('js/index.js', function() {});
    } else {
        $("#theme-var").text(themeVarLight);
        $("#modal-close-btn").removeClass("btn-close-white");
        $("footer").eq(0).removeClass("border-secondary");
        $("meta[name='theme-color']").attr('content', '#555555');
        $("#device-photo").removeClass("inverted");
        if ($("#favicon").text() == "linux.ico") {
            $("#icon").attr("href", "favicons/linux.ico");
            $("#shortcut-icon").attr("href", "favicons/linux.ico");
        } else {
            $("#icon").attr("href", "favicons/raspberrypi.ico");
            $("#shortcut-icon").attr("href", "favicons/raspberrypi.ico");
        }
        $.getScript('js/index.js', function() {});
    }
};
if (typeof media.addEventListener === 'function') {
    media.addEventListener('change', callback);
} else if (typeof media.addEventListener === 'function') {
    media.addEventListener(callback);
}

$('.dropdown-item').on('click',function() {
    $('.navbar-collapse').collapse('hide');
});
$('#logout').on('click',function() {
    $('.navbar-collapse').collapse('hide');
});

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