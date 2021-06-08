import fs from "fs";
import path from "path";
import { writeJSON } from "./src/function.mjs";
import { sqlmap, nmap, dirb } from "./src/filter.mjs";

const readEndpoint = async (dir) => {
  const json = {
    nmap: [],
    dirb: [],
    sqlmap: [],
  };
  const result = await read(dir, json);
  for (const key in result) {
    if (result[key].length == 0) {
      result[key] = null;
    }
  }
  return result;
};
//read file
const read = async (dir, json, tool = null) => {
  for await (const d of await fs.promises.opendir(dir)) {
    const entry = path.join(dir, d.name);
    if (d.isDirectory()) {
      await read(entry, json, d.name);
    } else if (d.isFile()) {
      //console.log("read: " + entry);
      switch (tool) {
        case "nmap":
          const nmapResult = await nmap(entry);
          if (nmapResult) {
            json.nmap.push(nmapResult);
          }
          break;
        case "dirb":
          const dirbResult = await dirb(entry);
          if (dirbResult) {
            json.dirb.push(dirbResult);
          }
          break;
        case "sqlmap":
          const sqlmapResult = await sqlmap(entry,d.name);
          if (sqlmapResult) {
            json.sqlmap.push(sqlmapResult);
          }
          break;
      }
    }
  }
  return json;
};

// main program
const main = async (company, endpoints, location = "/mnt/log/") => {
  const trimedEndpoints = endpoints.replace(/^\[+|\]+$/g, "");
  const arrayEndpoints = trimedEndpoints.split(",");
  const reportObj = {
    company: company,
  };
  const endpointsData = [];
  for (const endpoint of arrayEndpoints) {
    const dir = location + endpoint + "/log";
    const result = await readEndpoint(dir);
    const endpointObj = {
      [endpoint]: result,
    };
    endpointsData.push(endpointObj);
  }
  reportObj.endpoints = endpointsData;
  // writeJSON("/tmp/report.json", result);
  writeJSON("./report.json", reportObj);
};

// print process.argv
const args = process.argv.slice(2);
main(args[0], args[1]);
