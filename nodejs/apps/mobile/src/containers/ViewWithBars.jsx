import { Flex } from 'native-base'
import TopBar from '@Components/TopBar'
import NavBar from '@Components/NavBar'

const ViewWithBars = ({ pageName, children }) => {
  return (
    <Flex h='full' justifyContent='space-between' backgroundColor='gray.100'>
      <TopBar pageName={pageName} />
      {children}
      <Flex justify='flex-end'>
        <NavBar />
      </Flex>
    </Flex>
  )
}

export default ViewWithBars
