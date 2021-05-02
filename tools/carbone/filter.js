const fs = require('fs');
//var input_data = process.argv.slice(2);
var input_data = process.argv.slice(2);
input_data = input_data.toString();

const read = (input) => {
  return new Promise((fullfill, reject) => {
    fs.readFile(input_data, 'utf8', function (err, data) {
      if (err) throw err;
      const tmp = JSON.parse(data);
      for (k = 0; k < tmp.nmap.length; k++) {
        let tmp1 = tmp.nmap[k].nmaprun;
        delete tmp1.scaninfo;
        delete tmp1.verbose;
        delete tmp1.debugging;
        delete tmp1.taskbegin;
        delete tmp1.taskend;
        delete tmp1.taskprogress;

        for (i = 0; i < tmp1.host.length; i++) {
          let tmp2 = tmp1.host[i];
          for (j = 0; j < tmp2.ports.port.length; j++) {
            if (tmp2.ports.port[j].service === undefined) {
              tmp2.ports.port[j].service = {};
              tmp2.ports.port[j].service.name = "unknown";
            }
          }
        }
      }
      fullfill(tmp)
    })

  })
}

const filter = async (input) => {
  const data = await read(input)
  const pt = JSON.stringify(data, null, 4);
  fs.writeFileSync('report.json', pt, (err) => {
    if (err) throw err;
  });
  console.log('JSON file filtered');
}

filter(input_data)