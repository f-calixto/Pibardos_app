const jwt = require('jsonwebtoken')
const { ACCESS_TOKEN_SECRET, ACCESS_TOKEN_EXP_TIME } = require('../config')

class TokensService {
  constructor () {
    this.jwt = jwt
    this.accessTokenSecret = ACCESS_TOKEN_SECRET
    this.accessTokenOptions = {
      expiresIn: ACCESS_TOKEN_EXP_TIME
    }
  }

  signAccessToken ({ userId, username, email }) {
    return this.jwt.sign({ userId, username, email }, this.accessTokenSecret, this.accessTokenOptions)
  }

  verifyAccessToken (token) {
    return this.jwt.verify(token, this.accessTokenSecret)
  }
}

module.exports = TokensService
