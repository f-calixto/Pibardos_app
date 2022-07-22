import { Button, Container, Icon, Center, Row, ScrollView } from 'native-base'
import ViewWithBars from '@Containers/ViewWithBars'
import Ionicons from 'react-native-vector-icons/Ionicons'
import DebtsOverview from './components/DebtsOverview'
import DebtsSummary from './components/DebtsSummary'
import UserDebtItem from './components/UserDebtItem'

const DebtsScreenView = () => {
  return (
    <ViewWithBars pageName='Debts'>
      <ScrollView>
        <Center flexGrow={1}>
          <Container flexGrow={1} w='full'>
            <Row my={5}>
              <Button
                leftIcon={<Icon as={<Ionicons name='add-outline' />} />}
                mr={3}
              >
                Request a debt
              </Button>
              <Button
                leftIcon={<Icon as={<Ionicons name='checkmark-outline' />} />}
              >
                Confirm a payment
              </Button>
            </Row>

            <DebtsOverview
              title='Owe me'
              amount='420'
              amountColor='green.400'
              people='3'
              pending='2'
            />

            <DebtsOverview
              title='I owe'
              amount='250'
              amountColor='danger.400'
              people='2'
              pending='1'
            />

            <DebtsSummary
              title='Who owes me'
              buttonText='See all'
            >
              {['pending', 'accepted', 'rejected', 'paid'].map((item, idx) => (
                <UserDebtItem key={idx} status={item} />
              ))}
            </DebtsSummary>

            <DebtsSummary
              title='Who I owe'
              buttonText='See all'
            >
              {Array(3).fill(0).map((item, idx) => (
                <UserDebtItem key={idx} />
              ))}
            </DebtsSummary>
          </Container>
        </Center>
      </ScrollView>
    </ViewWithBars>
  )
}

export default DebtsScreenView
