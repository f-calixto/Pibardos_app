import { clientApi } from './index'

const baseUrl = '/auth'

const registerUser = async ({ email, password, username, birthdate, country }) => {
  const response = await clientApi.post(`${baseUrl}/register`, {
    email,
    password,
    username,
    birthdate,
    country
  })
  return response
}

const loginUser = async ({ email, password }) => {
  const response = await clientApi.post(`${baseUrl}/login`, {
    email,
    password
  })
  return response
}

export const authService = {
  registerUser,
  loginUser
}
