import App from './App'
import { NativeBaseProvider } from 'native-base'
// import SplashScreen from './src/containers/SplashScreen'
import useCustomTheme from '@Hooks/useCustomTheme'

import { NativeRouter } from 'react-router-native'
import { Provider as ReduxProvider } from 'react-redux'
import store from './src/redux/store'

export default function Main () {
  const theme = useCustomTheme()

  return (
    <ReduxProvider store={store}>
      <NativeBaseProvider theme={theme}>
        <NativeRouter>
          <App />
        </NativeRouter>
      </NativeBaseProvider>
    </ReduxProvider>
  )
}
