import { Button, Input, Flex, Text, Pressable, Select, Icon } from 'native-base'
import { Image } from 'react-native'
import theme from '../../theme'
import React, { useState } from 'react'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'

const LoginScreen = () => {
  const [showPswd, setShowPswd] = useState(false)
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [verifyPassword, setVerifyPassword] = useState('')
  const [userName, setUserName] = useState('')
  const [birthday, setBirthday] = useState('')
  const [country, setCountry] = useState('')

  return (
    <Flex flex='1'>
      <Flex justify='center' align='center' mt='30%'>
        <Image
          source={require('../assets/logos/favicon.png') }
          width='50%'
          height= '50%'
          alt='logo'
        />
        <Text mt='3' fontSize='35' fontWeight='bold'>Registrar Cuenta</Text>
      </Flex>
      <Flex mt='10%' justify='space-between' align='center' >
        <Input width='90%' keyboardType='email-address' placeholder='E-mail' onChangeText={email => setEmail(email)}></Input>
        <Input width='90%' type={showPswd ? 'text' : 'password'}
          InputRightElement={<Icon as={<MaterialIcons name={showPswd ? 'visibility' : 'visibility-off'} />}
          size={5} mr='2' color='muted.400' onPress={() => setShowPswd(!showPswd)} />}
          placeholder='Contraseña'
          onChangeText={pswd => setPassword(pswd)}
        />
        <Input width='90%' type={showPswd ? 'text' : 'password'}
          InputRightElement={<Icon as={<MaterialIcons name={showPswd ? 'visibility' : 'visibility-off'} />}
          size={5} mr='2' color='muted.400' onPress={() => setShowPswd(!showPswd)} />}
          placeholder='Repetir contraseña'
          onChangeText={verifyPswd => setVerifyPassword(verifyPswd)}
        />
        <Input width='90%' placeholder='Nombre de usuario' onChangeText={user => setUserName(user)}></Input>
        <Input width='90%' placeholder='Cumpleanos' onChangeText={bday => setBirthday(bday)}></Input>
        <Select minWidth='90%' accessibilityLabel='Pais' placeholder='Pais' onValueChange={(country) => setCountry(country) }>
          <Select.Item label='Argentina' value='AR' />
          <Select.Item label='Uruguay' value='UY' />
          <Select.Item label='Nueva Zelandia' value='NZ' />
          <Select.Item label='Mexico' value='MX' />
          <Select.Item label='Francia' value='FR' />
          <Select.Item label='Colombia' value='CO' />
          <Select.Item label='Inglaterra' value='UK' />
          <Select.Item label='Brasil' value='BR' />
        </Select>
        <Button
          width='70%'
          bg={theme.colors.green}
          height='12%'
          mt='5%'
          onPress={() => console.log({ email, password, verifyPassword, userName, birthday, country })}>
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
