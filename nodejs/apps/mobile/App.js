// 1. import `NativeBaseProvider` component
import { NativeBaseProvider } from 'native-base'
import SplashScreen from './src/containers/SplashScreen'

export default function App () {
  // 2. Use at the root of your app
  return (
    <NativeBaseProvider>
        <SplashScreen />
      {/* <Box flex={1} bg='#fff' alignItems='center' justifyContent='center'>
        <Text>Open up App.js to start working on your app!</Text>
      </Box> */}
    </NativeBaseProvider>
  )
}
