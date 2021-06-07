import { readJSON, readFile } from "./function.mjs";
import { nmapFilter } from "./nmap.mjs";
//each tools filter

const dnsmap = async (file) => {
  let json;
  //   const data = await readFile(file);
  //   const json = JSON.parse(data);
  // return json;
  return null;
};

const nmap = async (file) => {
  const filetype = file.split(".").pop();
  let filtered;
  if (filetype == "json") {
    const json = await readJSON(file);
    // console.log(json);
    filtered = await nmapFilter(json);
  }
  return filtered;
};

export { dnsmap, nmap };
