function init(data) {
    // console.log(data)
    // data = JSON.parse(data)
    // if (data.CODE==200){
    // }
}

$("#upload").click(function(){
    $.ajaxFileUpload({
        url:"uploadFile",
        secureuri: false,
        data: {
            'groupId': $("#groupId").val(),
            'artifactId': $("#artifactId").val(),
            'version': $("#version").val(),
            'fileExt': $("#fileExt").val(),
        },
        fileElementId:"file",
        dataType: 'json',
        success:function (data, status) {
            if (data&&status=="success") {
                if (data.CODE==200){
                    if(confirm("上传文件成功，是否跳转到列表页？")){
                        window.location.href = "/";
                    }else {

                    }
                } else {
                    alert(data.MSG);
                }
            }
        },
        error: function (data, status, e)
        {
            alert(e);
        }
    })
})