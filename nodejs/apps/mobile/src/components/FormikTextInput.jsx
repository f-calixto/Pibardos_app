import { Box, FormControl, Input, WarningOutlineIcon, Text } from 'native-base'
import { useField } from 'formik'
import theme from '../../theme'

const FormikTextInput = ({ name, placeholder, ...props }) => {
  const [field, meta, helpers] = useField(name)
  const showError = meta.touched && meta.error

  return (
    <Box mb={theme.fontSizes.medium}>
      <FormControl isInvalid={showError}>
        {field.value.length > 0 &&
          <FormControl.Label>
            {placeholder}
          </FormControl.Label>
        }
        <Input
          fontSize={theme.fontSizes.small}
          h={12}
          placeholder={placeholder}
          onChangeText={value => helpers.setValue(value)}
          onBlur={() => helpers.setTouched(true)}
          value={field.value}
          {...props}
        />
        <FormControl.ErrorMessage leftIcon={<WarningOutlineIcon size='xs' />}>
          <Text fontSize={theme.fontSizes.small}>{meta.error}</Text>
        </FormControl.ErrorMessage>
      </FormControl>
    </Box>
  )
}

export default FormikTextInput
