

window.onload = checkSpeed();

estimates = [];
proceed = true;

passNum = 0;

function checkSpeed() {
  document.body.innerHTML = "calculating...<br>";
    // Spawn 15 web workers for calculating the speed.
    for(i = 0; i < 1; i++) {
      var worker = new Worker('calcClockSpeedWorker.js');
      // Any time those workers get back to us...
      worker.onmessage = function(e) {
        // Check their result to make sure it's not negative
        // or above 6GHz (because the fastest CPU on the market right now is 5GHz)
        if(e.data >= 0 && e.data <= 10) {
          // If it's valid, push their result to an array.
          estimates.push(e.data);
          passNum++;
          // for each tenth pass, display such on the left
          if(passNum % 10 == 0) {
            document.body.innerHTML += "<span class='sidemessage'>pass "+passNum+"...</span><br>"
          }
        } else {
          // otherwise, ask it to run again.
          worker.postMessage([0]);
        }

        // Once that array reaches 15...
        if(estimates.length == 200) {
          // get the average of it.
          total = 0;
          for(i = 0; i < estimates.length; i++) {
            total += estimates[i];
          }
          total = total/estimates.length;
          total = Math.ceil(total*10)/10
          document.body.innerHTML = "<center><h1>"+total+"GHz</h1><small>(very roughly)</small></center>";
        }
      }
    }
}