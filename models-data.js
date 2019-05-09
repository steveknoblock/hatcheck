const fs		= require('fs');
const sha1		= require('sha1');
const path		= require('path');


class Data {

	constructor(data) {
		this.cid = "cryptographichash";// = sha1(data);
		this.contentData = "Thanks for all the content.";
		this.dirName = this.getDirName(this.cid);
		this.fileName = this.getFileName(this.cid);
		console.log('Constructor Done');
	}

	toJSON() {
	    return {
	      id: this.cid,
	      data: this.CASdata
	    };
	  }

	save() {

		console.log(this.cid);
		console.log(this.contentData);
		console.log(this.dirName); // shard dir
		console.log(this.fileName); // rest of hash

		this.cid = sha1(this.contentData);

		this.mkdirSync( path.join('objects', this.dirName) );

		fs.writeFile(path.join('objects', this.dirName, this.fileName), this.contentData, function(err) {
			if(err) {
				return console.log(err);
			}

			//console.log("Thank you for the content. Here's your ticket: ");
			//console.log(this.cid);
			}
		);
		
		console.log("Thank you for the content. Here's your ticket: ");
		console.log(this.cid);
	}

	getData(cid) {

		// shard content id
		this.dirName = getDirName(cid);
		this.fileName = getFileName(cid);
		fs.readFile(path.join('objects', this.dirName, this.fileName), function(err, buf) { console.log(buf.toString()); });
	}

	getDirName(id) {
 		return id.substr(0, 2);
	}

	getFileName(id) {
  		return id.slice(2);
	}
	
	mkdirSync(path) {
	  try {
	    fs.mkdirSync(path);
	  } catch(e) {
	    if ( e.code != 'EEXIST' ) throw e;
	  }
	}

}

module.exports = Data;