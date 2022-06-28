import Ionicons from 'react-native-vector-icons/Ionicons'
import TopBarItem from '../TopBarItem'
import GroupsActionsheet from './components/GroupsActionsheet'

const GroupsButtonView = ({ isOpen, onOpen, onClose }) => {
  return (
    <>
      <TopBarItem
        icon={<Ionicons name='people-outline' />}
        onPress={onOpen}
      />
      <GroupsActionsheet
        isOpen={isOpen}
        onClose={onClose}
      />
    </>
  )
}

export default GroupsButtonView
