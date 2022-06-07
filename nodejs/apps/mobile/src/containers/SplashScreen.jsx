import { Box, Text, Heading } from 'native-base'
import { StatusBar } from 'expo-status-bar'
import SVGImg from '../assets/icon.svg'

const SplashScreen = () => {
  return (
    <Box
      h='100%'
      w='100%'
      p={10}
      flex={1}
      flexDir='column'
      alignItems='center'
      justifyContent='center'
      // bgColor='#ff6575'
    >
      <StatusBar style='dark' />
      <Box mt='auto' alignItems='center'>
        <SVGImg width={150} height={150} />
        <Heading size='xl' color='#ff6575' mt={3}>PibardosApp</Heading>
      </Box>
      <Box mt='auto'>
        <Text color='#ff6575'>Made with ♥ by Franks, Ivanchu y Maxen</Text>
      </Box>
    </Box>
  )
}

export default SplashScreen
