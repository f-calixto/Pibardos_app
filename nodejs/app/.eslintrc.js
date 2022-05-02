module.exports = {
  env: {
    'react-native/react-native': true
  },
  extends: [
    'plugin:react/recommended',
    'standard'
  ],
  // parser: '@babel/eslint-parser',
  parserOptions: {
    ecmaFeatures: {
      jsx: true
    },
    ecmaVersion: 'latest',
    sourceType: 'module'
  },
  plugins: [
    'react',
    'react-native'
  ],
  settings: {
    react: {
      version: 'detect'
    }
  },
  rules: {
    // ...
    'react/jsx-uses-react': 'off',
    'react/react-in-jsx-scope': 'off'
  }
}
