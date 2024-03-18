import express, { NextFunction, Request, Response } from 'express'
import methodOverride from 'method-override'
import * as path from 'path'
import * as mg from 'mongoose'
import { ExpressError } from './utils/ExpressError'
import { campgroundsRouter } from './routes/campgrounds'
import { reviewsRouter } from './routes/reviews'
// @ts-ignore
import engine from 'ejs-mate'

const WEB_PORT = 8080
const WEB_HOST = '0.0.0.0'
const MG_PORT = 27017
const MG_HOST = 'localhost'

mg.connect(`mongodb://${MG_HOST}:${MG_PORT}/yelp-camp`, {
  autoIndex: true,
  autoCreate: true,
  serverSelectionTimeoutMS: 2000,
})

const app = express()

app.engine('ejs', engine)
app.set('view engine', 'ejs')
app.set('views', path.join(__dirname, 'views'))

app.use(express.urlencoded({ extended: true }))
app.use(express.json())
app.use(methodOverride('_method'))
app.use(express.static(path.join(__dirname, 'public')))

app.use('/campgrounds', campgroundsRouter)
app.use('/campgrounds/:id/reviews', reviewsRouter)

app.all('*', (_, __, next) => {
  next(new ExpressError('Page Not Found', 404))
})

app.use((err: ExpressError, _: Request, res: Response, __: NextFunction) => {
  const { statusCode = 500 } = err
  if (!err.message) err.message = 'Internal Server Error'
  res.status(statusCode).render('error', { err })
})

app.listen(WEB_PORT, WEB_HOST, 3, () => {
  console.log(`Listening on port http://${WEB_HOST}:${WEB_PORT}`)
})
