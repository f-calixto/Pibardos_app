import { Heading, Button, Flex, Text, ScrollView, Popover } from 'native-base'
import HeaderBar from '../HeaderBar'
import theme from '../../theme'

const GroupsScreen = () => {
  return (
    // <Header>Hola</Header>
    // <Text>TEst</Text>
    <Flex flex='1'>
      <HeaderBar></HeaderBar>

      <Flex align='center' pt='30%'>
        <Heading size='2xl'>Gestionar Grupos</Heading>
      </Flex>

    <Flex mt='10' align='center'>
        <ScrollView justify='center' bgColor='pink' maxHeight='100' w='90%' >
            <Text>Group 1</Text>
            <Text>Group 2</Text>
            <Text>Group 3</Text>
            <Text>Group 4</Text>
            <Text>Group 5</Text>
            <Text>Group 6</Text>
            <Text>Group 7</Text>
            <Text>Group 8</Text>
            <Text>Group 9</Text>
            <Text>Group 0</Text>
            <Text>Group 11</Text>
            <Text>Group 12</Text>

        </ScrollView>
    </Flex>

      <Flex justify='center' align='center'>
        <Button
          width='40%'
          bg={theme.colors.blue}
          height='18%'
          onPress={() => console.log('crear grupo press')}
        >
          Crear Grupo
        </Button>
        <Button
          mt='3%'
          bg={theme.colors.green}
          width='40%'
          height='18%'
          onPress={() => console.log('unirse a grupo pressed')}
        >
          Unirse a un grupo
        </Button>
      </Flex>
    </Flex>
  )
}

export default GroupsScreen
