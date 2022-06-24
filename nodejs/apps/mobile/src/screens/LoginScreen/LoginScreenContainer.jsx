import { useSelector, useDispatch } from 'react-redux'
import { loginUser } from '@ReduxSlices/slices/user'

import LoginScreenView from './LoginScreenView'

const initialValues = {
  email: '',
  password: ''
}

const LoginScreenContainer = () => {
  const dispatch = useDispatch()
  const userState = useSelector(state => state.user)

  const onSubmit = async (values, { setSubmitting }) => {
    await dispatch(loginUser({
      email: values.email,
      password: values.password
    }))
    setSubmitting(false)
  }

  return (
    <LoginScreenView
      onSubmit={onSubmit}
      initialValues={initialValues}
      fetchErrors={userState.errors}
    />
  )
}
export default LoginScreenContainer
