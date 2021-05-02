import {readJSON,readFile} from './function.mjs'
//each tools filter

const dnsmap = async (file) => {
    const data = await readFile(file);
    const json = JSON.parse(data)
    return json
}

const nmap = async (file) => {
    const json = await readJSON(file);
    return json
}


export{dnsmap,nmap}
