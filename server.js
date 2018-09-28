/** 
 * Hatcheck/CAS/Scratchpad/Lambda Server
 */

var http = require('http');
var url  = require('url');

 var server = http.createServer(function(request, response) {

/* 
	Accept JSON from endpoint, inspect the object, do what the command says
 */

	console.log('STATUS: ' + response.statusCode);
	//console.log('HEADERS: ' + JSON.stringify(response.request));
	console.log('HEADERS: ' + JSON.stringify(response.headers));


	console.log('REQUEST HEADERS' + JSON.stringify(request.headers));
	console.log('REQUEST URL' + request.url);

var url_parts = url.parse(request.url, true);
    var query = url_parts.query;

    console.log(query); //{Object}

  response.writeHead(200);
  response.end('Hello. This is Hatcheck.');
});

server.listen(8080);
