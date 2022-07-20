import TopBarItem from '../TopBarItem'
import GroupsActionsheet from './components/GroupsActionsheet'

const GroupsButtonView = ({ isOpen, onOpen, onClose }) => {
  return (
    <>
      <TopBarItem
        icon='people-outline'
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
