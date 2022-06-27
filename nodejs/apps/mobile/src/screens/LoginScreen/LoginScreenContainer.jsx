import { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { Navigate } from 'react-router-native'
import { loginUser } from '@ReduxSlices/slices/user'
import { useToast } from 'native-base'
import ToastAlert from '@Components/ToastAlert'

import LoginScreenView from './LoginScreenView'

const initialValues = {
  email: '',
  password: ''
}

const LoginScreenContainer = () => {
  const dispatch = useDispatch()
  const userState = useSelector(state => state.user)
  const toast = useToast()

  useEffect(() => {
    if (userState.isLoggedIn) {
      toast.show({
        render: () => (
          <ToastAlert
            status='success'
            message='You have successfully signed in!'
          />
        )
      })
    }
  }, [userState.isLoggedIn])

  const onSubmit = async (values, { setSubmitting }) => {
    await dispatch(loginUser({
      email: values.email,
      password: values.password
    }))
    setSubmitting(false)
  }

  if (userState.isLoggedIn) {
    return <Navigate to='/' />
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
