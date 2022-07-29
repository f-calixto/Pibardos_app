import { useState, useEffect } from 'react'
import { StatusBar, Box } from 'native-base'
import { useDispatch } from 'react-redux'
import RouterConfig from './src/router/RouterConfig'
import { secureStoreService } from '@Services/secureStore.service'
import { setUser } from './src/redux/slices/user'

export default function App () {
  const [isReady, setIsReady] = useState(false)
  const dispatch = useDispatch()

  useEffect(() => {
    const loadSavedUser = async () => {
      const savedUser = await secureStoreService.getUser()
      await dispatch(setUser(savedUser))
      setIsReady(true)
    }

    loadSavedUser()
  }, [])

  // TODO: add splash screen
  if (!isReady) return <Box />

  return (
    <Box flex={1}>
      <StatusBar
        barStyle='dark-content'
        backgroundColor=' #ffffff' // for android
      />
      <RouterConfig />
    </Box>
  )
}
