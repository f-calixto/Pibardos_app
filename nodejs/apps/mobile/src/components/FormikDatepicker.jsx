import { useState } from 'react'
import { FormControl, Icon, Input, Text, WarningOutlineIcon } from 'native-base'
import { useField } from 'formik'
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import DateTimePickerModal from 'react-native-modal-datetime-picker'
import theme from '../../theme'
import formatDate from '../utils/formatDate'

const FormikDatepicker = ({ name, placeholder, ...props }) => {
  const [field, meta, helpers] = useField(name)
  const showError = meta.touched && meta.error

  const [isDatePickerVisible, setDatePickerVisibility] = useState(false)

  const showDatePicker = () => {
    setDatePickerVisibility(true)
  }

  const hideDatePicker = () => {
    setDatePickerVisibility(false)
  }

  const handleConfirm = (date) => {
    helpers.setValue(date)
    hideDatePicker()
  }

  return (
    <FormControl mb={theme.fontSizes.medium} isInvalid={showError}>
      {field.value !== undefined &&
        <FormControl.Label>
          {placeholder}
        </FormControl.Label>
      }
        <Input
          showSoftInputOnFocus={false}
          onPressIn={showDatePicker}
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
            />
          }
          {...props}
        />
      <DateTimePickerModal
        isVisible={isDatePickerVisible}
        mode='date'
        date={field.value}
        onConfirm={handleConfirm}
        onCancel={hideDatePicker}
      />
      <FormControl.ErrorMessage leftIcon={<WarningOutlineIcon size='xs' />}>
        <Text fontSize={theme.fontSizes.small}>{meta.error}</Text>
      </FormControl.ErrorMessage>
    </FormControl>
  )
}

export default FormikDatepicker
