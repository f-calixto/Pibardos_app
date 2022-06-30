import { View, Text, Button, Flex } from 'native-base'
import ViewWithBars from '@Containers/ViewWithBars'
import DebtSummary from './components/DebtSummary'
import theme from '@Theme'
const DebtsScreenView = () => {
  return (
    <ViewWithBars>
      <View>
        <Flex direction='row' justifyContent='space-between' marginTop='10px'>
        <Text
        fontSize='30px'
        fontWeight='bold'
        marginBottom='10px'
        marginLeft='20px'
        >Group Debts</Text>
        <Button height='40px' width='100px' marginRight='20px' bgColor= {theme.colors.blue}>
        Requests
        </Button>
        </Flex>

        <DebtSummary
        user='ivo'
        owe='10'
        owed='20'/>
      </View>
    </ViewWithBars>
  )
}

export default DebtsScreenView
