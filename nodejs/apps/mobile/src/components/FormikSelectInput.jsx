import { FormControl, Select, CheckIcon, WarningOutlineIcon, Text } from 'native-base'
import { useField } from 'formik'
import theme from '../../theme'

const FormikSelectInput = ({ name, placeholder, items, ...props }) => {
  const [field, meta, helpers] = useField(name)
  const showError = meta.touched && meta.error

  return (
    <FormControl isInvalid={showError}>
      {field.value.length > 0 &&
        <FormControl.Label>
          {placeholder}
        </FormControl.Label>
      }
      <Select
        h={12}
        minWidth='200'
        accessibilityLabel={placeholder}
        placeholder={placeholder}
        onValueChange={value => helpers.setValue(value)}
        selectedValue={field.value}
        _selectedItem={{
          bg: theme.colors.lightGrey,
          endIcon: <CheckIcon size={5} />
        }}
        {...props}
      >
        {items && items.map((item, idx) => (
          <Select.Item
            key={idx}
            label={item.label}
            value={item.value}
          />
        ))}
      </Select>
      <FormControl.ErrorMessage leftIcon={<WarningOutlineIcon size='xs' />}>
        <Text fontSize={theme.fontSizes.small}>{meta.error}</Text>
      </FormControl.ErrorMessage>
    </FormControl>
  )
}

export default FormikSelectInput
