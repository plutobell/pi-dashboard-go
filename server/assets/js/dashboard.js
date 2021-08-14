// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-14
// @Version: 1.3.3

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

unScroll();

$(document).ready(function() {

    $("#loading").hide();
    removeUnScroll();

    Highcharts.setOptions({
        global: {
            useUTC: false
        },
        credits: {
            enabled: false
        },
        navigation: {
            buttonOptions: {
                enabled: false
            }
        }
    });

    var gaugeOptions = {

        chart: {
            type: 'solidgauge'
        },

        title: null,

        pane: {
            center: ['50%', '85%'],
            size: '140%',
            startAngle: -90,
            endAngle: 90,
            background: {
                backgroundColor: (Highcharts.theme && Highcharts.theme.background2) || '#FFFFFF',
                innerRadius: '60%',
                outerRadius: '100%',
                shape: 'arc'
            }
        },

        tooltip: {
            enabled: false
        },

        // the value axis
        yAxis: {
            stops: [
                [0.1, '#02c39a'],
                [0.5, '#dddf00'],
                [0.9, '#db3a34']
            ],
            lineWidth: 0,
            minorTickInterval: null,
            tickAmount: 2,
            title: {
                y: -70
            },
            labels: {
                y: 16
            }
        },

        plotOptions: {
            solidgauge: {
                dataLabels: {
                    y: 5,
                    borderWidth: 0,
                    useHTML: true
                }
            }
        }
    };


    var chartCPU = Highcharts.chart('container-cpu', Highcharts.merge(gaugeOptions, {
        yAxis: {
            min: 0,
            max: 100,
            title: {
                text: ''
            }
        },

        series: [{
            name: 'CPU',
            data: [0],
            dataLabels: {
                format: '<div style="text-align:center"><span style="font-size:28px;color:' +
                    ((Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black') + '">{y}</span>' +
                    '<span style="font-size:12px;color:silver">%</span></div>'
            },
            tooltip: {
                valueSuffix: ' %'
            }
        }]

    }));

    var chartRAM = Highcharts.chart('container-mem', Highcharts.merge(gaugeOptions, {
        yAxis: {
            min: 0,
            max: parseFloat(init_vals.mem.total),
            title: {
                text: ''
            }
        },

        series: [{
            name: 'MEMORY',
            data: [0],
            dataLabels: {
                format: '<div style="text-align:center"><span style="font-size:25px;color:' +
                    ((Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black') + '">{y:.1f}</span><br/>' +
                    '<span style="font-size:12px;color:silver">MB</span></div>'
            },
            tooltip: {
                valueSuffix: ' MB'
            }
        }]

    }));

    var chartCache = Highcharts.chart('container-cache', Highcharts.merge(gaugeOptions, {
        yAxis: {
            min: 0,
            max: parseFloat(init_vals.mem.total),
            title: {
                text: ''
            }
        },

        series: [{
            name: 'CACHE',
            data: [0],
            dataLabels: {
                format: '<div style="text-align:center"><span style="font-size:12px;color:' +
                    ((Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black') + '">{y:.1f}</span><br/>' +
                    '<span style="font-size:10px;color:silver">MB</span></div>'
            },
            tooltip: {
                valueSuffix: ' MB'
            }
        }]

    }));

    var chartRAM_real = Highcharts.chart('container-mem-real', Highcharts.merge(gaugeOptions, {
        yAxis: {
            min: 0,
            max: parseFloat(init_vals.mem.total),
            title: {
                text: ''
            }
        },

        series: [{
            name: 'REAL MEMORY',
            data: [0],
            dataLabels: {
                format: '<div style="text-align:center"><span style="font-size:12px;color:' +
                    ((Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black') + '">{y:.1f}</span><br/>' +
                    '<span style="font-size:10px;color:silver">MB</span></div>'
            },
            tooltip: {
                valueSuffix: ' MB'
            }
        }]

    }));

    var chartSWAP = Highcharts.chart('container-swap', Highcharts.merge(gaugeOptions, {
        yAxis: {
            min: 0,
            max: parseFloat(init_vals.mem.swap.total),
            title: {
                text: ''
            }
        },

        series: [{
            name: 'SWAP',
            data: [0],
            dataLabels: {
                format: '<div style="text-align:center"><span style="font-size:12px;color:' +
                    ((Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black') + '">{y:.1f}</span><br/>' +
                    '<span style="font-size:10px;color:silver">MB</span></div>'
            },
            tooltip: {
                valueSuffix: ' MB'
            }
        }]

    }));

    var chartDisk = Highcharts.chart('container-disk', Highcharts.merge(gaugeOptions, {
        yAxis: {
            min: 0,
            max: parseFloat(init_vals.disk.total),
            title: {
                text: ''
            }
        },

        series: [{
            name: 'DISK',
            data: [0],
            dataLabels: {
                format: '<div style="text-align:center"><span style="font-size:12px;color:' +
                    ((Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black') + '">{y:.1f}</span><br/>' +
                    '<span style="font-size:10px;color:silver">GB</span></div>'
            },
            tooltip: {
                valueSuffix: ' GB'
            }
        }]

    }));


    var chartNetInterface1 = Highcharts.chart('container-net-interface-1', {
        title: {
            text: ''
        },
        legend: {
            enabled: false
        },
        xAxis: {
            categories: [],
            title: {
                text: ''
            }
        },
        yAxis: {
            title: {
                text: '',
                style: {
                    fontWeight: 'normal'
                }
            }
        },
        series: [
            {
                name: 'IN',
                data: [0],
                color: '#3a86ff',
                marker: {
                    enabled: false
                }
            },
            {
                name: 'OUT',
                data: [0],
                color: '#16db93',
                marker: {
                    enabled: false
                }
            }
        ]
    });
    net_In1 = [0,0,0,0,0,0,0,0,0,0];
    net_Out1 = [0,0,0,0,0,0,0,0,0,0];

    var chartNetInterface2 = Highcharts.chart('container-net-interface-2', {
        title: {
            text: ''
        },
        legend: {
            enabled: false
        },
        xAxis: {
            categories: [],
            title: {
                text: ''
            }
        },
        yAxis: {
            title: {
                text: '',
                style: {
                    fontWeight: 'normal'
                }
            }
        },
        series: [
            {
                name: 'IN',
                data: [0],
                color: '#3a86ff',
                marker: {
                    enabled: false
                }
            },
            {
                name: 'OUT',
                data: [0],
                color: '#16db93',
                marker: {
                    enabled: false
                }
            }
        ]
    });
    net_In2 = [0,0,0,0,0,0,0,0,0,0];
    net_Out2 = [0,0,0,0,0,0,0,0,0,0];

    setInterval(function() {
        $.ajaxSetup(csrfAddToAjaxHeader());
        $.post('api/device', function(data){
            $("#loading").hide();
            removeUnScroll();

            $("#login-users").text(data.login_user_count);
            $("#hostname").text(data.hostname);
            $("#uname").text(data.uname);
            $("#system").text(data.system);
            $("#time").text(data.now_time_hms);
            $("#date").text(data.now_time_ymd);
            $("#uptime").text(data.uptime);
            $("#cpu-temp").text(data.cpu_temperature);
            $("#cpu-freq").text(data.cpu_freq);
            $("#cpu-stat-idl").text(data.cpu_status_idle);
            $("#cpu-stat-use").text(data.cpu_status_user);
            $("#cpu-stat-sys").text(data.cpu_status_system);
            $("#cpu-stat-nic").text(data.cpu_status_nice);
            $("#cpu-stat-iow").text(data.cpu_status_iowait);
            $("#cpu-stat-irq").text(data.cpu_status_irq);
            $("#cpu-stat-sirq").text(data.cpu_status_softirq);
            $("#mem-percent").text(data.memory_percent);
            $("#mem-free").text(data.memory_free);
            $("#mem-cached").text(data.memory_cached);
            $("#mem-swap-total").text(data.swap_total);
            $("#mem-cache-percent").text(data.memory_cached_percent);
            $("#mem-buffers").text(data.memory_buffers);
            $("#mem-real-percent").text(data.memory_real_percent);
            $("#mem-real-free").text(data.memory_available);
            $("#mem-swap-percent").text(data.swap_used_percent);
            $("#mem-swap-free").text(data.swap_free);
            $("#disk-percent").text(data.disk_used_percent);
            $("#disk-free").text(data.disk_free);
            $("#loadavg-1m").text(data.load_average_1m);
            $("#loadavg-5m").text(data.load_average_5m);
            $("#loadavg-10m").text(data.load_average_15m);
            $("#loadavg-running").text(data.load_average_process_running);
            $("#loadavg-threads").text(data.load_average_process_total);

            $("#net-interface-1-total-in").text(data.net_status_lo_in_data_format);
            $("#net-interface-1-total-out").text(data.net_status_lo_out_data_format);
            $("#net-interface-2-total-in").text(data.net_status_in_data_format);
            $("#net-interface-2-total-out").text(data.net_status_out_data_format);

            if(window.dashboard != null)
            {
                window.dashboard_old = window.dashboard;
            }
            window.dashboard = data;

        }).fail(function() {
                $("#loading").show();
                unScroll();
            });

        if(window.dashboard != null){
            var point;
            if (chartCPU) {
                point = chartCPU.series[0].points[0];
                point.update(parseFloat(window.dashboard.cpu_used));
            }
            if (chartRAM) {
                point = chartRAM.series[0].points[0];
                point.update(parseFloat(window.dashboard.memory_used));
            }
            if (chartCache) {
                point = chartCache.series[0].points[0];
                point.update(parseFloat(window.dashboard.memory_cached));
            }
            if (chartRAM_real) {
                point = chartRAM_real.series[0].points[0];
                point.update(parseFloat(window.dashboard.memory_real_used));
            }
            if (chartSWAP) {
                point = chartSWAP.series[0].points[0];
                point.update(parseFloat(window.dashboard.swap_used));
            }
            if (chartDisk) {
                point = chartDisk.series[0].points[0];
                point.update(parseFloat(window.dashboard.disk_used));
            }

            if(window.dashboard_old != null) {
                if(chartNetInterface1.series[0].data.length >=30){
                    chartNetInterface1.series[0].addPoint(parseInt(window.dashboard.net_status_lo_in_data) - parseInt(window.dashboard_old.net_status_lo_in_data), true, true);
                }
                else{
                    chartNetInterface1.series[0].addPoint(parseInt(window.dashboard.net_status_lo_in_data) - parseInt(window.dashboard_old.net_status_lo_in_data));
                }

                if(chartNetInterface1.series[1].data.length >=30){
                    chartNetInterface1.series[1].addPoint(parseInt(window.dashboard.net_status_lo_out_data) - parseInt(window.dashboard_old.net_status_lo_out_data), true, true);
                }
                else{
                    chartNetInterface1.series[1].addPoint(parseInt(window.dashboard.net_status_lo_out_data) - parseInt(window.dashboard_old.net_status_lo_out_data));
                }
            }

            if(window.dashboard_old != null) {
                if(chartNetInterface2.series[0].data.length >=30){
                    chartNetInterface2.series[0].addPoint(parseInt(window.dashboard.net_status_in_data) - parseInt(window.dashboard_old.net_status_in_data), true, true);
                }
                else{
                    chartNetInterface2.series[0].addPoint(parseInt(window.dashboard.net_status_in_data) - parseInt(window.dashboard_old.net_status_in_data));
                }

                if(chartNetInterface2.series[1].data.length >=30){
                    chartNetInterface2.series[1].addPoint(parseInt(window.dashboard.net_status_out_data) - parseInt(window.dashboard_old.net_status_out_data), true, true);
                }
                else{
                    chartNetInterface2.series[1].addPoint(parseInt(window.dashboard.net_status_out_data) - parseInt(window.dashboard_old.net_status_out_data));
                }
            }

        }

    }, (parseInt($("#interval").text()) * 1000) );
}
)

function unScroll() {
    var top = $(document).scrollTop();
    $(document).on('scroll.unable',function (e) {
        $(document).scrollTop(top);
    })
    $(document.body).css({
        "overflow-y": "hidden"
    });
}

function removeUnScroll() {
    $(document).unbind("scroll.unable");
    $(document.body).css({
        "overflow-y": "auto"
    });
}


$("#logout").click(function(){
    $("#logout").attr("disabled", true);
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/logout', function(result){
        if (result.status == true) {
            $("#logout").attr("disabled", false);
            $(window).attr('location','/login');
        } else {
            $("#logout").attr("disabled", false);
            window.alert("Sign out failed");
        }
    }).fail(function() {
        $("#login").attr("disabled", false);
        window.alert("Sign out failed");
    });
});

$("#reboot").click(function(){
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=reboot', function(data){
        if (data.status == true) {
            window.alert("OK")
            $("#loading").show();
            unScroll();
        }
    }).fail(function() {
        window.alert("Fail");
        $("#loading").show();
        unScroll();
    });
});

$("#shutdown").click(function(){
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=shutdown', function(data){
        if (data.status == true) {
            window.alert("OK");
            $("#loading").show();
            unScroll();
        }
    }).fail(function() {
        window.alert("Fail");
        $("#loading").show();
        unScroll();
    });
});

$("#dropcaches").click(function(){
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=dropcaches', function(data){ //$.getJSON()
        if (data.status == true) {
            window.alert("OK");
            // $("#loading").show();
            // unScroll();
        }
    }).fail(function() {
        window.alert("Fail");
        // $("#loading").show();
        // unScroll();
    });
});

// Check New Version
$(document).ready(function(){
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=checknewversion', function(data){
        if (data.new_version != "" && data.new_version_url != "") {
            $("#new-url").attr("href", data.new_version_url);
            $("#new-version").attr("title", "v" + data.new_version);
            $("#new-box").show(1000);
        } else {
            $("#new-box").hide(1000);
            $("#new-url").attr("href", "javascript:void(0);");
            $("#new-version").attr("title", "New Version");
        }
    }).fail(function() {
        $("#new-box").hide(1000);
        $("#new-url").attr("href", "javascript:void(0);");
        $("#new-version").attr("title", "New Version");
    });
});



// Login Page
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
        $.ajaxSetup(csrfAddToAjaxHeader());
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