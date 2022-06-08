import { Button, Input, Flex, Text, Pressable } from 'native-base'
import { Image } from 'react-native'
import theme from '../../theme'

const LoginScreen = () => {
  return (
    <Flex flex='1'>

    <Flex justify='center' align='center' mt='50%'>
        <Image
          source={require('../assets/logos/favicon.png') }
          width='50%'
          height= '50%'
          alt='logo'>

        </Image>
      <Text mt='3' fontSize='35' fontWeight='bold'>Iniciar Sesion</Text>
    </Flex>
    <Flex mt='10%' justify='space-between' align='center' >
        <Input width='90%' placeholder='E-mail o nombre de usuario'></Input>
        <Input width='90%' placeholder='Contraseña'></Input>
        <Button width='70%' bg={theme.colors.blue} height='20%' onPress={() => console.log('login press')}>
        Iniciar Sesion
      </Button>
      <Pressable onPress={() => console.log('Olvide contrasena')}>
          <Text>Olvide mi contraseña</Text>
      </Pressable>
      <Pressable onPress={() => console.log('No estoy registrado')}>
          <Text>No estoy registrado</Text>
      </Pressable>
    </Flex>
    </Flex>
  )
}
export default LoginScreen
