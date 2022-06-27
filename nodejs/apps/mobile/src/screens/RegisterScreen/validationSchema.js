import * as yup from 'yup'
import inputErrorMessages from '@Utils/inputErrorMessages'

const validationSchema = yup.object({
  email: yup
    .string()
    .email()
    .required(inputErrorMessages.REQUIRED_FIELD),

  password: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD),

  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password'), null], 'Passwords do not match'),

  username: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD),

  birthdate: yup
    .date()
    .required(inputErrorMessages.REQUIRED_FIELD),

  country: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD)
})

export default validationSchema
