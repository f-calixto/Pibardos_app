import * as yup from 'yup'
import inputErrorMessages from '@Utils/inputErrorMessages'

const validationSchema = yup.object({
  email: yup
    .string()
    .email()
    .required(inputErrorMessages.REQUIRED_FIELD),

  password: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD)
})

export default validationSchema
