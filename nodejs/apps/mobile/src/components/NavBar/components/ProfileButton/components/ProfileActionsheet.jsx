import { Actionsheet, Row, Text, Icon, Avatar } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'

const ProfileActionsheet = ({ isOpen, onClose }) => {
  return (
    <Actionsheet isOpen={isOpen} onClose={onClose}>
      <Actionsheet.Content>
        <Row alignItems='center' h={70} w='full' px={4}>
          <Avatar
            source={{ uri: 'https://images.unsplash.com/photo-1499996860823-5214fcc65f8f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=466&q=80' }}
            mr={3}
            size={12}
          />
          <Text fontWeight='semibold' fontSize={16}>olimar350</Text>
        </Row>
        <Actionsheet.Item
          leftIcon={<Icon as={Ionicons} name='cog-outline' />}
          _icon={{
            size: 6,
            color: 'black'
          }}
        >
          Account settings
        </Actionsheet.Item>
        <Actionsheet.Item
          leftIcon={<Icon as={Ionicons} name='log-out-outline' />}
          _icon={{
            size: 6,
            color: 'black'
          }}
        >
          Log out
        </Actionsheet.Item>
      </Actionsheet.Content>
    </Actionsheet>
  )
}

export default ProfileActionsheet
