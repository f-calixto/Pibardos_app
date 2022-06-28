import { Flex, Actionsheet, Avatar, Text } from 'native-base'

const ActionsheetItem = ({ name, imageUri }) => {
  return (
    <Actionsheet.Item>
      <Flex flexDir='row' alignItems='center'>
        <Avatar
          source={{ uri: imageUri }}
          mr={3}
          size={12}
        />
        <Text fontSize={16}>{name}</Text>
      </Flex>
    </Actionsheet.Item>
  )
}

export default ActionsheetItem
