const mongoose = require('mongoose')
const { v4: uuidv4 } = require('uuid')
const uniqueValidator = require('mongoose-unique-validator')
const { isEmail } = require('validator')

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
    unique: [true, 'Username already exists']
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
    required: [true, 'Password is a required field'],
  }
}, {
  timestamps: {
    createdAt: 'created_at'
  }
})

UserSchema.plugin(uniqueValidator)

UserSchema.set('toJSON', {
  transform: (document, returnedObject) => {
    returnedObject.id = document.id
    delete returnedObject._id
    delete returnedObject.password
    delete returnedObject.__v
  }
})

module.exports = mongoose.model('User', UserSchema)