import axios from 'axios';
import { argoServer } from '../settings';
export const startArgo = async (method, url, queryData) => {
  const headers = {};
  headers['Content-Type'] = 'application/json';
  try {
    // console.log(argoServer);
    const response = await axios({
      method,
      url,
      data: JSON.stringify(queryData),
      headers,
      proxy: { host: `${argoServer}`, port: 2746 },
    });
    return response.data;
  } catch (err) {
    throw new Error('failed to startArgo');
  }
};
