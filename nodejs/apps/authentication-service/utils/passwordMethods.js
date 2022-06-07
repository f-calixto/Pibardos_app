const bcrypt = require('bcryptjs')
const config = require('../config')

const hash = string => bcrypt.hash(string, config.BCRYPT_SALT_ROUNDS)

const compare = (string, hash) => bcrypt.compare(string, hash)

module.exports = {
  hash,
  compare
}
