// hatcheck router.js
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => res.send('Hello World!'))

app.get('/:foo', (req, res, next) => {
    res.send(req.params)
  });


app.listen(port, () => console.log(`Example app listening on port ${port}!`))
