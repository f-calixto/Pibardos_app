import { NativeBaseProvider } from 'native-base'
// import SplashScreen from './src/containers/SplashScreen'
import AuthScreen from './src/screens/AuthScreen'

export default function App () {
  return (
    <NativeBaseProvider>
      <AuthScreen />
    </NativeBaseProvider>
  )
}
