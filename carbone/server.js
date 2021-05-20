const express = require('express')
const convert = require('./convert.js')
const app = express()
const HOST = '0.0.0.0';
const PORT = 3000;

app.get('/', async (req, res) => {
    result = await convert.generateReport();
    res.send(result);
  })

app.listen(PORT,HOST, () => {
    console.log(`Example app listening at http://${HOST}:${PORT}`)
})

