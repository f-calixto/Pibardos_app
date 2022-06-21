import { useState } from 'react'
import { Input, KeyboardAvoidingView } from 'native-base'
import { Platform } from 'react-native'

const CountryFilterInput = ({ onFilterChange }) => {
  const [value, setValue] = useState('')

  const onValueChange = newValue => {
    setValue(newValue)
    onFilterChange(newValue)
  }

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
    >
      <Input
        h={12}
        mb={5}
        placeholder='Buscar país'
        value={value}
        onChangeText={onValueChange}
      />
    </KeyboardAvoidingView>
  )
}

export default CountryFilterInput
