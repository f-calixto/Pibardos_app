import { useDisclose } from 'native-base'

import GroupsButtonView from './GroupsButtonView'

const GroupsButtonContainer = () => {
  const {
    isOpen,
    onOpen,
    onClose
  } = useDisclose()

  return (
    <GroupsButtonView
      isOpen={isOpen}
      onOpen={onOpen}
      onClose={onClose}
    />
  )
}

export default GroupsButtonContainer
