const xml2js = require('xml2js');
var fs = require('fs');
// parse input from list to string
var xml = process.argv.slice(2);
xml = xml.toString();

var parser = new xml2js.Parser();
fs.readFile(xml, function(err, data){
    // convert XML to JSON
    if(err) throw err;
    parser.parseString(data, (err, result) => {
        if(err) throw err;
        // convert it to a JSON string
        const json = JSON.stringify(result, null, 4);
        // write file
        fs.writeFile('data.json', json, (err) => {
            if (err) throw err;
        });
    });
})
