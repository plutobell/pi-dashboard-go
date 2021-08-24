// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-24
// @Version: 1.4.0

var new_version = ""
var new_version_notes = ""
var new_version_url = ""

initTooltips();
unScroll();

$(document).ready(function() {
    var net_in_color = $(":root").css("--net-in-color");
    var net_out_color = $(":root").css("--net-out-color");

    $("#year").text(new Date().getFullYear());
    $("#loading").hide();
    $('#ModalBox').modal('hide');
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
                [0.1, '#D9E4DD'],
                [0.5, '#CDC9C3'],
                [0.9, '#919191']
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
                color: net_in_color,
                marker: {
                    enabled: false
                }
            },
            {
                name: 'OUT',
                data: [0],
                color: net_out_color,
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
                color: net_in_color,
                marker: {
                    enabled: false
                }
            },
            {
                name: 'OUT',
                data: [0],
                color: net_out_color,
                marker: {
                    enabled: false
                }
            }
        ]
    });
    net_In2 = [0,0,0,0,0,0,0,0,0,0];
    net_Out2 = [0,0,0,0,0,0,0,0,0,0];

    setInterval(function() {
        var date = new Date();
        var year = date.getFullYear();
        var month = date.getMonth();
        var day = date.getDate();

        $.ajaxSetup(csrfAddToAjaxHeader());
        $.post('api/device', function(data){
            $("#loading").hide();
            removeUnScroll();

            $("#login-users").text(data.login_user_count);
            $("#hostip").text(data.ip);
            $("#hostname").text(data.hostname);
            $("#uname").text(data.uname);
            $("#system").text(data.system);
            // $("#time").text(data.now_time_hms);
            // $("#date").text(data.now_time_ymd);
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

            $("#version").text("v" + data.version);
            $("#version").attr('data-bs-original-title', "Compiled with " + data.go_version);

            // $("#year").text(new Date().getFullYear());

            if(window.dashboard != null)
            {
                window.dashboard_old = window.dashboard;
            }
            window.dashboard = data;

        }).fail(function() {
                $("#loading").show();
                $('#ModalBox').modal('hide');
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
});

// Loading Consumption Time
$(window).on('load', function() {
    var endTime = new Date().getTime();
    loading_time = endTime - startTime
    if (loading_time < 1000) {
        $("#loading-time").attr('data-bs-original-title', "Loading time: " + loading_time + "ms");
    } else {
        $("#loading-time").attr('data-bs-original-title', "Loading time: " + (loading_time / 1000).toFixed(3) + "s");
    };
});

$(document).ready(function() {
    setInterval(function() {
        var nowDate = new Date();
        var year = nowDate.getFullYear();
        var month = nowDate.getMonth() + 1 < 10 ? "0" + (nowDate.getMonth() + 1) : nowDate.getMonth() + 1;
        var day = nowDate.getDate() < 10 ? "0" + nowDate.getDate() : nowDate.getDate();
        var hour = nowDate.getHours()< 10 ? "0" + nowDate.getHours() : nowDate.getHours();
        var minute = nowDate.getMinutes()< 10 ? "0" + nowDate.getMinutes() : nowDate.getMinutes();
        var second = nowDate.getSeconds()< 10 ? "0" + nowDate.getSeconds() : nowDate.getSeconds();

        var hms = hour + ":" + minute + ":" + second
        var ymd = year + "-" + month + "-" + day

        $("#time").text(hms);
        $("#date").text(ymd);
        $("#year").text(year);
    }, (500) );
});

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
    $("#logout").css("pointer-events", "none");
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/logout', function(result){
        if (result.status == true) {
            $("#logout").css("pointer-events", "auto");
            $(window).attr('location','/login');
        } else {
            $("#logout").css("pointer-events", "auto");
            showModalBox("Sign Out", "Sign out failed");
        }
    }).fail(function() {
        $("#logout").css("pointer-events", "auto");
        showModalBox("Sign Out", "Sign out failed");
    });
});

$("#reboot").click(function(){
    $("#reboot").css("pointer-events", "none");
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=reboot', function(data){
        if (data.status == true) {
            showModalBox("Reboot", "OK")
            $("#loading").show();
            unScroll();
            $("#reboot").css("pointer-events", "auto");
        } else {
            showModalBox("Reboot", "Fail");
            $("#loading").show();
            unScroll();
            $("#reboot").css("pointer-events", "auto");
        }
    }).fail(function() {
        showModalBox("Reboot", "Fail");
        $("#loading").show();
        unScroll();
        $("#reboot").css("pointer-events", "auto");
    });
});

$("#shutdown").click(function(){
    $("#shutdown").css("pointer-events", "none");
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=shutdown', function(data){
        if (data.status == true) {
            showModalBox("Shutdown", "OK");
            $("#loading").show();
            unScroll();
            $("#shutdown").css("pointer-events", "auto");
        } else {
            showModalBox("Shutdown", "Fail");
            $("#loading").show();
            unScroll();
            $("#shutdown").css("pointer-events", "auto");
        }
    }).fail(function() {
        showModalBox("Shutdown", "Fail");
        $("#loading").show();
        unScroll();
        $("#shutdown").css("pointer-events", "auto");
    });
});

$("#dropcaches").click(function(){
    $("#dropcaches").css("pointer-events", "none");
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=dropcaches', function(data){ //$.getJSON()
        if (data.status == true) {
            showModalBox("Drop Caches", "OK");
            // $("#loading").show();
            // unScroll();
            $("#dropcaches").css("pointer-events", "auto");
        } else {
            showModalBox("Drop Caches", "Fail");
            // $("#loading").show();
            // unScroll();
            $("#dropcaches").css("pointer-events", "auto");
        }
    }).fail(function() {
        showModalBox("Drop Caches", "Fail");
        // $("#loading").show();
        // unScroll();
        $("#dropcaches").css("pointer-events", "auto");
    });
});

// Check New Version
$(document).ready(function(){
    $.ajaxSetup(csrfAddToAjaxHeader());
    $.post('/api/operation?action=checknewversion', function(data){
        if (data.new_version != "" && data.new_version_url != "") {
            // $("#new-url").attr("href", data.new_version_url);
            new_version = data.new_version
            new_version_notes = data.new_version_notes
            new_version_url = data.new_version_url

            $('#new-box').attr('data-bs-original-title', "v" + data.new_version + " Available");
            $("#new-box").show(1000);
        } else {
            new_version = ""
            new_version_notes = ""
            new_version_url = ""

            $("#new-box").hide(1000);
            // $("#new-url").attr("href", "javascript:void(0);");
            $('#new-box').attr('data-bs-original-title', 'New Version');
        }
    }).fail(function() {
        new_version = ""
        new_version_notes = ""
        new_version_url = ""

        $("#new-box").hide(1000);
        // $("#new-url").attr("href", "javascript:void(0);");
        $('#new-box').attr('data-bs-original-title', 'New Version');
    });
});
$("#new-box").click(function(){
    $("#new-box").css("pointer-events", "none");
    $("#new-download").show();

    $('#ModalBox-Title').text("v" + new_version + " Available");
    $("#ModalBox-Body").empty();
    $("#ModalBox-Body").attr('style', 'text-align: left !important;');
    var notes_array = new_version_notes.split("*");
    $("#ModalBox-Body").append("<strong>Release Notes</strong>");
    $("#ModalBox-Body").append("<ul></ul>");
    for (var i = 0; i < notes_array.length; i++) {
        if (notes_array[i] != "") {
            insertHtml = "<li><small>" + $.trim(notes_array[i]) + "</small></li>"
            $('#ModalBox-Body').find('ul').eq(0).append(insertHtml);
        }
    }

    $('#ModalBox').modal('show');
    $("#new-box").css("pointer-events", "auto");
});
$("#new-download").click(function(){
    window.open(new_version_url, "_blank");
});

function showModalBox(title, body) {
    $("#new-download").hide();
    $('#ModalBox-Title').text(title);
    $("#ModalBox-Body").empty();
    $("#ModalBox-Body").attr('style', 'text-align: center !important;');
    $('#ModalBox-Body').text(body);
    $('#ModalBox').modal('show');
}

function initTooltips() {
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl)
    });
}

