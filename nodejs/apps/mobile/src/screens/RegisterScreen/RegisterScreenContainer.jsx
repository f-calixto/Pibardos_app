import { useSelector, useDispatch } from 'react-redux'
import { registerUser } from '../../redux/slices/user'
import RegisterScreenView from './RegisterScreenView'

const initialValues = {
  email: '',
  password: '',
  confirmPassword: '',
  username: '',
  birthdate: new Date('12/01/2000'),
  country: ''
}

const RegisterScreenContainer = () => {
  const dispatch = useDispatch()
  const userState = useSelector(state => state.user)

  const onSubmit = async (values, actions) => {
    await dispatch(registerUser(values))
    actions.setSubmitting(false)
  }

  return (
    <RegisterScreenView
      initialValues={initialValues}
      onSubmit={onSubmit}
      userState={userState}
    />
  )
}

export default RegisterScreenContainer
