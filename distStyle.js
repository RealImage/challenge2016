$(document).ready(function(){
	var city= [], province = [],flag = false;
	$(".box").click(function(){
		$(".dashboard,form .search,form .createDist,form .createSub").addClass('hide');
		$(".forms").removeClass('hide');
		if($(this).hasClass('search')){
			$('form .search').removeClass('hide');
		}
		else if($(this).hasClass('create_sub_dist')){
			$('form .createSub').removeClass('hide');
		}
		else{
			$('form .createDist').removeClass('hide');
		}
	})
	for(var i=0;i<Object.keys(cities).length;i++){
		var state = Object.keys(cities)[i];
		$.each(cities[state],function(i,Obj){
			city.push(Obj);
		})
		province.push(state);
	}
	$("form input[name=province]").focusout(function(){
		var val = $(this).val();
		if(val !== ""){
			autoComp("form input[name=city]", cities[val]);
		}
	})
	$("form input[name=city]").focusout(function(){
		var val = $(this).val();
		if(val !== ""){
			for(var j=0;j<Object.keys(cities).length;j++){
				var state = Object.keys(cities)[j];
				$.each(cities[state],function(i,Obj){
					if(Obj === val){
						j = Object.keys(cities).length;
						$("form input[name=province]").val(state);
						$("form input[name=country]").val("India");
						return false;
					}
				})
			}
		}
	})
	autoComp("form input[name=distributor]", Object.keys(dist));
	autoComp("form input[name=province]", province);
	autoComp("form input[name=city]", city);
	autoComp("form input[name=country]", ["India"]);
})
function back(){
	$('.dashboard').removeClass('hide');
	$('.forms').addClass('hide');
}
function home(){
	$('.dashboard').removeClass('hide');
	$('.result').addClass('hide');
}
function autoComp(element,src){
	$(element).autocomplete({
		source: src,
		minLength: 3
	})
}
function tablePopulate(dist){
	$(".forms").addClass('hide');
	$(".result").removeClass('hide');
	$("table tbody").empty();
	$.each(Object.keys(dist),function(i,Obj){
		$("table tbody").append("<tr><td>"+Obj+"</td><td>"+dist[Obj].include.City+"</td><td>"+dist[Obj].include.Province+"</td><td>"+dist[Obj].exclude.City+"</td><td>"+dist[Obj].exclude.Province+"</td></tr>")
	})
}
function search(){

}
function validate(e){
	e.preventDefault();
	if($('input[name=distributor]').val() !== ""){
		if(!$('.search').hasClass('hide')){
			var searchTemp = {};
			$.each(Object.keys(dist),function(i,Obj){
				if($('input[name=distributor]').val() === Obj){
					searchTemp[Obj] = dist[Obj];
					tablePopulate(searchTemp);
					return false;
				}
				else if(Object.keys(dist).length === i+1){
					alert("No matches found");
				}
			})
		}
		else{
			var City = $('input[name=city]').val(),Province=$('input[name=province]').val(),
			include={
				"City": City,
				"Province": Province,
				"Country": "India"
			},
			exclude={
				"City": "",
				"Province": "",
				"Country": ""
			}
			dist[$('input[name=distributor]').val()] = {
				"include": include,
				"exclude": exclude
			}
			tablePopulate(dist);
		}
	}
	else{
		alert("Please enter the Distributor's name");
	}
}