const nmapFilter = (data) => {
  return new Promise((fullfill, reject) => {
    const tmp = data;
    let tmp1 = tmp.nmaprun;
    delete tmp1.scaninfo;
    delete tmp1.verbose;
    delete tmp1.debugging;
    delete tmp1.taskbegin;
    delete tmp1.taskend;
    delete tmp1.taskprogress;

    for (let i = 0; i < tmp1.host.length; i++) {
      let tmp2 = tmp1.host[i];
      for (let j = 0; j < tmp2.ports.port.length; j++) {
        if (tmp2.ports.port[j].service === undefined) {
          tmp2.ports.port[j].service = {};
          tmp2.ports.port[j].service.name = "unknown";
        }
      }
    }
    fullfill(tmp);
  });
};

export { nmapFilter };
