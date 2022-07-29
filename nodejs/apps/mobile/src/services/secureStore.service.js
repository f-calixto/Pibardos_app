import * as SecureStore from 'expo-secure-store'

/**
 * Save user data in SecureStore
 * @param {*} userData
 * @returns
 */
const saveUser = async userData => await SecureStore.setItemAsync('user', JSON.stringify(userData))

/**
 * Get user data from SecureStore
 * @returns
 */
const getUser = async () => {
  const user = await SecureStore.getItemAsync('user')
  return JSON.parse(user)
}

export const secureStoreService = {
  saveUser,
  getUser
}
