import express, { Request, Response, NextFunction } from 'express'
import session, { SessionOptions } from 'express-session'
import methodOverride from 'method-override'
import mg from 'mongoose'
import path from 'path'
import { ExpressError } from './utils/ExpressError'
// @ts-ignore
import engine from 'ejs-mate'
import { transactionsRouter } from './routes/transactions'
import { accountsRouter } from './routes/accounts'
import flash from 'connect-flash'
import passport from 'passport'
import { Strategy as LocalStrategy } from 'passport-local'
import { User } from './models/User'

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

app.engine('ejs', engine)
app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'ejs')

app.use(express.urlencoded({ extended: true }))
app.use(express.json())
app.use(methodOverride('_method'))
app.use(express.static(path.join(__dirname, 'public')))

const sessionConfig: SessionOptions = {
  secret: 'fn5y68u932-fn684q2n7g5vy27895',
  resave: false,
  saveUninitialized: true,
  cookie: {
    maxAge: 1000 * 60 * 60 * 24 * 7, // a week in milliseconds
    httpOnly: true,
  },
}

app.use(session(sessionConfig))
app.use(flash())
// TODO: Improve authentication method
app.use(passport.initialize())
app.use(passport.session())

passport.use(new LocalStrategy(User.authenticate()))
passport.serializeUser(User.serializeUser()) // TODO: Improve passport workflow
passport.deserializeUser(User.deserializeUser())

app.use((req, res, next) => {
  res.locals.currentUser = req.user
  res.locals.success = req.flash('success')
  res.locals.error = req.flash('error')
  next()
})

app.use('/transactions', transactionsRouter)
app.use('/accounts', accountsRouter)

app.all('*', (_, __, next) => {
  next(new ExpressError('Page Not Found', 404))
})

app.use((err: ExpressError, _: Request, res: Response, __: NextFunction) => {
  const { statusCode = 500 } = err
  if (!err.message) err.message = 'Internal Server Error'
  res.status(statusCode).render('general/error', { err })
})

app.listen(WEB_PORT, WEB_HOST, () => {
  console.log(
    `${new Date().toLocaleString()} - Listening on http://${WEB_HOST}:${WEB_PORT}`
  )
})
