import fs from "fs";
import path from "path";
import { writeJSON } from "./src/function.mjs";
import { sqlmap, nmap, dirb } from "./src/filter.mjs";

const readEndpoint = async (dir, endpoint) => {
  const json = {
    endpoint: endpoint,
    tools: [],
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
      console.log("read: " + entry);
      switch (tool) {
        case "nmap":
          const nmapObj = { name: "nmap", data: [] };
          const nmapResult = await nmap(entry);
          if (nmapResult) {
            nmapObj.data.push(nmapResult);
            json.tools.push(nmapObj);
          }
          break;
        case "dirb":
          const dirbObj = { name: "dirb", data: [] };
          const dirbResult = await dirb(entry);
          if (dirbResult) {
            dirbObj.data.push(dirbResult);
            json.tools.push(dirbObj);
          }
          break;
        case "sqlmap":
          const sqlmapObj = { name: "sqlmap", data: [] };
          const sqlmapResult = await sqlmap(entry, d.name);
          if (sqlmapResult) {
            sqlmapObj.data.push(sqlmapResult);
            json.tools.push(sqlmapObj);
          }
          break;
      }
    }
  }
  return json;
};

// main program
const main = async (
  company,
  endpoints,
  output = "/tmp/report.json",
  location = "/mnt/log/"
) => {
  const trimedEndpoints = endpoints.replace(/^\[+|\]+$/g, "");
  const arrayEndpoints = trimedEndpoints.split(",");
  const reportObj = {
    company: company,
  };
  const endpointsData = [];
  for (const endpoint of arrayEndpoints) {
    const dir = location + endpoint + "/log";
    const result = await readEndpoint(dir, endpoint);
    endpointsData.push(result);
  }
  reportObj.endpoints = endpointsData;
  writeJSON(output, reportObj);
};

// print process.argv
const args = process.argv.slice(2);
main(args[0], args[1], args[2], args[3]);
