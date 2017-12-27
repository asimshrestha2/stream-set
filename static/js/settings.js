function GETRequest(opts = {
    url: "",
    success: function(){},
}){
    var xhr = new XMLHttpRequest();
    xhr.open('GET', opts.url, true);
    xhr.onload = function(){
        if(opts.success) opts.success(xhr.responseText);
    }
    xhr.send();
}

function POSTRequest(opts = {
    url: "",
    data: "",
    success: function(){},
}){
    var xhr = new XMLHttpRequest();
    xhr.open('POST', opts.url, true);
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
    xhr.onload = function(){
        if(opts.success) opts.success(xhr.responseText);
    }
    xhr.send(opts.data);
}

var filepath_input = document.getElementById('filepath'),
    color_input = document.getElementById('color-settings'),
    time_input = document.getElementById('time-settings'),
    logopath_input = document.getElementById('logopath');
var setfilepath_btn = document.getElementById('set-filepath'),
    savecustomization_btn = document.getElementById('save-customization');

GETRequest({
    url: "/api/getfilepath",
    success: function(resp){
        var jd = JSON.parse(resp);
        if(jd){
            filepath_input.value = jd.filelocation;
            color_input.value = jd.circlecolor;
            time_input.value = jd.timersetting;
            logopath_input.value = jd.logopath;
        }
    }
})

setfilepath_btn.addEventListener('click', function(){
    POSTRequest({
        url: "/api/setsettings",
        data: "filepath=" + filepath_input.value,
    })
})

savecustomization_btn.addEventListener('click', function(){
    POSTRequest({
        url: "/api/setsettings",
        data: "circlecolor=" + color_input.value + "&timersetting=" + time_input.value + "&logopath=" + logopath_input.value,
    })
})