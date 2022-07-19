import { Routes, Route, Navigate } from 'react-router-native'
import { useSelector } from 'react-redux'

// screens
import AuthScreen from '@Screens/AuthScreen'
import RegisterScreen from '@Screens/RegisterScreen'
import LoginScreen from '@Screens/LoginScreen'
import GroupScreen from '@Screens/GroupScreen'
import ManageGroupsScreen from '@Screens/ManageGroupsScreen'
import DebtsScreen from '@Screens/DebtsScreen'

const RouterConfig = () => {
  const isLoggedIn = useSelector(state => state.user.isLoggedIn)

  return (
    <Routes>
      {/* <Route exact path='/'
        element={isLoggedIn
          ? <GroupScreen />
          : <Navigate to='/auth' />}
      /> */}

      <Route exact path='/'
        element={isLoggedIn
          ? <Navigate to='/group' />
          : <Navigate to='/debts' />}
      />

      {/* Authentication */}
      <Route exact path='/auth' element={<AuthScreen />}/>
      <Route exact path='/register' element={<RegisterScreen />} />
      <Route exact path='/login' element={<LoginScreen />} />

      {/* Groups */}
      <Route exact path='/group' element={<GroupScreen />} />
      <Route exact path='/manage-groups' element={<ManageGroupsScreen />} />

      <Route exact path='/debts' element={<DebtsScreen />} />

      <Route path='*' element={
        <Navigate to='/' replace />}
      />
    </Routes>
  )
}

export default RouterConfig
