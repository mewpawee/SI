import { readJSON, readFile, command } from "./function.mjs";
import { nmapFilter } from "./nmap.mjs";
import { sqlmapFilter } from "./sqlmap.mjs";

//each tools filter

const sqlmap = async (file, name) => {
  const protocolPort=name.replace(".log", "");
  const data = await readFile(file);
  const Obj = await sqlmapFilter(data, protocolPort);
  if (!Obj) return undefined;
  return Obj;
};

const nmap = async (file) => {
  const filetype = file.split(".").pop();
  let filtered;
  if (filetype == "json") {
    const json = await readJSON(file);
    // console.log(json);
    filtered = await nmapFilter(json);
  }
  if (!filtered) return undefined;
  return filtered;
};

const dirb = async (file) => {
  // const cmd = `cat ${file} | awk -F" " '/http/{print $2}'`
  const cmd = `grep '+' ${file} | sed 's/+ /url-/;s/ (CODE:/,code-/;s/|SIZE:/,size-/;s/)//'`;
  const grepedData = await command(cmd);
  const obj = {};
  if (!grepedData) return undefined;
  const data = grepedData.split("\n").slice(0, -1);
  for (const d of data) {
    const arr = d.split(",");
    arr.forEach((string) => {
      const [key, value] = string.split("-");
      obj[key] = value;
    });
  }
  return obj;
};

export { sqlmap, nmap, dirb };
