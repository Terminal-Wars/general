//
// JSBenchmark by Aaron Becker on StackOverflow, 
// modified to be much more accurate and less error prone
// 

if (self.Worker) {
	function calc() {
	  // DISCLAIMER: I DO NOT UNDERSTAND THE THOUGHT PROCESS BEHIND
	  // ANY OF THE SHIT BELOW. ALL I CHANGED WAS THE amount, and estprocesser,
	  // and i made it make sure it's always in a reasonable range.
	  //if speed=(c*a)/t, then constant=(s*t)/a and time=(a*c)/s
	  var _speedconstant = 1.15600e-8; 
	  var d = new Date();
	  var amount = 500000000;
	  // average processor speed according to steam hardware survey (in GHZ)
	  var estprocessor = 2.3; 
	  for (var i = amount; i > 0; i--) {}
	  var newd = new Date();
	  var accnewd = Number(String(newd.getSeconds()) + "." + String(newd.getMilliseconds()));
	  var accd = Number(String(d.getSeconds()) + "." + String(d.getMilliseconds()));
	  var di = accnewd - accd;
	  //console.log(accnewd,accd,di);
	  if (d.getMinutes() != newd.getMinutes()) {
	    di = (60 * (newd.getMinutes() - d.getMinutes())) + di
	  }
	  spd = ((_speedconstant * amount) / di);
	  console.log(spd);
	  final = Math.round(spd * 1000) / 1000;
	  // sometimes javascript gets a lil' quirky and needs to be divided by 10 again.
	  if(final >= 10) {
	  	final = final / 10;
	  }
	  postMessage(final);
	}
	// Upon receieving a message...recalculate because that means we fucked up.
	onmessage = function(e) {
		calc();
	}
	calc();
}