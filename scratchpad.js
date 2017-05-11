// Scratchpad

// template string to support multiline strings
var data = `[{
    "name": "qBert",
    "arcade": "Yes",
    "color": "Yes",
    "note": "none"
  },
  {
    "name": "Frogger",
    "arcade": "Yes",
    "color": "Yes",
    "note": "Fun frog hopping across a busy road and dangerous river."
  },
  {
    "name": "Xevious",
    "arcade": "Yes",
    "color": "Yes",
    "note": "2D vertical scroller shooter with deep game play."
  },
  {
    "name": "Galaga",
    "arcade": "Yes",
    "color": "Yes",
    "note": "A twist on Space Invaders."
  },
  {
    "name": "Pong",
    "arcade": "Yes",
    "color": "No",
    "note": "The first popular video arcade game."
  }
]`;


// Test
/*
try{
    var obj = JSON.parse(data);
}
catch(e){
   throw new Error("Sorry, man, I could not parse your JSON data.");
  }
console.log(obj);
*/

// sans fluent way of writing this
var myArray = JSON.parse(data);
var result = myArray.filter(function (something) {
  //console.log(something);
  //console.log(something.color);
  return something.color === 'No';
});

// fluent way of writing the query
var result = JSON.parse(data).filter(function(obj) {
  return obj.color === "No";
});
console.log(result);



/*
var result = JSON.parse(data).filter(obj, function () {
	return obj.color === 'No';
});
console.log(result);
*/

