import { Column, Row, Box, Heading, Text, Icon, Skeleton } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'

const DebtsOverview = ({
  title = 'Untitled',
  amount = '0',
  amountColor = 'text.900',
  people = '0',
  pending = '0',
  isLoading
}) => {
  return (
    <Column
      width='full'
      alignItems='center'
      p={5}
      mb={5}
      borderRadius={10}
      shadow='5'
      bgColor='white'
    >
      {isLoading
        ? (
          <>
            <Skeleton rounded='full' w={20} h={5} />
            <Skeleton rounded='full' w={40} h={10} my={5} />
            <Row justifyContent='space-around' w='full'>
              <Skeleton rounded='full' w={24} h={5} />
              <Skeleton rounded='full' w={24} h={5} />
            </Row>
          </>
          )
        : (
        <>
          <Text fontSize='lg' fontWeight='medium' color='text.400'>{title}</Text>
          <Heading fontSize='4xl' color={amountColor}>{`$${amount}.00`}</Heading>
          <Box w='full' h='1px' my={3} bgColor='gray.200' />
          <Row w='full' justifyContent='space-around'>
            <Row alignItems='center'>
              <Icon
                as={<Ionicons name='person' />}
                size={5}
                color='amber.400'
                mr={2}
              />
              <Text fontSize='md' fontWeight='medium' color='text.600'>{`${people} people`}</Text>
            </Row>
            <Row alignItems='center'>
              <Icon
                as={<Ionicons name='time' />}
                size={5}
                color='amber.400'
                mr={2}
              />
              <Text fontSize='md' fontWeight='medium' color='text.600'>{`${pending} pending`}</Text>
            </Row>
          </Row>
        </>
          )}
    </Column>
  )
}

export default DebtsOverview
