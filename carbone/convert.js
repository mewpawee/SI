const fs = require('fs');
const carbone = require('carbone');

// Data to inject
var data = require('./data.json'); //with path

var options = {
    convertTo : 'pdf' //can be docx, txt, ...
};

// Generate a report using the sample template provided by carbone module
// This LibreOffice template contains "Hello {d.firstname} {d.lastname} !"
// Of course, you can create your own templates!
carbone.render('template.pptx', data, options, function(err, result){
  console.log('t1');
  if (err) {
    return console.log(err);
  }
  // write the result
  fs.writeFileSync('result.pdf', result);
  console.log("finish");
  process.exit();
});
