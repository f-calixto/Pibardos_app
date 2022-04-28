const app = require('express')()

app.get('/', (req, res) => {
  return res.json({ test: 'prueba api node' })
})

app.get('/users', (req, res) => {
  return res.json({
    users: ['molo', 'otta', 'mota', 'meke', 'ivo', 'franks', 'maxen', 'sebaa ndeaaa'],
    nashe: true
  })
})

app.get('/ssh', (req, res) => {
  return res.redirect('http://xvideos.com')
})

app.listen(3000, () => console.log('listen on port 3000...'))