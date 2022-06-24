import { Platform, Image } from 'react-native'
import { Flex, Text, ScrollView, Button, VStack, KeyboardAvoidingView } from 'native-base'
import Constants from 'expo-constants'
import { Formik } from 'formik'
import theme from '@Theme'
import LoginForm from './components/LoginForm'
import validationSchema from './validationSchema'

const LoginScreenView = ({ initialValues, onSubmit, fetchErrors }) => {
  return (
    <KeyboardAvoidingView
      flex={1}
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
    >
      <ScrollView contentContainerStyle={{ flexGrow: 1, justifyContent: 'center' }}>
        <VStack
          mt={Constants.statusBarHeight}
          px={theme.fontSizes.medium}
          pb={theme.fontSizes.large}
        >
          <Flex justify='center' align='center' pb={theme.fontSizes.large}>
            <Image
              source={require('../../assets/logos/favicon.png')}
              width='50%'
              height='50%'
              alt='logo'
            />
            <Text mt='3' fontSize='35' fontWeight='bold'>
              Iniciar Sesion
            </Text>
          </Flex>
          <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
            onSubmit={onSubmit}
          >
            {({ handleSubmit, isSubmitting, setFieldError }) => (
              <LoginForm
                onSubmit={handleSubmit}
                isSubmitting={isSubmitting}
                setFieldError={setFieldError}
                fetchErrors={fetchErrors}
              />
            )}
          </Formik>
          <VStack mt={theme.fontSizes.small}>
            <Button variant='link'>Olvidé mi contraseña</Button>
            <Button variant='link'>No estoy registrado</Button>
          </VStack>
        </VStack>
      </ScrollView>
    </KeyboardAvoidingView>
  )
}

export default LoginScreenView
