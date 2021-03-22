import { request } from './api'

export const getEndpoints = async () => {
  const poolUrl = 'https://csi.cmkl.ac.th/api/privateUser/getPools'

  try {
    const poolResponse = await request('get', poolUrl, {}, true)
    const poolId = poolResponse.data[0].poolid
    const endpointsUrl = `https://csi.cmkl.ac.th/api/privateUser/getEndpoints?poolid=${poolId}`
    const test = await request('get', endpointsUrl, {}, true)
    return test
  } catch (err) {
    throw new Error('cant submit')
  }
}

export const addEndpoint = async (endpoint) => {
  const poolUrl = 'https://csi.cmkl.ac.th/api/privateUser/getPools'

  try {
    const poolResponse = await request('get', poolUrl, {}, true)
    const poolId = poolResponse.data[0].poolid
    const addEndpointsUrl = `https://csi.cmkl.ac.th/api/privateUser/addEndpoint`
    const body = {}
    body.poolid = poolId
    body.endpoint = endpoint
    const test = await request('post', addEndpointsUrl, body, true)
    return test
  } catch (err) {
    throw new Error('cant submit')
  }
}

export const deleteEndpoint = async (endpoint) => {
  try {
    const deleteEndpointsUrl = `https://csi.cmkl.ac.th/api/privateUser/deleteEndpoint`
    const body = {}
    body.endpoint = endpoint
    const test = await request('post', deleteEndpointsUrl, body, true)
    return test
  } catch (err) {
    throw new Error('cant submit')
  }
}
