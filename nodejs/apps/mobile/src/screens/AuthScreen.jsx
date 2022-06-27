import { Image } from 'react-native'
import { Flex, Text } from 'native-base'
import ButtonLink from '@Components/ButtonLink'
import theme from '@Theme'

const AuthScreen = () => {
  return (
    <Flex flex='1'>
      <Flex direction='row' justify='center' align='center' pt='50%'>
        <Image
          source={require('../assets/logos/favicon.png')}
          width='20%'
          height='20%'
          alt='logo'
        ></Image>

        <Text pl='2' fontSize='40' fontWeight='bold'>
          Pibardos App
        </Text>
      </Flex>
      <Flex flex='1' justify='center' mt='-30%' px={5}>
        <ButtonLink
          to='/login'
          bg={theme.colors.blue}
          h={12}
        >
          Sign In
        </ButtonLink>

        <ButtonLink
          to='/register'
          mt={3}
          h={12}
          bg={theme.colors.green}
        >
          Create an account
        </ButtonLink>
      </Flex>
      <Flex align='center' pb='5%'>
        <Text>Made with 🖤 by Frank, Mazen e Ivanchu </Text>
      </Flex>
    </Flex>
  )
}

export default AuthScreen
