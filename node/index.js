const app = require('express')()

app.get('/', (req, res) => {
  return res.json({ test: 'prueba api node' })
})

app.get('/users', (req, res) => {
  return res.json({
    users: ['molo', 'otta', 'mota', 'meke', 'ivo', 'franks', 'maxen'],
    nashe: true
  })
})

app.listen(80, () => console.log('listen on port 80...'))