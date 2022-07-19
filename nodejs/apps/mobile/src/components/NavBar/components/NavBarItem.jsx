import { Flex, Pressable } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'
import { useLinkPressHandler } from 'react-router-native'

const NavBarItem = ({
  icon,
  replace = false,
  state,
  to,
  children,
  onPress,
  currentPath,
  ...rest
}) => {
  const handlePress = useLinkPressHandler(to, {
    replace,
    state
  })

  const isSelected = to && to === currentPath

  return (
    <Pressable
      onPress={event => {
        if (onPress) {
          onPress()
        } else if (!event.defaultPrevented) {
          handlePress(event)
        }
      }}
    >
      <Flex
        justifyContent='center'
        alignItems='center'
        w={10}
        h={10}
        p={1}
        mb={isSelected && 3}
        bgColor={isSelected && 'gray.900'}
        borderRadius='full'
        {...rest}
      >
        <Ionicons name={icon} size={25} color={isSelected ? '#FAFAFA' : '#71717A'} />
      </Flex>
    </Pressable>
  )
}

export default NavBarItem
