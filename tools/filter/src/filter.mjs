import { readJSON, readFile } from "./function.mjs";
//each tools filter

const dnsmap = async (file) => {
  let json;
//   const data = await readFile(file);
//   const json = JSON.parse(data);
  return json;
};

const nmap = async (file) => {
  const filetype = file.split(".").pop();
  let json;
  if (filetype == "json") {
    json = await readJSON(file);
  }
  return json;
};

export { dnsmap, nmap };
