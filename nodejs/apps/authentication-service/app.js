const express = require('express')
const logger = require('morgan')
const errorMiddleware = require('./middlewares/error')

// init app
const app = express()

// Middlewares
app.use(express.json())
app.use(logger('dev'))

// Routes
// app.use('/token', require('./routes/token.routes')) // token routes
app.use('/', require('./routes/users.routes')) // users routes

// Errors middleware
app.use(errorMiddleware)

module.exports = app
