import { useEffect } from 'react'
import { useToast } from 'native-base'
import { useSelector, useDispatch } from 'react-redux'
import RegisterScreenView from './RegisterScreenView'
import { registerUser } from '../../redux/slices/user'
import { Navigate } from 'react-router-native'
import ToastAlert from '@Components/ToastAlert'

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
  const toast = useToast()

  useEffect(() => {
    if (userState.isLoggedIn) {
      toast.show({
        render: () => <ToastAlert
          status='success'
          message='Account successfully created!'
        />
      })
    }
  }, [userState.isLoggedIn])

  const onSubmit = async (values, actions) => {
    await dispatch(registerUser(values))
    actions.setSubmitting(false)
  }

  if (userState.isLoggedIn) {
    return <Navigate to='/' />
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
