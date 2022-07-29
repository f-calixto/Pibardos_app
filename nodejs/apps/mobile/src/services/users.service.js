import { clientApi } from './index'

const baseUrl = '/users'

const getUser = async id => {
  const response = await clientApi.get(`${baseUrl}/${id}`)
  return response
}

export const usersService = {
  getUser
}
