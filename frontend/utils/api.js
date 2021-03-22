import axios from 'axios'
const getToken = function() {
  if (process.server) {
    // server side
    return
  }
  if (window.$nuxt) {
    return window.$nuxt.$auth.strategy.token.get()
  }
}

export async function request(method, url, data, auth = false) {
  const headers = {}
  headers['Content-Type'] = 'application/json'
  if (auth) {
    headers.Authorization = await getToken()
  }
  try {
    // call api
    const response = await axios({
      method,
      url,
      headers,
      data: JSON.stringify(data)
    })
    return response.data
  } catch (e) {
    return e
  }
}
