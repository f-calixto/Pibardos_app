import { useState, useEffect } from 'react'
import { Box, Button, Icon } from 'native-base'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import FormikTextInput from '@Components/FormikTextInput'
import theme from '@Theme'

const LoginForm = ({ onSubmit, isSubmitting, setFieldError, fetchErrors }) => {
  const [showPassword, setShowPassword] = useState(false)

  // TODO: Create a custom hook with this logic. Do the same with RegisterForm component.
  useEffect(() => {
    if (fetchErrors && fetchErrors.length > 0) {
      fetchErrors.forEach(error => setFieldError(error.field, error.userMessage))
    }
  }, [fetchErrors])

  return (
    <Box>
      <FormikTextInput
        name='email'
        placeholder='E-mail'
        autoCorrect={false}
        keyboardType='email-address'
      />

      <FormikTextInput
        name='password'
        placeholder='Password'
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

      <Button
        bgColor={theme.colors.blue}
        mt={theme.fontSizes.large}
        onPress={onSubmit}
        isLoading={isSubmitting}
        isLoadingText='Logging in to account'
        _loading={{
          _text: {
            color: theme.colors.secondary
          }
        }}
        _spinner={{
          color: theme.colors.secondary
        }}
      >
        Sign In
      </Button>
    </Box>
  )
}

export default LoginForm
