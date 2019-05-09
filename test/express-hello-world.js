const express = require('express')
const app = express()
const port = 3000

// The req (request) and res (response) are the exact same objects that Node provides, so you can invoke req.pipe(), req.on('data', callback), and anything else you would do without Express involved.

/*
	Route definition

	app.HTTP_METHOD(PATH, HANDLER_FUNCTION)

*/

app.get('/', (req, res) => res.send('Hello World!'))

app.get('/about', function (req, res) {
  res.send('about')
})

app.get('/users/:userId/books/:bookId', function (req, res) {
  res.send(req.params)
})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))

