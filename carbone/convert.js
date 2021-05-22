const carbone = require('carbone');

var options = {
    convertTo : 'pdf' //can be docx, txt, ...
};

const carboneGenerate = (data)=>{
  return new Promise((resolve,reject)=>{
    carbone.render('CSI.pptx', data,options, (err, result) => {
      if (err) {
        reject(err);
      }
      //return the result
      resolve(result);
    });
  })
}
exports.generateReport = async(dataInput)=>{
  let pdfFile = await carboneGenerate(dataInput)
  return pdfFile
}


