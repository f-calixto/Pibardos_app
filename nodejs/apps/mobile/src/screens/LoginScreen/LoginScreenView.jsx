import { Platform, Image } from 'react-native'
import { Flex, Text, ScrollView, VStack, KeyboardAvoidingView } from 'native-base'
import Constants from 'expo-constants'
import { Formik } from 'formik'
import LoginForm from './components/LoginForm'
import ButtonLink from '@Components/ButtonLink'
import validationSchema from './validationSchema'
import theme from '@Theme'

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
              Sign In
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
            {/* TODO: Add recovery password feature */}
            {/* <Button variant='link'>Olvidé mi contraseña</Button> */}
            <ButtonLink
              to='/register'
              variant='link'
            >
              I am not registered
            </ButtonLink>
          </VStack>
        </VStack>
      </ScrollView>
    </KeyboardAvoidingView>
  )
}

export default LoginScreenView
