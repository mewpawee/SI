const express = require('express')
const convert = require('./convert.js')

const HOST = '0.0.0.0';
const PORT = 3000;
const app = express()

// parse application/json
app.use(express.json())

app.post('/generateReport', async (req, res) => {
    result = await convert.generateReport(req.body);
    res.send(result);
  })

app.listen(PORT,HOST, () => {
    console.log(`Example app listening at http://${HOST}:${PORT}`)
})

