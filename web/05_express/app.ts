import express from 'express'
import * as path from 'path'
import { v4 as uuid } from 'uuid'
import methodOverride from 'method-override'

const PORT = 8080
const HOST = '0.0.0.0'

const app = express()

app.use(express.urlencoded({ extended: true }))
app.use(express.json())
app.use(methodOverride('_method'))
app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'ejs')

let comments = [
  {
    id: uuid(),
    author: 'Corina',
    comment: 'lol that is so funny',
  },
  {
    id: uuid(),
    author: 'Oliver',
    comment: 'I love Frida',
  },
  {
    id: uuid(),
    author: 'Tati',
    comment: 'I wanna party',
  },
  {
    id: uuid(),
    author: 'Frida',
    comment: 'woof woof',
  },
  {
    id: uuid(),
    author: 'Mia',
    comment: 'meow meow',
  },
  {
    id: uuid(),
    author: 'Fiona',
    comment: '... ...',
  },
]

app.get('/comments', (_, res) => {
  res.render('comments/index', { comments })
})

app.get('/comments/new', (_, res) => {
  res.render('comments/new')
})

app.post('/comments', (req, res) => {
  const { author, comment } = req.body
  comments.push({ id: uuid(), author: author, comment: comment })
  res.redirect('/comments')
})

app.get('/comments/:id', (req, res) => {
  const { id } = req.params
  const comment = comments.find((c) => c.id === id)
  res.render('comments/show', { comment })
})

app.get('/comments/:id/edit', (req, res) => {
  const { id } = req.params
  const comment = comments.find((c) => c.id === id)
  res.render('comments/update', { comment })
})

app.patch('/comments/:id', (req, res) => {
  const { id } = req.params
  const newC = { id: id, author: req.body.author, comment: req.body.comment }
  const oldCI = comments.findIndex((c) => c.id === id)
  if (oldCI !== undefined) {
    comments.splice(oldCI, 1, newC)
  }
  res.redirect('/comments')
})

app.delete('/comments/:id', (req, res) => {
  const { id } = req.params
  const comI = comments.findIndex((c) => c.id === id)
  comments.splice(comI, 1)
  res.redirect('/comments')
})

app.listen(PORT, HOST, 3, () => {
  console.log(`Listening on port http://${HOST}:${PORT}`)
})
