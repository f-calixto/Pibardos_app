const mongoose = require('mongoose')

const handleMongooseValidationError = (err, res) => {
  const errors = Object.values(err.errors).map(error => {
    const errorMessage = error.properties.message

    return {
      field: error.path,
      userMessage: errorMessage
    }
  })

  return res.status(400).json({ errors })
}

const handleJoiValidationError = (err, res) => {
  const errors = err.details.map(error => {
    return {
      field: error.path[0],
      userMessage: error.message
    }
  })

  return res.status(400).json({ errors })
}

module.exports = (err, req, res, next) => {
  console.log(err)

  if (err instanceof mongoose.Error.ValidationError) return handleMongooseValidationError(err, res)
  if (err.isJoi) return handleJoiValidationError(err, res)

  return res.status(500).json({ error: 'An unknown error occurred' })
}
