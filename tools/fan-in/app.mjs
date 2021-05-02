import fs from 'fs'
import path from 'path'
import {writeJSON} from './function.mjs'
import {dnsmap,nmap} from './filter.mjs'


//global variable
const json = {
    dnsmap:[],
    nmap:[]
}

//read file
const read = async(dir,tool=null)=>{   
    for await(const d of await fs.promises.opendir(dir)){
        const entry = path.join(dir,d.name);
        console.log(entry,d.name,tool)
        if(d.isDirectory()){
            await read(entry,d.name)
        }else if (d.isFile()){
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
    const result = await read("./log")
    console.log(result)
    writeJSON('report.json',result)
}

main()

