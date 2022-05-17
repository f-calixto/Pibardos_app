const mongoose = require('mongoose')

const connect = (uri, options) => {
  const parsedOptions = options || {
    useNewUrlParser: true,
    useUnifiedTopology: true
  }

  return mongoose.connect(uri, parsedOptions)
    .then(() => console.log('Succesful database connection'))
    .catch(error => console.error('An error occurred when trying to establish a connection with database:', error))
}

const closeConnection = () => (
  mongoose.connection.close()
    .then(() => console.log('Database connection closed'))
    .catch(error => console.error('An error occured whwn trying to close the database connection:', error))
)

module.exports = {
  connect,
  closeConnection
}
