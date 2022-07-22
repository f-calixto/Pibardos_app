import { Row, Column, Text, Badge, Avatar, Icon } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'

const StatusBadge = ({ status }) => {
  if (status === 'pending') return <Badge colorScheme='yellow' rounded='full' leftIcon={<Icon as={<Ionicons name='time-outline' />} />}>Pending</Badge>
  if (status === 'accepted') return <Badge colorScheme='darkBlue' rounded='full' leftIcon={<Icon as={<Ionicons name='checkmark-done-outline' />} />}>Accepted</Badge>
  if (status === 'rejected') return <Badge colorScheme='red' rounded='full' leftIcon={<Icon as={<Ionicons name='close-outline' />} />}>Rejected</Badge>
  if (status === 'paid') return <Badge colorScheme='green' rounded='full' leftIcon={<Icon as={<Ionicons name='checkmark-outline' />} />}>Paid</Badge>
}

const UserDebtItem = ({
  username = 'username',
  amount = '250',
  status,
  createdAt
}) => {
  return (
    <Row justifyContent='space-between' my={2}>
      <Row>
        <Avatar bg='gray.400' size='md' mr={3} source={{
          uri: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80'
        }}>
            {username[0].toUpperCase()}
        </Avatar>
        <Column>
          <Text fontSize='lg' fontWeight='medium' color='text.900'>{username}</Text>
          <Text fontSize='sm' fontWeight='medium' color='text.400'>2 hours ago</Text>
        </Column>
      </Row>
      <Column justifyContent='center' alignItems='flex-end'>
        <Text fontSize='xl' fontWeight='bold' color='text.900'>{`$${amount}.00`}</Text>
        {status && <StatusBadge status={status} />}
      </Column>
    </Row>
  )
}

export default UserDebtItem
