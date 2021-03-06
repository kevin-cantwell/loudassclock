package clock

const Html = `
<html>
<head>
  <style>
    body { background-color: black; }
    .clock { margin: 0 auto; width: 1220px; position: relative; }
    .digit { background-image: url(/images/digits.png); position: absolute; }
    .hour-ten { top: 0; left: 0; }
    .hour-one { top: 0; left: 273px; }
    .dots { top: 125px; left: 546px; }
    .min-ten { top: 0; left: 624px; }
    .min-one { top: 0; left: 897px; }
    .hidden { display: none; }
    .dnil { background: none;}
    .d1{ background-position: 0 0; width: 283px; height: 499px; }
    .d2{ background-position: -333px 0; width: 283px; height: 499px; }
    .d3{ background-position: -666px 0; width: 283px; height: 499px; }
    .d4{ background-position: -999px 0; width: 283px; height: 499px; }
    .d5{ background-position: -1332px 0; width: 283px; height: 499px; }
    .d6{ background-position: -1665px 0; width: 283px; height: 499px; }
    .d7{ background-position: 0 -549px; width: 283px; height: 499px; }
    .d8{ background-position: -333px -549px; width: 283px; height: 499px; }
    .d9{ background-position: -666px -549px; width: 283px; height: 499px; }
    .d0{ background-position: -1137px -549px; width: 283px; height: 499px; }
    .dots{ background-position: -999px -549px; width: 88px; height: 250px; }
  </style>
</head>
<body>
  <div class="clock">
    <div class="digit hour-ten d0"></div>
    <div class="digit hour-one d0"></div>
    <div class="digit dots"></div>
    <div class="digit min-ten d0"></div>
    <div class="digit min-one d0"></div>
  </div>
  <script type="text/javascript">
    var displayDots = false;
    update();
    setInterval(update, 1000);
    
    function update() {
      var date = new Date();
      var hours = date.getHours();
      var hoursStr = hours.toString();
      var hourTen = "0";
      var hourOne = hoursStr;
      if (hoursStr.length > 1) {
        hourTen = hoursStr[0];
        hourOne = hoursStr[1];
      }
      var minutesStr = date.getMinutes().toString();
      var minTen = "0";
      var minOne = minutesStr;
      if (minutesStr.length > 1) {
        minTen = minutesStr[0];
        minOne = minutesStr[1];
      }
      document.getElementsByClassName('hour-ten')[0].setAttribute('class', 'digit hour-ten d' + hourTen);
      document.getElementsByClassName('hour-one')[0].setAttribute('class', 'digit hour-one d' + hourOne);
      document.getElementsByClassName('min-ten')[0].setAttribute('class', 'digit min-ten d' + minTen);
      document.getElementsByClassName('min-one')[0].setAttribute('class', 'digit min-one d' + minOne);

      displayDots = !displayDots
      var display = ' hidden';
      if (displayDots) {
        display = '';
      }
      document.getElementsByClassName('dots')[0].setAttribute('class', 'digit dots' + display);
    }
  </script>
</body>
</html>
`
