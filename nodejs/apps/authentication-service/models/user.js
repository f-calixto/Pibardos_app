const mongoose = require('mongoose')
const { v4: uuidv4 } = require('uuid')
const uniqueValidator = require('mongoose-unique-validator')
const { isEmail, isAlphanumeric } = require('validator')
const passwordMethods = require('../utils/passwordMethods')

const UserSchema = new mongoose.Schema({
  _id: {
    type: String,
    default: uuidv4
  },

  username: {
    type: String,
    required: [true, 'Username is a required field'],
    trim: true,
    lowercase: true,
    minlength: [5, 'Username minimum length is 8 characters'],
    maxlength: [24, 'Username maximum length is 24 characters'],
    unique: [true, 'Username already exists'],
    validate: {
      validator: isAlphanumeric,
      message: 'Username may only have letters and numbers'
    }
  },

  email: {
    type: String,
    required: [true, 'Email is a required field'],
    trim: true,
    lowercase: true,
    unique: [true, 'Email already exists'],
    validate: {
      validator: isEmail,
      message: 'Email is invalid'
    }
  },

  password: {
    type: String,
    required: [true, 'Password is a required field']
  }
}, {
  timestamps: {
    createdAt: 'created_at'
  }
})

UserSchema.pre('save', function (next) {
  const user = this

  // if user is created or if the password is modified, then save the hashed password
  if (this.isModified('password') || this.isNew) {
    passwordMethods.hash(user.password)
      .then(hashedPassword => {
        user.password = hashedPassword
        next()
      })
      .catch(err => next(err))
  } else {
    return next()
  }
})

UserSchema.methods.comparePassword = function (string) {
  const hash = this.password
  return passwordMethods.compare(string, hash)
}

UserSchema.plugin(uniqueValidator)

// remove password and __v fields from returned json
// and rename _id field to id
UserSchema.set('toJSON', {
  transform: (document, returnedObject) => {
    returnedObject.id = document.id
    delete returnedObject._id
    delete returnedObject.password
    delete returnedObject.__v
  }
})

// remove password field from returned user document
UserSchema.set('toObject', {
  transform: (document, returnedObject) => {
    delete returnedObject.password
  }
})

module.exports = mongoose.model('User', UserSchema)
