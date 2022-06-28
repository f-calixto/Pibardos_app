import Ionicons from 'react-native-vector-icons/Ionicons'
import TopBarItem from '../TopBarItem'
import ProfileActionsheet from './components/ProfileActionsheet'

const ProfileButtonView = ({ isOpen, onOpen, onClose }) => {
  return (
    <>
      <TopBarItem
        icon={<Ionicons name='person-outline'/>}
        onPress={onOpen}
      />
      <ProfileActionsheet
        isOpen={isOpen}
        onClose={onClose}
      />
    </>
  )
}

export default ProfileButtonView
