// hatcheck router.js
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => res.send('Hello World!'))

app.get('/:foo', (req, res, next) => {
    res.send(req.params)
  });


/*
	I want to make it possible to config a path to the query endpoint

	var endpoint = 'foo/bar/baz';

	This is not for navigation of data but for seperation of concerns.

	Default to POST method.


*/

var endpoint = 'foo/bar/baz';

app.post('/' + endpoint, (req, res, next) => {
	/* express can handle a HTTP body that consists entirely of JSON
	   and no URL-encoded data
	  */
	app.use(express.json()); // for parsing application/json

	console.log(req.body)
  res.json(req.body)
    res.send('POST request to the homepage')
  });


app.listen(port, () => console.log(`Example app listening on port ${port}!`))
