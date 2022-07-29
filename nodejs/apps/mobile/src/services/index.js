import axios from 'axios'
import { secureStoreService } from './secureStore.service'

const clientApi = axios.create({
  baseURL: 'https://api.pibardosapp.com/',
  timeout: 8000
})

// Add accessToken in header on all requests
clientApi.interceptors.request.use(async function (config) {
  const { accessToken } = await secureStoreService.getUser()
  config.headers.Authorization = `Bearer ${accessToken}`

  return config
})

export { clientApi }
