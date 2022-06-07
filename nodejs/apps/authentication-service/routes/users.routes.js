const express = require('express')
const router = express.Router()
const usersController = require('../controllers/users')
const validator = require('../middlewares/validator')
const userValidation = require('../validation/user.validation')

// TODO: Add change password route
// TODO: add change email route

router.post('/register', validator(userValidation.createUser), usersController.register)
router.post('/login', validator(userValidation.authenticateUser), usersController.authenticate)

module.exports = router
