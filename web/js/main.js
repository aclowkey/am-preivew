$(document).ready(function () {
  $('#render').click(function () {
    let template = $('#template').val()
    let data = $('#data').val()

    $.post( "render", { template: template, data: data })
      .done(function( data ) {
        $('#result').val(""+data)
      });
  })
})
