import NavBarItem from '../NavBarItem'
import ProfileActionsheet from './components/ProfileActionsheet'

const ProfileButtonView = ({ isOpen, onOpen, onClose }) => {
  return (
    <>
      <NavBarItem
        icon='person-outline'
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
