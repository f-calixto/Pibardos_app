import { Button } from 'native-base'
import { useLinkPressHandler } from 'react-router-native'

const ButtonLink = ({
  children,
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
    <Button
      {...rest}
      onPress={event => {
        if (!event.defaultPrevented) {
          handlePress(event)
        }
      }}
    >
      {children}
    </Button>
  )
}

export default ButtonLink
