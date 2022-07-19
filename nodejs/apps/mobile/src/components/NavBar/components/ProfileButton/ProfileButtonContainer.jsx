import { useDisclose } from 'native-base'
import ProfileButtonView from './ProfileButtonView'

const ProfileButtonContainer = () => {
  const {
    isOpen,
    onOpen,
    onClose
  } = useDisclose()

  return (
    <ProfileButtonView
      isOpen={isOpen}
      onOpen={onOpen}
      onClose={onClose}
    />
  )
}

export default ProfileButtonContainer
