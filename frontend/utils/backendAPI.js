import { request } from './api'
export const addEndpoint = async (googleid, endpoint) => {
  const url = 'https://csi.cmkl.ac.th/api/addEndpoint'
  const body = {
    endpoint: `'${googleid}'`,
    googleid: `'${endpoint}'`
  }
  try {
    const response = await request('post', url, body, true)
    return response
  } catch (err) {
    throw new Error('cant submit')
  }
}
