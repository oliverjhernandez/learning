import express from 'express'
import { User } from '../models/User'
import { ExpressError } from '../utils/ExpressError'
import { errorCatcher } from '../utils/ErrorCatcher'
import passport from 'passport'
import { storeReturnTo } from '../middleware'

const router = express.Router()

router.get('/register', (_, res) => {
  res.render('users/register')
})

router.post(
  '/register',
  errorCatcher(async (req, res, next) => {
    try {
      const { username, email, password } = req.body
      const newUser = new User({ email, username })
      const registered = await User.register(newUser, password)
      req.login(registered, (error) => {
        if (error) {
          next(error)
        }
        req.flash('success', 'Welcome to YelpCamp!')
        res.redirect('/campgrounds')
      })
    } catch (error) {
      const er = error as ExpressError
      req.flash('error', er.message)
      res.redirect('/register')
    }
  })
)

router.get('/login', (req, res) => {
  console.log('GET', req.session)
  res.render('users/login')
})

router.post(
  '/login',
  storeReturnTo,
  passport.authenticate('local', {
    failureFlash: true,
    failureRedirect: '/login',
  }),
  (req, res) => {
    req.flash('success', 'Welcome back!')
    console.log('POST', req.session)

    if (res.locals.returnTo) {
      const returnTo = res.locals.returnTo
      delete req.session.returnTo
      delete res.locals.returnTo
      return res.redirect(returnTo)
    }
    res.redirect('/campgrounds')
  }
)

router.get('/logout', (req, res, next) => {
  req.logout((err) => {
    if (err) {
      return next(err)
    }
    req.flash('success', 'Bye!')
  })
  res.redirect('/campgrounds')
})

export { router as userRouter }
