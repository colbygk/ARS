
$(function() {
  $('#depicker').datetimepicker({format: 'YYYY-MM-DD'});
  $('#repicker').datetimepicker({format: 'YYYY-MM-DD'});

/*
  $("#depicker").on("dp.change",function (e) {
    $('#repicker').data("DateTimePicker").minDate(e.date);
  });
  $("#repicker").on("dp.change",function (e) {
    $('#depicker').data("DateTimePicker").maxDate(e.date); });
  $('#depicker').data("DateTimePicker").minDate(moment());
  $('#repicker').data("DateTimePicker").minDate(moment());
  */
  });
