import express from 'express'
import * as path from 'path'
import methodOverride from 'method-override'
import { commentsRepo, IComment } from './repository/Comments'
import { connect } from 'mongoose'

const PORT = 8080
const HOST = '0.0.0.0'

const app = express()

app.use(express.urlencoded({ extended: true }))
app.use(express.json())
app.use(methodOverride('_method'))
app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'ejs')

connect('mongodb://localhost:27017/comments', {
  serverSelectionTimeoutMS: 5000,
})

app.get('/comments', async (_, res) => {
  const comments = await commentsRepo.fetchAll()
  res.render('comments/index', { comments })
})

app.get('/comments/new', (_, res) => {
  res.render('comments/new')
})

app.post('/comments', (req, res) => {
  const { author, comment } = req.body
  const comm: Partial<IComment> = {
    author: author,
    comment: comment,
  }
  commentsRepo.storeComment(comm)
  res.redirect('/comments')
})

app.get('/comments/:id', async (req, res) => {
  const { id } = req.params
  const comment = await commentsRepo.fetchCommentById(id)
  res.render('comments/show', { comment })
})

app.get('/comments/:id/edit', async (req, res) => {
  const { id } = req.params
  const comment = await commentsRepo.fetchCommentById(id)
  res.render('comments/update', { comment })
})

app.patch('/comments/:id', async (req, res) => {
  const { id } = req.params
  const newComment: Partial<IComment> = {
    author: req.body.author,
    comment: req.body.comment,
  }
  commentsRepo.updateDoc(id, newComment)
  res.redirect('/comments')
})

app.delete('/comments/:id', (req, res) => {
  const { id } = req.params
  commentsRepo.deleteDoc(id)
  res.redirect('/comments')
})

app.listen(PORT, HOST, 3, () => {
  console.log(`Listening on port http://${HOST}:${PORT}`)
})
