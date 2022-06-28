import { SafeAreaView } from 'react-native'
import { View, Flex, Image, StatusBar } from 'native-base'
import theme from '@Theme'

// images
import Favicon from '@Assets/logos/favicon.png'
import GroupsButton from './components/GroupsButton'
import ProfileButton from './components/ProfileButton'

const TopBarView = () => {
  return (
    <View backgroundColor={theme.colors.primary} borderBottomRadius={20}>
      <StatusBar barStyle='light-content' backgroundColor={theme.colors.primary}/>
      <SafeAreaView>
        <Flex
          flexDir='row'
          justifyContent='space-between'
          alignItems='center'
          py={3}
          px={4}
        >
          <GroupsButton />
          <Image source={Favicon} alt='app-logo' w={10} h={10} />
          <ProfileButton />
        </Flex>
      </SafeAreaView>
    </View>
  )
}

export default TopBarView
