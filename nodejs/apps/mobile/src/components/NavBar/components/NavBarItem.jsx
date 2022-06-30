import { IconButton } from 'native-base'

const NavBarItem = ({ icon, ...rest }) => {
  return (
    <IconButton
      icon={icon}
      borderRadius='full'
      _icon={{
        color: 'white',
        size: 25
      }}
      {...rest}
    />
  )
}

export default NavBarItem
