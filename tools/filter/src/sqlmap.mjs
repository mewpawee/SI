const sqlmapFilter = (data,name) => {
  return new Promise((fullfill, reject) => {
    const myObj = {};
    myObj["protocolPort"] = name;
    const startOS = data.search("web server operating system: ");
    let pa = data.search("Parameter");
    let firstSection = data.substring(pa, startOS);
    let DatabaseInfoStr = data.substring(startOS, data.search("Table: "));
    let paramlist = [];
    while (pa != -1) {
      let param = firstSection.search("Paramiter") + 12;
      firstSection = firstSection.substring(param, startOS);
      let paramInfo = {};
      let ty = firstSection.search("Type");
      pa = firstSection.search("Parameter");
      paramInfo["Prameter_name"] = firstSection.substring(
        0,
        firstSection.search("\r")
      );

      const injectlist = [];
      while ((pa == -1 || ty < pa) && ty != -1) {
        const inject = {};
        firstSection = firstSection.substring(ty, startOS);
        const ti = firstSection.search("Title");
        let pl = firstSection.search("Payload");
        inject["Type"] = firstSection.substring(6, firstSection.search("\r"));
        firstSection = firstSection.substring(ti, startOS);
        pl = firstSection.search("Payload");
        inject["Title"] = firstSection.substring(7, firstSection.search("\r"));
        firstSection = firstSection.substring(pl, startOS);
        ty = firstSection.search("Type");
        pa = firstSection.search("Parameter");
        inject["Payload"] = firstSection.substring(
          8,
          firstSection.search("\r")
        );
        injectlist.push(inject);
      }
      paramInfo["injection"] = injectlist;
      paramlist.push(paramInfo);
    }
    myObj["Parameter"] = paramlist;
    myObj["OS"] = DatabaseInfoStr.substring(
      DatabaseInfoStr.search("web server operating system: ") + 29,
      DatabaseInfoStr.search("\r")
    );
    DatabaseInfoStr = DatabaseInfoStr.substring(
      DatabaseInfoStr.search("web application technology: "),
      1000
    );
    myObj["Application"] = DatabaseInfoStr.substring(
      28,
      DatabaseInfoStr.search("\r")
    );
    DatabaseInfoStr = DatabaseInfoStr.substring(
      DatabaseInfoStr.search("back-end DBMS: "),
      1000
    );
    myObj["DBMS"] = DatabaseInfoStr.substring(17, DatabaseInfoStr.search("\r"));
    DatabaseInfoStr = DatabaseInfoStr.substring(
      DatabaseInfoStr.search("Database: "),
      1000
    );
    myObj["Database name"] = DatabaseInfoStr.substring(
      10,
      DatabaseInfoStr.search("\r")
    );
    fullfill(myObj);
  });
};

export { sqlmapFilter };
