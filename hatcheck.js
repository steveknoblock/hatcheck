/*
 hatcheck
 prototype for a CAS over HTTP
*/

var fs = require('fs');
var sha1 = require('sha1');
var path  = require('path');

var content = 'asdfasdfasdf';

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

console.log(id);

var dirName = getDirName(id);

console.log(dirName);

var fileName = getFileName(id);

console.log(fileName);

mkdirSync( path.join(dirName) );

fs.writeFile(path.join(dirName, fileName), "content", function(err) {
    if(err) {
        return console.log(err);
    }

    console.log("The file was saved!");
});




mkdirSync( path.join('first') );
mkdirSync( path.join('first', 'second') );

fs.writeFile("hey", "Hey there!", function(err) {
    if(err) {
        return console.log(err);
    }

    console.log("The file was saved!");
}); 

/*
$file_path = $this->pathToObjects;
		$id = sha1($content);

		//echo "<p>Hash for object is: {$id}</p>\n";

		$dirname = substr($id, 0, 2);
		$filename = substr($id, 2, 40-2);
		*/