import { readJSON, readFile, command } from "./function.mjs";
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

const dirb = async (file) => {
  // const cmd = `cat ${file} | awk -F" " '/http/{print $2}'`
  const cmd = `grep '+' ${file} | sed 's/+ /url-/;s/ (CODE:/,code-/;s/|SIZE:/,size-/;s/)//'`;
  const grepedData = await command(cmd);
  const obj = {};
  if (grepedData) {
    const data = grepedData.split("\n").slice(0, -1);
    for (const d of data) {
      const arr = d.split(",");
      arr.forEach((string) => {
        //console.log(arr)
        const [key, value] = string.split("-");
        obj[key] = value;
      });
    }
  } else {
    return undefined;
  }
  return obj;
};

export { dnsmap, nmap, dirb };
