import { extendTheme } from 'native-base'

const useCustomTheme = () => {
  const theme = extendTheme({
    components: {
      Container: {
        defaultProps: {
          maxWidth: '90%'
        }
      },
      Button: {
        defaultProps: {
          colorScheme: 'gray',
          bg: 'gray.900',
          _text: {
            color: 'text.50',
            fontSize: '12',
            fontWeight: 'semibold'
          },
          _icon: {
            size: 'lg'
          }
        }
      }
    }
  })

  return theme
}

export default useCustomTheme
