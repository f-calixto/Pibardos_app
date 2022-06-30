import { Flex } from 'native-base'
import Ionicons from 'react-native-vector-icons/Ionicons'
import theme from '@Theme'

import NavBarItem from './components/NavBarItem'

const NavBarView = () => {
  return (
    <Flex
      flexDir='row'
      justifyContent='space-between'
      alignItems='center'
      px={5}
      bgColor={theme.colors.primary}
      h={60}
    >
      <NavBarItem
        icon={<Ionicons name='card-outline' />}
        to='/debts'
      />
      <NavBarItem
        icon={<Ionicons name='home-outline' />}
      />
      <NavBarItem
        icon={<Ionicons name='calendar-outline' />}
      />
    </Flex>
  )
}

export default NavBarView
