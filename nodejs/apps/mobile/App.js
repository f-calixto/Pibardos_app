import { NativeBaseProvider, StatusBar, Box } from 'native-base'
// import SplashScreen from './src/containers/SplashScreen'
// import AuthScreen from './src/screens/AuthScreen'
// import LoginScreen from './src/screens/LoginScreen'
import RegisterScreen from './src/screens/RegisterScreen/RegisterScreen'
import theme from './theme'
// import HeaderBar from './src/HeaderBar'
// import GroupsScreen from './src/screens/GroupsScreen'

export default function App () {
  return (
    <NativeBaseProvider>
      <Box flex={1}>
        <StatusBar
          barStyle='dark-content'
          backgroundColor={theme.colors.white} // for android
        />

        <RegisterScreen />
        {/* <GroupsScreen /> */}
      </Box>
    </NativeBaseProvider>
  )
}
