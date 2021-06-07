import fs from "fs";
import path from "path";
import { writeJSON } from "./src/function.mjs";
import { dnsmap, nmap } from "./src/filter.mjs";

const readEndpoint = async (dir) => {
  const json = {
    dnsmap: [],
    nmap: [],
  };
  const result = await read(dir,json);
  for (const key in result) {
    if (result[key].length == 0) {
      result[key] = null;
    }
  }
  return result
}
//read file
const read = async (dir, json, tool = null) => {
  for await (const d of await fs.promises.opendir(dir)) {
    const entry = path.join(dir, d.name);
    if (d.isDirectory()) {
      await read(entry, json,d.name);
    } else if (d.isFile()) {
      //console.log("read: " + entry);
      switch (tool) {
        case "dnsmap":
          const result = await dnsmap(entry);
          if (result == null) {
            json[tool] = null;
          } else {
            json.dnsmap.push(result);
          }
          break;
        case "nmap":
          const test = await nmap(entry);
          if (test) {
            json.nmap.push(test);
          }
          break;
      }
    }
  }
  return json;
};

// main program
const main = async (endpoints, company, location = "/mnt/log/") => {
  const trimedEndpoints = endpoints.replace(/^\[+|\]+$/g, '')
  const arrayEndpoints = trimedEndpoints.split(',')
  const reportObj = {
    company: company
  }
  const endpointsData = []
  for(const endpoint of arrayEndpoints){
    const dir = location+endpoint+"/log"
    const result = await readEndpoint(dir);
    const endpointObj = {
      [endpoint] : result
    }
    endpointsData.push(endpointObj)
  }
  reportObj.endpoints = endpointsData
  console.log(reportObj)
  // writeJSON("/tmp/report.json", result);
};

// print process.argv
const args = process.argv.slice(2);
main(args[0],args[1]);
