import express, { Request, Response, NextFunction } from 'express'
import { Transaction } from '../models/Transaction'
import { errorHandler } from '../utils/ErrorHandler'
import { transactionsSchema } from '../schemas/transactions'
import { ExpressError } from '../utils/ExpressError'

const router = express.Router({})

const validateTransaction = (req: Request, _: Response, next: NextFunction) => {
  const { error } = transactionsSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

router.get('/', async (_, res) => {
  const transactions = await Transaction.find()
  res.render('transactions/index', { transactions })
})

router.get('/new', async (_, res) => {
  res.render('transactions/new')
})

router.get('/:id', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findById(id)
  res.render('transactions/show', { transaction })
})

router.post(
  '/',
  validateTransaction,
  errorHandler(async (req: Request, res: Response, _: NextFunction) => {
    const transaction = new Transaction({ ...req.body })
    await transaction.save()
    req.flash('success', 'Transaction created successfully')
    res.redirect(`transactions/${transaction._id}`)
  })
)

router.get('/:id/edit', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findById(id)
  if (!transaction) {
    req.flash('error', 'Could not find transaction')
  }

  res.render('transactions/edit', { transaction })
})

router.patch(
  '/:id',
  validateTransaction,
  errorHandler(async (req, res) => {
    const { id } = req.params
    const transaction = await Transaction.findByIdAndUpdate(id, { ...req.body })
    if (!transaction) {
      req.flash('error', 'Could not update transaction')
    }
    res.redirect(`transactions/${id}`)
  })
)

router.delete('/transaction/:id', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findByIdAndDelete(id)
  if (!transaction) {
    req.flash('error', 'Could not delete transaction')
  }
  res.redirect('transactions')
})

export { router as transactionsRouter }
