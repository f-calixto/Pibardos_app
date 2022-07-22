import { Column, Row, Text, Button, Icon, Skeleton } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'

const DebtsSummary = ({
  title = 'Untitled',
  buttonTo,
  buttonText = 'Untitled',
  children,
  isLoading
}) => {
  return (
    <Column
      w='full'
      p={5}
      mb={5}
      borderRadius={10}
      shadow='5'
      bgColor='white'
    >
      {isLoading
        ? (
          <>
            <Row w='full' justifyContent='space-between'>
              <Skeleton w='100px' h={5} rounded='full' />
              <Skeleton w='60px' h={5} rounded='full' />
            </Row>
            {Array(3).fill(0).map((i, idx) => (
              <Row key={idx} alignItems='center' mt={5}>
                <Skeleton h={8} size={8} rounded='full' />
                <Skeleton h={8} flex={1} mx={3} rounded='full' />
                <Skeleton h={8} flex={2} rounded='full' />
              </Row>
            ))}
          </>
          )
        : (
          <>
            <Row justifyContent='space-between' alignItems='center' mb={5}>
              <Text fontSize='md' fontWeight='medium' color='text.400'>{title}</Text>
              <Button
                py='3px'
                rounded='full'
                rightIcon={<Icon as={<Ionicons name='chevron-forward-outline' />} size='sm' />}
              >
                {buttonText}
              </Button>
            </Row>
            {children}
          </>
          )}
    </Column>
  )
}

export default DebtsSummary
