import fs from 'fs'
import path from 'path'
import {writeJSON} from './function.mjs'
import {dnsmap,nmap} from './filter.mjs'

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
const main = async() => {
    const result = await read("./report")
    writeJSON('report.json',result)
}

main()

