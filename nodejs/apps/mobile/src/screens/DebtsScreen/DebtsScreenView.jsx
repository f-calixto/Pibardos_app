import {
  View,
  Heading,
  Button,
  Flex,
  Box,
  Text,
  Spacer,
  HStack,
  Avatar,
  VStack,
  FlatList,
  ScrollView
} from 'native-base'
import ViewWithBars from '@Containers/ViewWithBars'
import DebtSummary from './components/DebtSummary'
import theme from '@Theme'
import IonIcon from 'react-native-vector-icons/Ionicons'

const DebtsScreenView = () => {
  const data = [
    {
      id: 'bd7acbea-c1b1-46c2-aed5-3ad53abb28ba',
      fullName: 'Moliber',
      timeStamp: '12:47 PM',
      recentText: 'Good Day!',
      avatarUrl:
        'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500'
    },
    {
      id: '3ac68afc-c605-48d3-a4f8-fbd91aa97f63',
      fullName: 'Mazen',
      timeStamp: '11:11 PM',
      recentText: 'Cheer up, there!',
      avatarUrl:
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyEaZqT3fHeNrPGcnjLLX1v_W4mvBlgpwxnA&usqp=CAU'
    },
    {
      id: '58694a0f-3da1-471f-bd96-145571e29d72',
      fullName: 'Meke',
      timeStamp: '6:22 PM',
      recentText: 'Good Day!',
      avatarUrl: 'https://miro.medium.com/max/1400/0*0fClPmIScV5pTLoE.jpg'
    },
    {
      id: '68694a0f-3da1-431f-bd56-142371e29d72',
      fullName: 'Miko',
      timeStamp: '8:56 PM',
      recentText: 'All the best',
      avatarUrl:
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSr01zI37DYuR8bMV5exWQBSw28C1v_71CAh8d7GP1mplcmTgQA6Q66Oo--QedAN1B4E1k&usqp=CAU'
    },
    {
      id: '28694a0f-3da1-471f-bd96-142456e29d72',
      fullName: 'Kiara',
      timeStamp: '12:47 PM',
      recentText: 'I will call today.',
      avatarUrl:
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRBwgu1A5zgPSvfE83nurkuzNEoXs9DMNr8Ww&usqp=CAU'
    },
    {
      id: '28694a0f-3da1-471f-bd96-142456e29d72',
      fullName: 'Kiara',
      avatarUrl:
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRBwgu1A5zgPSvfE83nurkuzNEoXs9DMNr8Ww&usqp=CAU'
    }
  ]
  return (
    <ViewWithBars>
      <View>
        <Flex direction='row' justifyContent='space-between' marginTop='10px'>
          <Heading
            fontSize='30px'
            fontWeight='bold'
            marginBottom='10px'
            marginLeft='20px'
          >
            Group Debts
          </Heading>
          <Button
            height='40px'
            width='100px'
            marginRight='20px'
            bgColor={theme.colors.blue}
          >
            Requests
          </Button>
        </Flex>

        <DebtSummary user='Ivo' owe='100' owed='200' />
        <ScrollView
          axW='300'
          h='80'
          _contentContainerStyle={{
            px: '20px',
            mb: '4',
            minW: '72'
          }}
        >
          <FlatList
            data={data}
            renderItem={({ item }) => (
              <Box
                borderBottomWidth='1'
                _dark={{
                  borderColor: 'gray.600'
                }}
                borderColor='coolGray.200'
                pl='4'
                pr='5'
                py='2'
              >
                <HStack space={3} justifyContent='space-between'>
                  <Avatar
                    size='48px'
                    source={{
                      uri: item.avatarUrl
                    }}
                  />
                  <VStack>
                    <Text
                      _dark={{
                        color: 'warmGray.50'
                      }}
                      color='coolGray.800'
                      bold
                    >
                      {item.fullName} owes:
                    </Text>
                    <Text
                      color='coolGray.600'
                      _dark={{
                        color: 'warmGray.200'
                      }}
                    >
                      amount aca
                    </Text>
                    <Text color='coolGray.600' _dark={{
                      color: 'warmGray.200'
                    }}>
                  debt count aca
                </Text>
                  </VStack>
                  <Spacer />
                  <View alignSelf='flex-start'>
                  <IonIcon name='arrow-forward' size={20} />
                  </View>
                </HStack>
              </Box>
            )}
            keyExtractor={(item) => item.id}
          />
        </ScrollView>
      </View>
    </ViewWithBars>
  )
}

export default DebtsScreenView