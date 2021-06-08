import { request, download } from './api'

export const adminGetEndpoints = async () => {
  try {
    const endpointUrl = `https://csi.cmkl.ac.th/api/getAllEndpointsAdmin`
    const result = await request('get', endpointUrl, {}, true)
    return result
  } catch (err) {
    throw new Error(`can't get endpoints`)
  }
}
export const downloadReport = async (scanId) => {
  try {
    const endpointUrl = 'https://csi.cmkl.ac.th/api/generateReport'
    const body = {}
    body.scanid = scanId
    const result = await download(endpointUrl, body, true)
    return result
  } catch (err) {
    throw new Error(`can't download report`)
  }
}
export const getAllScans = async () => {
  try {
    const endpointUrl = `https://csi.cmkl.ac.th/api/privateUser/getAllScans`
    const result = await request('get', endpointUrl, {}, true)
    return result
  } catch (err) {
    throw new Error('cant retrieve')
  }
}

export const getEndpoints = async () => {
  try {
    const endpointsUrl = `https://csi.cmkl.ac.th/api/privateUser/getEndpoints`
    const result = await request('get', endpointsUrl, {}, true)
    return result
  } catch (err) {
    throw new Error('cant submit')
  }
}

export const addEndpoint = async (endpoint) => {
  try {
    const addEndpointsUrl = `https://csi.cmkl.ac.th/api/privateUser/addEndpoint`
    const body = {}
    body.endpoint = endpoint
    const result = await request('post', addEndpointsUrl, body, true)
    return result
  } catch (err) {
    throw new Error('cant submit')
  }
}

export const deleteEndpoint = async (endpoint) => {
  try {
    const deleteEndpointsUrl = `https://csi.cmkl.ac.th/api/privateUser/deleteEndpoint`
    const body = {}
    body.endpoint = endpoint
    const result = await request('post', deleteEndpointsUrl, body, true)
    return result
  } catch (err) {
    throw new Error('cant submit')
  }
}
