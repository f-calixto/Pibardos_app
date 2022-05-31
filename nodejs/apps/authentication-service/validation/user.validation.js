const Joi = require('joi')

const createUser = Joi.object({
  email: Joi.string().email().required(),
  password: Joi.string().required(),
  username: Joi.string().required(),
  birthdate: Joi.date().required(),
  country: Joi.string().required()
})

module.exports = {
  createUser
}
