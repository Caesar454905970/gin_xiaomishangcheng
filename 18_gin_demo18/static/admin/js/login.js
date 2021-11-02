$(function(){
    app.init();
})
var app={
    init:function(){
        this.getCaptcha()
        this.captchaImgChage()
    },
    getCaptcha:function(){
        $.get("/admin/captcha?t="+Math.random(),function(response){
            console.log(response)
            $("#captchaId").val(response.CaptchaId)
            $("#captchaImg").attr("src",response.CaptchaImage)
        })
    },
    captchaImgChage:function(){
        var that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}