import { SafeAreaView } from 'react-native'
import { View, Flex, StatusBar, Heading } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'

const TopBarView = ({ pageName = 'Untitled' }) => {
  return (
    <View backgroundColor='amber.400'>
      <StatusBar barStyle='dark-content' backgroundColor='amber.400'/>
      <SafeAreaView>
        <Flex
          flexDir='row'
          justifyContent='space-between'
          alignItems='center'
          py={4}
          px={4}
        >
          <Heading size='xl'>{pageName}</Heading>
          <Ionicons name='notifications-outline' size={28} />
          {/* <GroupsButton />
          <ProfileButton /> */}
        </Flex>
      </SafeAreaView>
    </View>
  )
}

export default TopBarView
