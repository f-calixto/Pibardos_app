import { Icon, Box, Image, Menu } from 'native-base'
import { Pressable } from 'react-native'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
// import group from './assets/logos/group.png'

// use props to pass name of groups user is in and icon - pass as assoc array?
const HeaderBar = () => {
  return (
    // Header Box
    <Box
      borderBottomWidth='2'
      p='3'
      safeAreaTop='5'
      flexDirection='row'
      justifyContent='space-between'
    >
      {/* Groups Menu */}
      <Menu
        w='190'
        trigger={(triggerProps) => {
          return (
            <Pressable {...triggerProps}>
              <Image
                source={require('../assets/logos/group.png')}
                alt='logo'
              ></Image>
            </Pressable>
          )
        }}
      >
        <Menu.Item flexDirection='row'>
          <Icon mr='2' size='md' as={<MaterialIcons name='settings' />} />
          Gestionar grupos
        </Menu.Item>
      </Menu>

      <Image source={require('../assets/logos/favicon.png')} alt='logo'></Image>

      <Menu
        w='190'
        trigger={(triggerProps) => {
          return (
            <Pressable {...triggerProps}>
              <Image
                source={require('../assets/logos/user.png')}
                alt='logo'
              ></Image>
            </Pressable>
          )
        }}
      >
        <Menu.Item flexDirection='row'>
          <Icon mr='2' size='md' as={<MaterialIcons name='settings' />} />
          Ajustes de perfil
        </Menu.Item>
        <Menu.Item flexDirection='row'>
          <Icon mr='2' size='md' as={<MaterialIcons name='logout' />} />
          Cerrar sesion
        </Menu.Item>
      </Menu>
    </Box>
  )
}

export default HeaderBar
