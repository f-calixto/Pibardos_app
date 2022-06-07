class ServiceError extends Error {
  constructor (name, path, message, stack = '') {
    super(message)
    this.name = name
    this.path = path

    if (stack) {
      this.stack = stack
    } else {
      Error.captureStackTrace(this, this.constructor)
    }

    this.errors = [{ field: path, userMessage: message }]
  }
}

module.exports = ServiceError
