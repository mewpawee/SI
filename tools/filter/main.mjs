import fs from "fs";
import path from "path";
import { writeJSON } from "./src/function.mjs";
import { sqlmap, nmap, dirb } from "./src/filter.mjs";

const readEndpoint = async (dir, endpoint) => {
  const json = {
    endpoint: endpoint,
    tools: [
      { name: "nmap", data: [] },
      { name: "dirb", data: [] },
      { name: "sqlmap", data: [] },
    ],
  };
  const result = await read(dir, json);
  result.tools = result.tools.filter((value, index, arr) => {
    return value.data.length > 0;
  });
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
          const nmapResult = await nmap(entry);
          if (nmapResult) {
            json.tools[0].data.push(nmapResult);
          }
          break;
        case "dirb":
          const dirbResult = await dirb(entry);
          if (dirbResult) {
            json.tools[1].data.push(dirbResult);
          }
          break;
        case "sqlmap":
          const sqlmapResult = await sqlmap(entry, d.name);
          if (sqlmapResult) {
            json.tools[2].data.push(sqlmapResult);
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
