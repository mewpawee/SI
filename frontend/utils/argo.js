import { request } from './api'
const url = 'https://csi.cmkl.ac.th/express/api/argo'
export const argoSubmit = async (targetDomains) => {
  try {
    const response = await request('post', url, targetDomains, true)
    return response
  } catch (err) {
    throw new Error('cant submit')
  }
}
