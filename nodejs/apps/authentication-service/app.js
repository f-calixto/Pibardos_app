const express = require('express')
const app = express()
const logger = require('morgan')

app.use(express.json())
app.use(logger('dev'))

module.exports = app