import { useState, useEffect } from 'react'
import { Button, Icon, Box } from 'native-base'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import { countries } from 'countries-list'
import theme from '../../../../theme'

// formik input components
import FormikTextInput from '../../../components/FormikTextInput'
import FormikDatepicker from '../../../components/FormikDatepicker'
import FormikSelectInput from '../../../components/FormikSelectInput'
import CountryFilterInput from './CountryFilterInput'

const countriesList = Object.entries(countries).map(([key, value]) => ({
  label: `${value.emoji} ${value.name}`,
  value: key
}))

const RegisterForm = ({ onSubmit, isSubmitting, setFieldError, errors }) => {
  const [showPassword, setShowPassword] = useState(false)
  const [filteredCountriesList, setFilteredCountriesList] = useState([])

  useEffect(() => {
    if (errors && errors.length > 0) {
      errors.forEach(error => {
        setFieldError(error.field, error.userMessage)
      })
    }
  }, [errors])

  const onFilterChange = newValue => {
    if (newValue.length === 0) return setFilteredCountriesList([])

    const filtered = countriesList.filter(country => (
      country.label.toLowerCase().includes(newValue.toLowerCase())
    ))
    setFilteredCountriesList(filtered)
  }

  return (
    <Box>
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

      <Button
        bgColor={theme.colors.green}
        mt={theme.fontSizes.large}
        onPress={onSubmit}
        isLoading={isSubmitting}
        isLoadingText='Registrando cuenta'
        _loading={{
          _text: {
            color: theme.colors.primary
          }
        }}
        _spinner={{
          color: theme.colors.primary
        }}
      >
        Registrarse
      </Button>
      <Button variant='link' mt={theme.fontSizes.small}>Ya tengo una cuenta</Button>
    </Box>
  )
}

export default RegisterForm
