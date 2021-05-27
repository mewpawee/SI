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

export async function download(url, data, auth = false) {
  const headers = {}
  headers.Accept = '*/*'
  headers['Content-type'] = 'application/pdf'
  if (auth) {
    headers.Authorization = await getToken()
  }
  try {
    // call api
    const response = await axios({
      method: 'post',
      url,
      headers,
      data: JSON.stringify(data),
      responseType: 'blob'
    }).then((response) => {
      const url = window.URL.createObjectURL(new Blob([response.data]))
      const link = document.createElement('a')
      link.href = url
      link.setAttribute('download', `${data.scanid}.pdf`)
      document.body.appendChild(link)
      link.click()
    })
    return response.data
  } catch (e) {
    return e
  }
}
