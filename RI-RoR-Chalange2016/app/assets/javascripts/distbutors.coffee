# Place all the behaviors and hooks related to the matching controller here.
# All this logic will automatically be available in application.js.
# You can use CoffeeScript in this file: http://coffeescript.org/
ready = ()->
  $('#distbutor_included_states, #distbutor_excluded_states').select2
    ajax:
      url: "/states/search",
      minimumInputLength: 3
      dataType: 'json'
      delay: 250
      data: (params) ->
        return {state: params.term, countries: $('#distbutor_included_countries').val()}
      processResults: (data)->
        processed_data = []
        for datum in data
          processed_data.push
            id: datum.id
            text: datum.name
        return {results: processed_data}
  $('#permision_states').select2
    ajax:
      url: "/states/search",
      minimumInputLength: 3
      dataType: 'json'
      delay: 250
      data: (params) ->
        return {state: params.term, countries: $('#permision_countries').val()}
      processResults: (data)->
        processed_data = []
        for datum in data
          processed_data.push
            id: datum.id
            text: datum.name
        return {results: processed_data}
  $('#distbutor_included_cities').select2
    ajax:
      url: "/cities/search",
      minimumInputLength: 3
      dataType: 'json'
      delay: 250
      data: (params) ->
        return {city: params.term, states: $('#distbutor_included_states').val()}
      processResults: (data)->
        processed_data = []
        for datum in data
          processed_data.push
            id: datum.id
            text: datum.name
        return {results: processed_data}
  $('#distbutor_excluded_cities').select2
    ajax:
      url: "/cities/search",
      minimumInputLength: 3
      dataType: 'json'
      delay: 250
      data: (params) ->
        return {city: params.term, states: $('#distbutor_excluded_states').val().concat($('#distbutor_included_states').val())}
      processResults: (data)->
        processed_data = []
        for datum in data
          processed_data.push
            id: datum.id
            text: datum.name
        return {results: processed_data}
  $('#permision_cities').select2
    ajax:
      url: "/cities/search",
      minimumInputLength: 3
      dataType: 'json'
      delay: 250
      data: (params) ->
        return {city: params.term, states: $('#permision_states').val().concat($('#distbutor_included_states').val())}
      processResults: (data)->
        processed_data = []
        for datum in data
          processed_data.push
            id: datum.id
            text: datum.name
        return {results: processed_data}
  $('#ds_permision_check').click ()->
    $.ajax
      method: "get"
      context: "json"
      url: "/distbutors/"+$('#permision_distbutor').val()+"/permision"
      data:
        permision:
          countries: $('#permision_countries').val()
          states: $('#permision_states').val()
          cities: $('#permision_cities').val()
      success: (data)->
        $('div.callout').remove()
        $('.box-header').after($('<div>').attr("class","callout callout-"+data.status).html($("<h4>").text(data.status)).append($("<p>").text(data.message)))

document.addEventListener 'turbolinks:load', ready
