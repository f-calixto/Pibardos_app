import React from 'react'
import { Box, Text, Flex } from 'native-base'
import theme from '@Theme'

const DebtSummary = ({ user, owed, owe }) => {
  return (
    <Flex borderWidth='1' borderRadius='14px' alignItems='center' marginX='10px' bgColor={theme.colors.lightGrey}>
      <Box>
        <Text fontSize='15px'>{user}`s Summary</Text>
      </Box>

      <Flex width='full' flexDirection='row' justifyContent='space-around' marginY='20px'>
        <Box >
          <Text fontWeight='bold' fontSize='20px'>Owed</Text>
          <Text fontWeight='bold' fontSize='20px' color='green.500'>${owed}</Text>
        </Box>
        <Box>
          <Text fontWeight='bold' fontSize='20px'>Owe</Text>
          <Text fontWeight='bold' fontSize='20px' color='red.500'>${owe}</Text>
        </Box>
      </Flex>
    </Flex>
  )
}

export default DebtSummary
