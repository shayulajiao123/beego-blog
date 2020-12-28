$(function(){

	$(window).resize(function(){
		change();
	});
	change();
	function change(){
		var win_wid=$(window).width();
		if(win_wid>=1200){
			
		}else{
			
		}
	}
	// 导航滚动
	$(window).scroll(function(){
		var top = 100;//根据id获取指定的位置
		var scrollTop = $(window).scrollTop();//获取当前滑动的位置
		console.log('scrollTop----------',scrollTop)
		if(scrollTop > top){
			$("header.normal").addClass("active");
		}else{
			$("header.normal").removeClass("active");
		}
	})
	// footer点击回到顶部
	$(".goBack").click(function(){   
		$('body,html').animate({ 
			scrollTop:0 
		},700);
		return false; //防止冒泡
	});
	$(".menu-btn.menu-open-btn").click(function(){
		$(".mask").show()
		$(".menu-open-btn").hide()
		$(".menu-close-btn").show()
		$('html,body').css({'overflow':'hidden'})
	})
	$(".menu-btn.menu-close-btn").click(function(){
		$(".mask").hide()
		$(".menu-open-btn").show()
		$(".menu-close-btn").hide()
		$('html,body').css({'overflow':'auto'})
	})
});
