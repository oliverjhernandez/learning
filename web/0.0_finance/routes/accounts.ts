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

router.get('/', async (_, res) => {
  const accounts = await Account.find()
  res.render('accounts/index', { accounts })
})

router.get('/new', async (_, res) => {
  res.render('accounts/new')
})

router.get('/:id', async (req, res) => {
  const { id } = req.params
  const account = await Account.findById(id)
  res.render('accounts/show', { account })
})

router.post(
  '/',
  validateSchema,
  errorHandler(async (req: Request, res: Response, _: NextFunction) => {
    const account = new Account({ ...req.body })
    await account.save()
    res.redirect('')
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
    await Account.findByIdAndUpdate(id, { ...req.body })
    res.redirect('')
  })
)

router.delete('/:id', async (req, res) => {
  const { id } = req.params
  await Account.findByIdAndDelete(id)
  res.redirect('')
})

export { router as accountsRouter }
