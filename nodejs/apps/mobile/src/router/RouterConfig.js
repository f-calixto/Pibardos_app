import { Routes, Route, Navigate } from 'react-router-native'
import { useSelector } from 'react-redux'

// screens
import AuthScreen from '@Screens/AuthScreen'
import RegisterScreen from '@Screens/RegisterScreen'
import LoginScreen from '@Screens/LoginScreen'
import GroupsScreen from '@Screens/GroupsScreen'

const RouterConfig = () => {
  const isLoggedIn = useSelector(state => state.user.isLoggedIn)

  return (
    <Routes>
      <Route exact path='/'
        element={isLoggedIn
          ? <GroupsScreen />
          : <Navigate to='/auth' />}
      />

      <Route exact path='/auth' element={<AuthScreen />}/>

      <Route exact path='/register' element={<RegisterScreen />} />

      <Route exact path='/login' element={<LoginScreen />} />

      <Route path='*' element={
        <Navigate to='/' replace />}
      />
    </Routes>
  )
}

export default RouterConfig
