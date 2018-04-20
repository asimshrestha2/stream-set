var logo = document.getElementById('logo');
var circle = document.querySelector('.circle');
var radius = circle.getAttribute('r');
var c = 2 * Math.PI * radius;

circle.style.strokeDasharray = c + "";

function setPercentage(num){
  if (num < 0) { num = 0;}
  if (num > 100) { num = 100;}

  var pct = ((100-num)/100)*c;
  
  circle.style.strokeDashoffset = pct + "";
}

function setAnimationSpeed(num){
  logo.style.animationDuration = num+"s";
}

var lastTime = Date.now();
var id = null;
var idrunning = false;

function start(){
  lastTime = Date.now();
  id = setInterval(frame, 100);
  idrunning = true;
}

var totalTime = 60000;
function frame() {
  var currentTime = Date.now();
  var timeDiff = currentTime - lastTime;
  
  if(!logo.classList.contains("animate-fast")){
    logo.classList.remove("animate-normal");
    logo.classList.add("animate-fast");
  }
  
  if (timeDiff >= totalTime) {
    setPercentage(0);
    if(!logo.classList.contains("animate-normal")){
      logo.classList.add("animate-normal");
      logo.classList.remove("animate-fast");
    }
    // setAnimationSpeed(5);
    idrunning = false;
    clearInterval(id);
  } else {
    // setAnimationSpeed((.5) +( (timeDiff/totalTime) * 4.5));
    setPercentage((1-(timeDiff/totalTime))*100);
  }
  // console.log(timeDiff == totalTime)
}

function restartTimer(){
  if(idrunning){
    lastTime = Date.now();
  } else {
    lastTime = Date.now();
    id = setInterval(frame, 100);
    idrunning = true;
  }
}