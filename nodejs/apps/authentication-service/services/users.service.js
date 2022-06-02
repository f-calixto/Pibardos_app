const ServiceError = require('../utils/ServiceError')
const { SERVICE_ERRORS } = require('../constants/errors')

class UsersService {
  constructor (UserModel) {
    this.UserModel = UserModel
  }

  /**
  * Register an user.
  * @param {Object} user - The user who will be registered.
  * @param {string} user.email - The user's email.
  * @param {string} user.username - The user's username.
  * @param {string} user.password - The user's password.
  * @param {string} user.birthdate - The user's birthdate.
  * @param {string} user.country - The user's country.
  * @returns {Promise} Promise object represents the registered user.
  */
  async register ({ email, password, username, birthdate, country }) {
    if (!birthdate || !country) {
      throw new Error('Birthdate or country is missing')
    }

    const user = await this.UserModel.create({
      email,
      password,
      username
    })

    return user
  }

  /**
   * Authenticate an user.
   * @param {Object} credentials - The user's credentials.
   * @param {string} credentials.email - The user's email.
   * @param {string} credentials.password - The user's password.
   * @returns {Promise} Promise object represents the authenticated user.
   */

  async authenticate ({ email, password }) {
    const user = await this.UserModel.findOne({ email })

    if (!user) {
      throw new ServiceError(SERVICE_ERRORS.USER_NOT_FOUND, 'email', 'User not found')
    }

    if (!await user.comparePassword(password)) {
      throw new ServiceError(SERVICE_ERRORS.INVALID_CREDENTIALS, 'email', 'Email or password is invalid')
    }

    return user
  }
}

module.exports = UsersService
