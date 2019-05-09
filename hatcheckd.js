/**
 * Hatcheckd Server Module
 * CAS over HTTP
 */

import Data from "./models-data.js";
import { createServer } from "http";

const hostname = 'localhost';
const port ='8080';

/*
var props,
Hatcheck = {
 	
 	properties: { 
 		objectDir:'objects'
 	},

 	init: function() {
 		props = properties;
 	}
};
*/

const data = new Data();
console.log(data);
var content = { 'name':'Fedora', 'color':'black' };

//const data = new Data();
//console.log(data);

var id = data.save(content);

data.getData(id);

createServer(function (request, response) {

	//request.
	//console.log('STATUS: ' + res.statusCode);
  	//console.log('HEADERS: ' + JSON.stringify(res.headers));

	response.writeHead(200, {'Content-Type': 'text/plain'});
	response.write("Hi");
	response.end('Hello World\n');


}).listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
}
)
