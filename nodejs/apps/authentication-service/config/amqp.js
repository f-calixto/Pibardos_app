const { RABBITMQ_URI_CONNECTION } = require('../config/index')
const amqplib = require('amqplib')

const options = {
  clientProperties: {
    connection_name: 'users-service'
  }
}

let channel = null

const connect = async () => {
  amqplib.connect(RABBITMQ_URI_CONNECTION, options, (err, conn) => {
    if (err) throw err

    conn.createChannel((err, connChannel) => {
      if (err) throw err

      connChannel.assertQueue('usersQueue', { durable: true })

      channel = connChannel
    })
  })
}

module.exports = {
  connect,
  channel
}
