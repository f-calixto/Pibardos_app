import { useState } from 'react'
import { Button, Flex, Text, Icon, ScrollView, VStack } from 'native-base'
import { Image } from 'react-native'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import { Formik } from 'formik'
import * as yup from 'yup'
import FormikTextInput from '../components/FormikTextInput'
import FormikDatepicker from '../components/FormikDatepicker'
import FormikSelectInput from '../components/FormikSelectInput'
import Constants from 'expo-constants'
import { countries } from 'countries-list'
import theme from '../../theme'

const initialValues = {
  email: '',
  password: '',
  repeatPassword: '',
  username: '',
  birthdate: new Date('12/01/2000'),
  country: ''
}

const validationSchema = yup.object({
  username: yup
    .string()
    .required('Username is required'),

  password: yup
    .string()
    .required('Password is required'),

  birthdate: yup
    .date()
    .required('Date is required')
})

const countriesList = Object.entries(countries).map(([key, value]) => ({
  label: `${value.emoji} ${value.name}`,
  value: key
}))

const Form = ({ onSubmit }) => {
  const [showPassword, setShowPassword] = useState(false)

  return (
    <VStack>
      <FormikTextInput
        name='email'
        placeholder='Correo electrónico'
        autoCorrect={false}
        keyboardType='email-address'
      />

      <FormikTextInput
        name='password'
        placeholder='Contraseña'
        autoCorrect={false}
        InputRightElement={
          <Icon
            as={
              <MaterialIcons
                name={showPassword ? 'visibility' : 'visibility-off'}
              />
            }
            size={5}
            mr='2'
            color='muted.400'
            onPress={() => setShowPassword(!showPassword)}
          />
        }
        secureTextEntry={showPassword}
      />

      <FormikTextInput
        name='repeatPassword'
        placeholder='Repetir contraseña'
        autoCorrect={false}
        InputRightElement={
          <Icon
            as={
              <MaterialIcons
                name={showPassword ? 'visibility' : 'visibility-off'}
              />
            }
            size={5}
            mr='2'
            color='muted.400'
            onPress={() => setShowPassword(!showPassword)}
          />
        }
        secureTextEntry={showPassword}
      />

      <FormikTextInput
        name='username'
        placeholder='Nombre de usuario'
        autoCorrect={false}
      />

      <FormikDatepicker name='birthdate' placeholder='Fecha de nacimiento'/>

      <FormikSelectInput
        name='country'
        placeholder='País'
        items={countriesList}
      />

      <Button bgColor={theme.colors.green} mt={theme.fontSizes.large} onPress={onSubmit}>Sign In</Button>
      <Button variant='link' mt={theme.fontSizes.small}>Ya tengo una cuenta</Button>
    </VStack>
  )
}

const LoginScreen = () => {
  const onSubmit = val => console.log(val)
  return (
    <ScrollView>
      <VStack
        mt={Constants.statusBarHeight}
        px={theme.fontSizes.medium}
        pb={theme.fontSizes.large}
      >
        <Flex justify='center' align='center' pb={theme.fontSizes.large}>
          <Image
            source={require('../assets/logos/favicon.png')}
            width='100%'
            height='100%'
            alt='logo'
          />
          <Text mt='3' fontSize='35' fontWeight='bold'>
            Registrar Cuenta
          </Text>
        </Flex>
        <Formik
          initialValues={initialValues}
          onSubmit={onSubmit}
          validationSchema={validationSchema}
        >
          {({ handleSubmit }) => <Form onSubmit={handleSubmit} />}
        </Formik>
      </VStack>
    </ScrollView>
  )
}

export default LoginScreen
