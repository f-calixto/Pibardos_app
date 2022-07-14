import { Flex } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'
import { useLinkPressHandler } from 'react-router-native'

const NavBarItem = ({
  icon,
  replace = false,
  state,
  to,
  selected,
  children,
  onPress,
  ...rest
}) => {
  const handlePress = useLinkPressHandler(to, {
    replace,
    state
  })

  return (
    <Flex
      justifyContent='center'
      alignItems='center'
      w={10}
      h={10}
      p={1}
      mb={selected && 3}
      bgColor={selected && 'gray.900'}
      borderRadius='full'
      {...rest}
      onPress={event => {
        if (onPress) {
          onPress()
        } else if (!event.defaultPrevented) {
          handlePress(event)
        }
      }}
    >
      <Ionicons name={icon} size={25} color={selected ? '#FAFAFA' : '#71717A'} />
    </Flex>
  )
}

export default NavBarItem
