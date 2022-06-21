import { useState } from 'react'
import { Button, Flex, Text, Icon, ScrollView, VStack, KeyboardAvoidingView } from 'native-base'
import { Image, Platform } from 'react-native'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import { Formik } from 'formik'
import * as yup from 'yup'
import FormikTextInput from '../../components/FormikTextInput'
import FormikDatepicker from '../../components/FormikDatepicker'
import FormikSelectInput from '../../components/FormikSelectInput'
import CountryFilterInput from './components/CountryFilterInput'
import Constants from 'expo-constants'
import { countries } from 'countries-list'
import inputErrorMessages from '../../utils/inputErrorMessages'
import theme from '../../../theme'

const initialValues = {
  email: '',
  password: '',
  confirmPassword: '',
  username: '',
  birthdate: new Date('12/01/2000'),
  country: ''
}

const validationSchema = yup.object({
  email: yup
    .string()
    .email()
    .required(inputErrorMessages.REQUIRED_FIELD),

  password: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD),

  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password'), null], 'Las constraseñas no coinciden'),

  username: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD),

  birthdate: yup
    .date()
    .required(inputErrorMessages.REQUIRED_FIELD),

  country: yup
    .string()
    .required(inputErrorMessages.REQUIRED_FIELD)
})

const countriesList = Object.entries(countries).map(([key, value]) => ({
  label: `${value.emoji} ${value.name}`,
  value: key
}))

const Form = ({ onSubmit }) => {
  const [showPassword, setShowPassword] = useState(false)
  const [filteredCountriesList, setFilteredCountriesList] = useState([])

  const onFilterChange = newValue => {
    if (newValue.length === 0) return setFilteredCountriesList([])

    const filtered = countriesList.filter(country => (
      country.label.toLowerCase().includes(newValue.toLowerCase())
    ))
    setFilteredCountriesList(filtered)
  }

  return (
    <>
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
                name={showPassword ? 'visibility-off' : 'visibility'}
              />
            }
            size={5}
            mr='2'
            color='muted.400'
            onPress={() => setShowPassword(!showPassword)}
          />
        }
        secureTextEntry={!showPassword}
      />

      <FormikTextInput
        name='confirmPassword'
        placeholder='Confirmar contraseña'
        autoCorrect={false}
        InputRightElement={
          <Icon
            as={
              <MaterialIcons
                name={showPassword ? 'visibility-off' : 'visibility'}
              />
            }
            size={5}
            mr='2'
            color='muted.400'
            onPress={() => setShowPassword(!showPassword)}
          />
        }
        secureTextEntry={!showPassword}
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
        items={filteredCountriesList}
        _actionSheetBody={{
          minH: '100%',
          ListHeaderComponent: (
            <CountryFilterInput onFilterChange={onFilterChange} />
          )
        }}
      />

      <Button bgColor={theme.colors.green} mt={theme.fontSizes.large} onPress={onSubmit}>Sign In</Button>
      <Button variant='link' mt={theme.fontSizes.small}>Ya tengo una cuenta</Button>
    </>
  )
}

const LoginScreen = () => {
  const onSubmit = val => console.log(val)
  return (
    <KeyboardAvoidingView
      h={{
        base: '800px',
        lg: 'auto'
      }}
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
    >
      <ScrollView>
        <VStack
          mt={Constants.statusBarHeight}
          px={theme.fontSizes.medium}
          pb={theme.fontSizes.large}
        >
          <Flex justify='center' align='center' pb={theme.fontSizes.large}>
            <Image
              source={require('../../assets/logos/favicon.png')}
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
    </KeyboardAvoidingView>
  )
}

export default LoginScreen
