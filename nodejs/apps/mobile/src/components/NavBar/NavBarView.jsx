import { Flex } from 'native-base'

import NavBarItem from './components/NavBarItem'

const NavBarView = () => {
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
      />
      <NavBarItem
        icon='calendar-outline'
      />
      <NavBarItem
        selected
        icon='card-outline'
        to='/debts'
      />
      <NavBarItem
        icon='person-outline'
        to='/profile'
      />
    </Flex>
  )
}

export default NavBarView
