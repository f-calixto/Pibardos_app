import { IconButton } from 'native-base'
import { useLinkPressHandler } from 'react-router-native'

const NavBarItem = ({
  icon,
  onPress,
  replace = false,
  state,
  to,
  ...rest
}) => {
  const handlePress = useLinkPressHandler(to, {
    replace,
    state
  })

  return (
    <IconButton
      icon={icon}
      borderRadius='full'
      _icon={{
        color: 'white',
        size: 25
      }}
      {...rest}
      onPress={event => {
        if (!event.defaultPrevented) {
          handlePress(event)
        }
      }}
    />
  )
}

export default NavBarItem
