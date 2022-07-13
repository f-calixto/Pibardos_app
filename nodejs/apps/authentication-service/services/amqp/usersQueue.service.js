const { RABBITMQ_URI_CONNECTION } = require('../../config/index')
const amqplib = require('amqplib')

const queue = 'usersQueue'
const options = {
  clientProperties: {
    connection_name: 'authentication-service'
  }
}

const publish = async (payload) => {
  const connection = await amqplib.connect(RABBITMQ_URI_CONNECTION, options)
  const channel = await connection.createChannel()
  await channel.assertQueue(queue, { durable: true })
  channel.sendToQueue(
    queue,
    Buffer.from(JSON.stringify(payload)),
    { persistent: true }
  )

  // close connection
  if (channel) await channel.close()
  await connection.close()
}

module.exports = {
  publish
}
