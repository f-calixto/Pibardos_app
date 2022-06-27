import { Flex, Text, ScrollView, VStack, KeyboardAvoidingView } from 'native-base'
import { Image, Platform } from 'react-native'
import { Formik } from 'formik'
import Constants from 'expo-constants'
import validationSchema from './validationSchema'
import RegisterForm from './components/RegisterForm'
import theme from '@Theme'

const RegisterScreenView = ({ initialValues, onSubmit, userState }) => {
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
              Create account
            </Text>
          </Flex>
          <Formik
            initialValues={initialValues}
            onSubmit={onSubmit}
            validationSchema={validationSchema}
          >
            {({ handleSubmit, isSubmitting, setFieldError }) => (
              <RegisterForm
                onSubmit={handleSubmit}
                isSubmitting={isSubmitting}
                setFieldError={setFieldError}
                errors={userState.errors}
              />
            )}
          </Formik>
        </VStack>
      </ScrollView>
    </KeyboardAvoidingView>
  )
}

export default RegisterScreenView
