
const fs = require('fs')



fs.readFile('/mnt/log/103.153.118.68/log/sqlmap/http:10254.log', null, (err, data) => {
    if (err) {
        console.error(err)
        return
    }
    
    var decoder = new TextDecoder('utf-8'),
        decodedMessage;
    decodedMessage = decoder.decode(data);
    var myObj = {}
    var startOS = decodedMessage.search('web server operating system: ')
    var pa = decodedMessage.search('Parameter')
    firstSection = decodedMessage.substring(pa, startOS)
    DatabaseInfoStr = decodedMessage.substring(startOS, decodedMessage.search('Table: '))
    var paramlist = []
    while (pa != -1) {
        var param = firstSection.search('Paramiter') + 12
        var firstSection = firstSection.substring(param, startOS)
        var paramInfo = {}
        var ty = firstSection.search('Type')
        pa = firstSection.search('Parameter')
        paramInfo['Prameter_name'] = firstSection.substring(0, firstSection.search('\r'))
        
        var injectlist = []
        while (((pa == -1) || (ty < pa)) && (ty != -1)){
            var inject = {}
            firstSection = firstSection.substring(ty, startOS)
            var ti = firstSection.search('Title')
            var pl = firstSection.search('Payload')
            inject['Type'] = firstSection.substring(6, firstSection.search('\r'))
            firstSection = firstSection.substring(ti, startOS)
            var pl = firstSection.search('Payload')
            inject['Title'] = firstSection.substring(7, firstSection.search('\r'))
            firstSection = firstSection.substring(pl, startOS)
            ty = firstSection.search('Type')
            pa = firstSection.search('Parameter')
            inject['Payload'] = firstSection.substring(8, firstSection.search('\r'))
            injectlist.push(inject)
        } 
        paramInfo['injection'] = injectlist
        paramlist.push(paramInfo)

    } 
    myObj['Parameter'] = paramlist
    myObj['OS'] = DatabaseInfoStr.substring(DatabaseInfoStr.search('web server operating system: ') + 29, DatabaseInfoStr.search('\r'))
    DatabaseInfoStr = DatabaseInfoStr.substring(DatabaseInfoStr.search('web application technology: '), 1000)
    myObj['Application'] = DatabaseInfoStr.substring(28, DatabaseInfoStr.search('\r'))
    DatabaseInfoStr = DatabaseInfoStr.substring(DatabaseInfoStr.search('back-end DBMS: '), 1000)
    myObj['DBMS'] = DatabaseInfoStr.substring(17, DatabaseInfoStr.search('\r'))
    DatabaseInfoStr = DatabaseInfoStr.substring(DatabaseInfoStr.search('Database: '), 1000)
    myObj['Database name'] = DatabaseInfoStr.substring(10, DatabaseInfoStr.search('\r'))
    var str = JSON.stringify(myObj)
    console.log(str)
})