/**
 * Hatcheck Server Module
 * CAS over HTTP
 */

var http = require("http");


var props,
Hatcheck = {
 	
 	properties: { 
 		objectDir:'objects'
 	},

 	init: function() {
 		props = properties;
 	}
};


http.createServer(function (request, response) {

	request.
	console.log('STATUS: ' + res.statusCode);
  	console.log('HEADERS: ' + JSON.stringify(res.headers));

	response.writeHead(200, {'Content-Type': 'text/plain'});
	response.end('Hello World\n');
}).listen(8081);

