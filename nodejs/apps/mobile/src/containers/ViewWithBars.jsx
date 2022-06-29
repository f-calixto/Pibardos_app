import { Flex, View } from 'native-base'
import TopBar from '@Components/TopBar'
import NavBar from '@Components/NavBar'

const ViewWithBars = ({ children }) => {
  return (
    <Flex h='full' maxH='full' justifyContent='space-between'>
      <TopBar />
      <View flexGrow={1}>
        {children}
      </View>
      <Flex justify='flex-end'>
        <NavBar />
      </Flex>
    </Flex>
  )
}

export default ViewWithBars
