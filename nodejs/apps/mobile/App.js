import { NativeBaseProvider } from 'native-base'
// import SplashScreen from './src/containers/SplashScreen'
// import AuthScreen from './src/screens/AuthScreen'
// import LoginScreen from './src/screens/LoginScreen'
// import RegisterScreen from './src/screens/RegisterScreen'
// import HeaderBar from './src/HeaderBar'
import GroupsScreen from './src/screens/GroupsScreen'

export default function App () {
  return (
    <NativeBaseProvider>
      <GroupsScreen />
    </NativeBaseProvider>
  )
}
