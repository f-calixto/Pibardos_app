import { Alert, VStack, HStack, Text } from 'native-base'

const ToastAlert = ({ status, message }) => {
  return (
    <Alert w='100%' variant='subtle' colorScheme={status} status={status}>
      <VStack space={2} flexShrink={1} w='100%'>
        <HStack flexShrink={1} space={2} alignItems='center' justifyContent='space-between'>
          <HStack space={2} flexShrink={1} alignItems='center'>
            <Alert.Icon />
            <Text color='coolGray.800'>
              {message}
            </Text>
          </HStack>
        </HStack>
      </VStack>
    </Alert>
  )
}

export default ToastAlert
