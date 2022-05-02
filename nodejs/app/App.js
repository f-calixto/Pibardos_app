import { StyleSheet, View, Text, Pressable, Alert } from 'react-native'

const PressableText = props => {
  return (
    <View style={styles.container}>
      <Pressable
        onPress={() => Alert.prompt('eee', 'anasheee')}
      >
        <Text>You can press me</Text>
      </Pressable>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center'
  }
})

export default PressableText
