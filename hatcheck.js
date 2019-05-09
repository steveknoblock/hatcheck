/*
 hatcheck
 prototype for a CAS over HTTP
*/

var fs = require('fs');
var sha1 = require('sha1');
var path  = require('path');

var content = { 'name':'Fedora', 'color':'black' };

var id = sha1(content);


var getDirName = function (id) {
  return id.substr(0, 2);
}

var getFileName = function (id) {
  return id.slice(2);
}

var mkdirSync = function (path) {
  try {
    fs.mkdirSync(path);
  } catch(e) {
    if ( e.code != 'EEXIST' ) throw e;
  }
}

/* console.log(id); */

var dirName = getDirName(id);

/* console.log(dirName); */

var fileName = getFileName(id);

/* console.log(fileName); */

mkdirSync( path.join('objects', dirName) );

fs.writeFile(path.join('objects', dirName, fileName), content, function(err) {
    if(err) {
        return console.log(err);
    }
    console.log("Thank you for the content. Here's your ticket: " + id);
});

