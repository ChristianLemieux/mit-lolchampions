$(function () {
	$(window).load(function() {
		tags = ["All"];
		$("input.tag").each(function() {
			if($.inArray($(this).val(), tags) === -1) {
				tags.push($(this).val());
			}
		});
		container = $("#filter-container");
		$.each(tags, function(idx, value){
			var btn = $('<button/>', {
				text: value,
				class: "btn",
				click: function () { filterChampions(value); }
			})
			container.append(btn);
		});

		function filterChampions(val){
			if(val == "All"){
				$("div.champion").show();
			}else{
				$("div.champion").hide();
				$("div.champion."+val).show();
			}
		}
	});
});