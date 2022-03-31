jQuery(document).ready(function ($) {

	/************** Menu Content Opening *********************/
	$(".main_menu a, .responsive_menu a").click(function () {
		var id = $(this).attr('class');
		id = id.split('-');
		$("#menu-container .content").hide();
		$("#menu-container #menu-" + id[1]).addClass("animated fadeInDown").show();
		$("#menu-container .homepage").hide();
		$("#menu-container .mapconsert").hide();
		$("#menu-container .search").hide();
		$("#menu-container .randompage").hide();
		$(".support").hide();
		$(".testimonials").hide();
		return false;
	});


	$(".main_menu a.HomePage").addClass('active');
	// $(".main_menu a.MapConsert").removeClass('active');

	$(".main_menu a.HomePage, .responsive_menu a.HomePage").click(function () {
		$("#menu-container .homepage").addClass("animated fadeInDown").show();
		$(this).addClass('active');
		$(".main_menu a.MapConcert, .responsive_menu a.MapConcert").removeClass('active');
		$(".main_menu a.Search, .responsive_menu a.Search").removeClass('active');
		$(".main_menu a.RandomPage, .responsive_menu a.RandomPage").removeClass('active');
		return false;
	});

	$(".main_menu a.MapConcert, .responsive_menu a.MapConcert").click(function () {
		$("#menu-container .mapconsert").addClass("animated fadeInDown").show();
		$(this).addClass('active');
		$(".main_menu a.HomePage, .responsive_menu a.HomePage").removeClass('active');
		$(".main_menu a.Search, .responsive_menu a.Search").removeClass('active');
		$(".main_menu a.RandomPage, .responsive_menu a.RandomPage").removeClass('active');
		return false;
	});

	$(".main_menu a.Search, .responsive_menu a.Search").click(function () {
		$("#menu-container .search").addClass("animated fadeInDown").show();
		$(this).addClass('active');
		$(".main_menu a.MapConcert, .responsive_menu a.MapConcert").removeClass('active');
		$(".main_menu a.HomePage, .responsive_menu a.HomePage").removeClass('active');
		$(".main_menu a.RandomPage, .responsive_menu a.RandomPage").removeClass('active');
		return false;
	});

	$(".main_menu a.RandomPage, .responsive_menu a.RandomPage").click(function () {
		$("#menu-container .randompage").addClass("animated fadeInDown").show();
		$(this).addClass('active');
		$(".main_menu a.MapConcert, .responsive_menu a.MapConcert").removeClass('active');
		$(".main_menu a.Search, .responsive_menu a.Search").removeClass('active');
		$(".main_menu a.HomePage, .responsive_menu a.HomePage").removeClass('active');
		return false;
	});

	/************** Gallery Hover Effect *********************/
	$(".overlay").hide();

	$('.gallery-item').hover(
		function () {
			$(this).find('.overlay').addClass('animated fadeIn').show();			
		},
		function () {
			$(this).find('.overlay').removeClass('animated fadeIn').hide();
		}
	);



});

function randomArtist() {
	var random = Math.floor(Math.random() * 52) + 1;
	return String(random)
}

function searchlist() {
	var input, filter, ul, li, a, i, txtValue;
	input = document.getElementById("searchbar");
	filter = input.value.toUpperCase();
	ul = document.getElementById("searchlist");
	li = ul.getElementsByTagName("li");

	if (filter.length == 0) {
		ul.style.display = "none";
	} else {
		ul.style.display = "";
	}

	for (i = 0; i < li.length; i++) {
		a = li[i].getElementsByTagName("a")[0];
		txtValue = a.textContent || a.innerText;
		if (txtValue.toUpperCase().indexOf(filter) > -1) {
			li[i].style.display = "";
		} else {
			li[i].style.display = "none";
		}
	}
}
