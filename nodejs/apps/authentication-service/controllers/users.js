const UsersService = require('../services/users.service')
const User = require('../models/user')
const { SERVICE_ERRORS } = require('../constants/errors')

const usersService = new UsersService(User)

module.exports = {
  register: async (req, res, next) => {
    const { email, password, username, birthdate, country } = req.body

    try {
      const user = await usersService.register({ email, password, username, birthdate, country })
      return res.status(201).json(user)
    } catch (err) {
      next(err)
    }
  },

  authenticate: async (req, res, next) => {
    const { email, password } = req.body

    try {
      const user = await usersService.authenticate({ email, password })
      return res.status(200).json(user)
    } catch (err) {
      if (err.name === SERVICE_ERRORS.USER_NOT_FOUND) { return res.status(404).json({ errors: err.errors }) }
      if (err.name === SERVICE_ERRORS.INVALID_CREDENTIALS) { return res.status(401).json({ errors: err.errors }) }
      next(err)
    }
  }
}
