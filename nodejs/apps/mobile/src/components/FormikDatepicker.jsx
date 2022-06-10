import { useState } from 'react'
import { Box, FormControl, Icon, Input, Text, WarningOutlineIcon } from 'native-base'
import { useField } from 'formik'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import DateTimePickerModal from 'react-native-modal-datetime-picker'
import theme from '../../theme'
import formatDate from '../utils/formatDate'

const FormikDatepicker = ({ name, placeholder, ...props }) => {
  const [field, meta, helpers] = useField(name)
  const [show, setShow] = useState(false)
  const showError = meta.touched && meta.error

  const [isDatePickerVisible, setDatePickerVisibility] = useState(false)

  const showDatePicker = () => {
    setDatePickerVisibility(true)
  }

  const hideDatePicker = () => {
    setDatePickerVisibility(false)
  }

  const handleConfirm = (date) => {
    console.warn('A date has been picked: ', date)
    hideDatePicker()
  }

  return (
    <Box mb={theme.fontSizes.medium}>
      <FormControl isInvalid={showError}>
        {field.value.length > 0 &&
          <FormControl.Label>
            {placeholder}
          </FormControl.Label>
        }
        <Input
          isReadOnly
          fontSize={theme.fontSizes.small}
          h={12}
          placeholder={placeholder}
          onBlur={() => helpers.setTouched(true)}
          value={formatDate(field.value)}
          rightElement={
            <Icon
              as={<MaterialIcons name='event' />}
              mr={theme.fontSizes.small}
              size={theme.fontSizes.large}
              onPress={() => setShow(true)}
            />
          }
          {...props}
        />
        <Button title='Show Date Picker' onPress={showDatePicker} />
        <DateTimePickerModal
          isVisible={isDatePickerVisible}
          mode='date'
          onConfirm={handleConfirm}
          onCancel={hideDatePicker}
        />
        <FormControl.ErrorMessage leftIcon={<WarningOutlineIcon size='xs' />}>
          <Text fontSize={theme.fontSizes.small}>{meta.error}</Text>
        </FormControl.ErrorMessage>
      </FormControl>
    </Box>
  )
}

export default FormikDatepicker
