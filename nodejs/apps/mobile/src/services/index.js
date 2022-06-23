import axios from 'axios'

const clientApi = axios.create({
  baseURL: 'https://api.pibardosapp.com/',
  timeout: 8000
})

export { clientApi }
