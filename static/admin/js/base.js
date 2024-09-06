$(function () {
	baseApp.init();
})

var baseApp = {
	// 加载页面时，执行 init
	init: function () {
		this.initAside()
		this.resizeIframe()
		this.confirmDelete()
		this.changeStatus()
		this.changeNum()
	},

	// 定义 initAside 的函数，给页面上类名为 aside 的元素内所有 <h4> 标签绑定一个点击事件。
	// 当 <h4> 标签被点击时，各自的同级 <ul> 元素会执行一个滑动切换（slideToggle）效果。
	initAside: function () {
		$('.aside h4').click(function () {
			$(this).siblings('ul').slideToggle();
		})
	},

	// 计算、设置 iframe 高度
	resizeIframe: function () {
		$("#rightMain").height($(window).height() - 80)
	},

	// 删除提示
	confirmDelete: function () {
		$(".delete").click(function () {
			var flag = confirm("确定删除吗？")
			return flag
		})
	},

	// ajax 异步修改图片状态
	changeStatus: function () {
		$(".chStatus").click(function () {
			var id = $(this).attr("data-id")
			var table = $(this).attr("data-table")
			var field = $(this).attr("data-field")
			var el = $(this)
			$.get("/admin/changeStatus", { id: id, table: table, field: field }, function (response) {
				if (response.success) {
					if (el.attr("src").indexOf("yes") != -1) {
						el.attr("src", "/static/admin/images/no.gif")
					} else {
						el.attr("src", "/static/admin/images/yes.gif")
					}
				}
			})
		})
	},

	// ajax 异步修改排序
	changeNum: function () {
		/*
		1、获取el里面的值  var spanNum=$(this).html()
		2、创建一个input的dom节点   var input=$("<input value='' />");
		3、把input放在el里面   $(this).html(input);
		4、让input获取焦点  给input赋值    $(input).trigger('focus').val(val);
		5、点击input的时候阻止冒泡 
					$(input).click(function(e){
						e.stopPropagation();				
					})					
		6、鼠标离开的时候给span赋值,并触发ajax请求
			$(input).blur(function(){
				var inputNum=$(this).val();
				spanEl.html(inputNum);
				触发ajax请求
			})
		*/
		$(".chSpanNum").click(function () {
			// 1. 获取el 以及el里面的属性值
			var id = $(this).attr("data-id")
			var table = $(this).attr("data-table")
			var field = $(this).attr("data-field")
			var num = $(this).html().trim()
			var spanEl = $(this)

			// 2. 创建一个input的dom节点   var input=$("<input value='' />");
			var input = $("<input style='width:60px'  value='' />");

			// 3. 把input放在el里面   $(this).html(input);
			$(this).html(input);

			// 4. 让input获取焦点  给input赋值    $(input).trigger('focus').val(val);
			$(input).trigger("focus").val(num);

			// 5. 点击input的时候阻止冒泡 
			$(input).click(function (e) {
				e.stopPropagation();
			})

			// 6.鼠标离开的时候给span赋值,并触发ajax请求
			$(input).blur(function () {
				var inputNum = $(this).val()
				spanEl.html(inputNum)
				//触发ajax请求
				$.get("/admin/changeNum", { id: id, table: table, field: field, num: inputNum }, function (response) {
					console.log(response)
				})
			})
		})
	},
}
