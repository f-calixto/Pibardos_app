import { NativeBaseProvider, StatusBar, Box } from 'native-base'
// import SplashScreen from './src/containers/SplashScreen'
import useCustomTheme from '@Hooks/useCustomTheme'

import { NativeRouter } from 'react-router-native'
import { Provider as ReduxProvider } from 'react-redux'
import store from './src/redux/store'
import RouterConfig from './src/router/RouterConfig'

export default function App () {
  const theme = useCustomTheme()

  return (
    <ReduxProvider store={store}>
      <NativeBaseProvider theme={theme}>
        <NativeRouter>
          <Box flex={1}>
            <StatusBar
              barStyle='dark-content'
              backgroundColor=' #ffffff' // for android
            />
            <RouterConfig />
          </Box>
        </NativeRouter>
      </NativeBaseProvider>
    </ReduxProvider>
  )
}
