import { Actionsheet, Box, Text, Icon } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'
import ActionsheetItem from './ActionsheetItem'

const GroupsActionsheet = ({ isOpen, onClose }) => {
  return (
    <Actionsheet isOpen={isOpen} onClose={onClose}>
      <Actionsheet.Content>
        <Box w='100%' h={60} px={4} justifyContent='center'>
          <Text fontSize='16' color='gray.500' _dark={{
            color: 'gray.300'
          }}>
            My groups
          </Text>
        </Box>
        <ActionsheetItem
          name='Los Inmigrantes'
          imageUri='https://images.unsplash.com/photo-1543807535-eceef0bc6599?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80'
        />
        <ActionsheetItem
          name='Los Pibardos'
          imageUri='https://images.unsplash.com/photo-1490578474895-699cd4e2cf59?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=871&q=80'
        />
        <Actionsheet.Item
          leftIcon={<Icon as={Ionicons} name='cog-outline' />}
          _icon={{
            size: 6,
            color: 'black'
          }}
        >
          Manage groups
        </Actionsheet.Item>
      </Actionsheet.Content>
    </Actionsheet>
  )
}

export default GroupsActionsheet
