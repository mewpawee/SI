const fs = require('fs');
const carbone = require('carbone');

// Data to inject
var data = require('./data.json'); //with path

var options = {
    convertTo : 'pdf:writer_pdf_Export' //can be docx, txt, ...
};

// Generate a report using the sample template provided by carbone module
// This LibreOffice template contains "Hello {d.firstname} {d.lastname} !"
// Of course, you can create your own templates!
carbone.render('CSI.pptx', data, function(err, result){
  console.log('t1');
  if (err) {
    return console.log(err);
  }
  // write the result
  fs.writeFileSync('result.pptx', result);
  console.log("finish");
  process.exit();
});

