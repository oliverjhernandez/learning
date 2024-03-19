import express, { Request, Response, NextFunction } from 'express'
import methodOverride from 'method-override'
import mg from 'mongoose'
import path from 'path'
import { ExpressError } from './utils/ExpressError'
// @ts-ignore
import engine from 'ejs-mate'
import { transactionsRouter } from './routes/transactions'

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
app.use('/transactions', transactionsRouter)

app.all('*', (_, __, next) => {
  next(new ExpressError('Page Not Found', 404))
})

app.use((err: ExpressError, _: Request, res: Response, __: NextFunction) => {
  const { statusCode = 500 } = err
  if (!err.message) err.message = 'Internal Server Error'
  res.status(statusCode).render('error', { err })
})

app.listen(WEB_PORT, WEB_HOST, () => {
  console.log(`Listening on http://${WEB_HOST}:${WEB_PORT}`)
})
