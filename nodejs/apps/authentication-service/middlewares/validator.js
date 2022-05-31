module.exports = schema => (req, res, next) => {
  const { error } = schema.validate(req.body, {
    abortEarly: false,
    errors: {
      wrap: { label: '' }
    }
  })

  if (error) {
    return next(error)
  }

  next()
}
