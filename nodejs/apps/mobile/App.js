import { NativeBaseProvider, StatusBar, Box } from 'native-base'
// import SplashScreen from './src/containers/SplashScreen'
import theme from './theme'

import { NativeRouter } from 'react-router-native'
import { Provider as ReduxProvider } from 'react-redux'
import store from './src/redux/store'
import RouterConfig from './src/router/RouterConfig'

export default function App () {
  return (
    <ReduxProvider store={store}>
      <NativeBaseProvider>
        <NativeRouter>
          <Box flex={1}>
            <StatusBar
              barStyle='dark-content'
              backgroundColor={theme.colors.white} // for android
            />
            <RouterConfig />
          </Box>
        </NativeRouter>
      </NativeBaseProvider>
    </ReduxProvider>
  )
}
