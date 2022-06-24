module.exports = function (api) {
  api.cache(true)
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      ['module-resolver', {
        alias: {
          '@Components': './src/components',
          '@Containers': './src/containers',
          '@Screens': './src/screens',
          '@ReduxSlices': './src/redux',
          '@Hooks': './src/hooks',
          '@Assets': './src/assets',
          '@Theme': './theme.js',
          '@Utils': './src/utils',
          '@Services': './src/services'
        },
        extensions: [
          '.js',
          '.jsx',
          '.ts',
          '.tsx'
        ]
      }]
    ]
  }
}
