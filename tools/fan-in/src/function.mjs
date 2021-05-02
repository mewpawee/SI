import fs from 'fs'

//read text file
const readFile = (filename) => {
    return new Promise((fullfill,reject)=>{
        fs.readFile(filename,"utf-8",(err,res)=>{
            if(err) reject(err);    
            else{
                fullfill(res)
            }
        });
    });
}

//read jsonfile
const readJSON = (filename) => {
    return new Promise((fullfill,reject)=>{
        fs.readFile(filename,(err,res)=>{
            if(err) reject(err);    
            else{
                const result =  JSON.parse(res);
                fullfill(result)
            }
        });
    });
}

//parse string and write to json file
const writeJSON = (filename,data) => {
    return new Promise((_,reject) => {
        fs.writeFile(filename, JSON.stringify(data,null,2),(err)=>{
            if(err) reject(err);
            console.log("write: " + filename)
            console.log(data)
        });
    });
}

export {readFile,readJSON, writeJSON}