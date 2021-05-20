const carbone = require('carbone');

var options = {
    convertTo : 'pdf' //can be docx, txt, ...
};


const getData = () => {
    var data = require('./report.json'); //with path
    return data
}
const carboneGenerate = new Promise((resolve,reject)=>{
  data = getData()
  carbone.render('CSI.pptx', data,options, (err, result) => {
    if (err) {
      reject(err);
    }
    //return the result
    resolve(result);
  });
})

exports.generateReport = async()=>{
  let pdfFile = await carboneGenerate
  return pdfFile
}


