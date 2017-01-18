$(function () {
	$(window).load(function() {
		tags = [];
		$("input.tag").each(function() {
			console.log($(this).val());
			if($.inArray($(this).val(), tags) === -1) {
				tags.push($(this).val());
			}
		});
	});
});