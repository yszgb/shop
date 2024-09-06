// 通过接口获取验证码图片，供 login.html 传入到前端页面
$(function () {
    loginApp.init();
})

var loginApp = {
    // 加载页面时，执行 init
    init: function () {
        this.getCaptcha()
        this.captchaImgChange()
    },
    
    getCaptcha: function () {
        // 加随机数解决缓存问题
        $.get("/admin/captcha?t=" + Math.random(), function (response) {
            console.log(response)
            $("#captchaId").val(response.captchaId)
            $("#captchaImg").attr("src", response.captchaImg)
        })
    },

    // 点一下刷新验证码图片
    captchaImgChange: function () {
        var that = this;
        // $("#captchaImg") 点击 captchaImg 时触发：
        $("#captchaImg").click(function () {
            that.getCaptcha()
        })
    }
}