import { useTheme } from 'native-base'
import { Pressable } from 'react-native'
import Ionicons from 'react-native-vector-icons/Ionicons'

const TopBarItem = ({ icon, onPress, ...rest }) => {
  const theme = useTheme()

  return (
    <Pressable onPress={onPress}>
      <Ionicons
        name={icon}
        color={theme.colors.text[900]}
        size={28}
        {...rest}
      />
    </Pressable>
  )
}

export default TopBarItem
