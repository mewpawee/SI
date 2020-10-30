import axios from 'axios';

export const startArgo = async (method, url, queryData) => {
  const headers = {};
  headers['Content-Type'] = 'application/json';
  try {
    const response = await axios({
      method,
      url,
      data: JSON.stringify(queryData),
      headers,
      proxy: { host: 'argo-server', port: 2746 },
    });
    console.log(response);
    return response.data;
  } catch (err) {
    return err;
  }
};
