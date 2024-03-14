import express from 'express'
import methodOverride from 'method-override'
import * as mg from 'mongoose'
import * as path from 'path'
import { Transaction } from './models/Transactions'
// @ts-ignore
import engine from 'ejs-mate'

const WEB_PORT = 8080
const WEB_HOST = '0.0.0.0'
const MG_PORT = 27017
const MG_HOST = '0.0.0.0'

mg.connect(`mongodb://${MG_HOST}:${MG_PORT}/finance`, {
  autoIndex: true,
  autoCreate: true,
  serverSelectionTimeoutMS: 2000,
})

const app = express()

app.use(express.urlencoded({ extended: true }))
app.use(express.json())
app.use(methodOverride('_method'))
app.engine('ejs', engine)
app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'ejs')

//// Entries

app.get('/transactions', async (_, res) => {
  const transactions = await Transaction.find()
  res.render('index', { transactions })
})

app.get('/transactions/new', async (_, res) => {
  res.render('new')
})

app.get('/transactions/:id', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findById(id)
  res.render('show', { transaction })
})

app.post('/transactions', async (req, res) => {
  const transaction = new Transaction({ ...req.body })
  await transaction.save()
  res.redirect('/transactions')
})

app.get('/transactions/:id/edit', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findById(id)
  res.render('edit', { transaction })
})

app.patch('/transactions/:id', async (req, res) => {
  const { id } = req.params
  await Transaction.findByIdAndUpdate(id, { ...req.body })
  res.redirect('/transactions')
})

app.delete('/transaction/:id', async (req, res) => {
  const { id } = req.params
  await Transaction.findByIdAndDelete(id)
  res.redirect('/transactions')
})

app.listen(WEB_PORT, WEB_HOST, () => {
  console.log(`Listening on http://${WEB_HOST}:${WEB_PORT}`)
})
