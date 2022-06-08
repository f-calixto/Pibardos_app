import { Button, Input, Flex, Text, Pressable, Select } from 'native-base'
import { Image } from 'react-native'
import theme from '../../theme'

const LoginScreen = () => {
  return (

    <Flex flex='1'>

    <Flex justify='center' align='center' mt='30%'>
        <Image
          source={require('../assets/logos/favicon.png') }
          width='50%'
          height= '50%'
          alt='logo'>

        </Image>
      <Text mt='3' fontSize='35' fontWeight='bold'>Registrar Cuenta</Text>
    </Flex>
    <Flex mt='10%' justify='space-between' align='center' >
        <Input width='90%' placeholder='E-mail'></Input>
        <Input width='90%' placeholder='Contraseña'></Input>
        <Input width='90%' placeholder='Repetir contraseña'></Input>
        <Input width='90%' placeholder='Nombre de usuario'></Input>
        <Input width='90%' placeholder='Cumpleanos'></Input>
        <Select minWidth='90%' accessibilityLabel='Pais' placeholder='Pais' onValueChange={(value) => console.log('Country changed: ' + value)}>
          <Select.Item label='Argentina' value='AR' />
          <Select.Item label='Uruguay' value='UY' />
          <Select.Item label='Nueva Zelandia' value='NZ' />
          <Select.Item label='Mexico' value='MX' />
          <Select.Item label='Francia' value='FR' />
          <Select.Item label='Colombia' value='CO' />
          <Select.Item label='Inglaterra' value='UK' />
          <Select.Item label='Brasil' value='BR' />
        </Select>
        <Button width='70%' bg={theme.colors.green} height='12%' mt='5%' onPress={() => console.log('register press')}>
        Registrarse
      </Button>
      <Pressable onPress={() => console.log('Ya tengo una cuenta')}>
          <Text>Ya tengo una cuenta</Text>
      </Pressable>
    </Flex>
    </Flex>
  )
}
export default LoginScreen
