import express, { Request, Response, NextFunction } from 'express'
import { Account } from '../models/Account'
import { accountsSchema } from '../schemas/accounts'
import { ExpressError } from '../utils/ExpressError'
import { errorHandler } from '../utils/ErrorHandler'

const router = express.Router()

const validateSchema = (req: Request, _: Response, next: NextFunction) => {
  const { error } = accountsSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

router.get(
  '/',
  errorHandler(async (_, res) => {
    const accounts = await Account.find()
    res.render('accounts/index', { accounts })
  })
)

router.get(
  '/new',
  errorHandler(async (_, res) => {
    res.render('accounts/new')
  })
)

router.get(
  '/:id',
  errorHandler(async (req, res) => {
    const { id } = req.params
    const account = await Account.findById(id)
    console.log('hello')

    if (!account) {
      req.flash('error', 'Could not find account')
      res.redirect('accounts')
    }
    res.render('accounts/show', { account })
  })
)

router.post(
  '/',
  validateSchema,
  errorHandler(async (req: Request, res: Response, _: NextFunction) => {
    const account = new Account({ ...req.body })
    const savedAccount = await account.save()
    console.log(savedAccount)
    if (!savedAccount) {
      req.flash('error', 'Could not store new account')
      res.redirect('accounts')
    }
    req.flash('success', 'Account created successfully')
    res.redirect(`accounts/${account._id}`)
  })
)

router.get('/:id/edit', async (req, res) => {
  const { id } = req.params
  const account = await Account.findById(id)

  res.render('accounts/edit', { account })
})

router.patch(
  '/:id',
  validateSchema,
  errorHandler(async (req, res) => {
    const { id } = req.params
    const account = await Account.findByIdAndUpdate(id, { ...req.body })
    if (!account) {
      req.flash('error', 'Could not update account')
    }
    res.redirect(`accounts/${id}`)
  })
)

router.delete('/:id', async (req, res) => {
  const { id } = req.params
  const account = await Account.findByIdAndDelete(id)
  if (!account) {
    req.flash('error', 'Could not delete account')
  }
  res.redirect('accounts')
})

export { router as accountsRouter }
