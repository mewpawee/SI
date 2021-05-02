import fs from 'fs'
import path from 'path'
import {writeJSON} from './src/function.mjs'
import {dnsmap,nmap} from './src/filter.mjs'

//global variable
const json = {
    nmap:null
}

//read file
const read = async(dir,tool=null)=>{   
    for await(const d of await fs.promises.opendir(dir)){
        const entry = path.join(dir,d.name);
        if(d.isDirectory()){
            await read(entry,d.name)
        }else if (d.isFile()){
            console.log("read: " + entry);
            if(json[tool] == null){
                json[tool] = []
            }
            switch(tool){
                case "dnsmap":
                    const result = await dnsmap(entry)
                    json.dnsmap.push(result)
                    break;
                case "nmap":
                    const test = await nmap(entry)
                    json.nmap.push(test)
                    break;
            }
        }
    }

    return json
}

// main program
const main = async(location="./report") => {
    const result = await read(location)
    writeJSON('/tmp/report.json',result)
}


// print process.argv
const args = process.argv.slice(2);
if(args.length > 0 && args[0] == "-d"){
    main(args[1])
}else{
    main()
}

