import jwtdecode from 'jwt-decode'

/**
 * Verify if authorization token is valid.
 *
 * @returns {Boolean} Boolean
 */
export default function (token) {
  try {
    const decodedToken = jwtdecode(token)
    const expToken = decodedToken.exp * 1000
    return expToken > new Date()
  } catch (error) {
    return false
  }
}
