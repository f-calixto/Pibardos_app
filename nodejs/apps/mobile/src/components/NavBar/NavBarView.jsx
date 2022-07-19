import { Flex } from 'native-base'
import { useLocation } from 'react-router-native'

import NavBarItem from './components/NavBarItem'
import ProfileButton from './components/ProfileButton'

const NavBarView = () => {
  const location = useLocation()

  return (
    <Flex
      flexDir='row'
      justifyContent='space-between'
      alignItems='center'
      px={5}
      bgColor='white'
      h={60}
      >
      <NavBarItem
        icon='home-outline'
        to='/'
        currentPath={location.pathname}
      />
      <NavBarItem
        icon='calendar-outline'
        to='/calendar'
        currentPath={location.pathname}
      />
      <NavBarItem
        icon='card-outline'
        to='/debts'
        currentPath={location.pathname}
      />
      <ProfileButton />
    </Flex>
  )
}

export default NavBarView
