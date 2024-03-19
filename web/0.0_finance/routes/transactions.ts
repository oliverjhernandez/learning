import express, { Request, Response, NextFunction } from 'express'
import { Transaction } from '../models/Transactions'
import { errorHandler } from '../utils/ErrorHandler'
import { transactionSchema } from '../schemas/transactions'
import { ExpressError } from '../utils/ExpressError'

const router = express.Router({})

const validateSchema = (req: Request, _: Response, next: NextFunction) => {
  const { error } = transactionSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

router.get('/', async (_, res) => {
  const transactions = await Transaction.find()
  res.render('index', { transactions })
})

router.get('/new', async (_, res) => {
  res.render('new')
})

router.get('/:id', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findById(id)
  res.render('show', { transaction })
})

router.post(
  '',
  validateSchema,
  errorHandler(async (req: Request, res: Response, _: NextFunction) => {
    const transaction = new Transaction({ ...req.body })
    await transaction.save()
    res.redirect('')
  })
)

router.get('/:id/edit', async (req, res) => {
  const { id } = req.params
  const transaction = await Transaction.findById(id)
  res.render('edit', { transaction })
})

router.patch(
  '/:id',
  validateSchema,
  errorHandler(async (req, res) => {
    const { id } = req.params
    await Transaction.findByIdAndUpdate(id, { ...req.body })
    res.redirect('')
  })
)

router.delete('/transaction/:id', async (req, res) => {
  const { id } = req.params
  await Transaction.findByIdAndDelete(id)
  res.redirect('')
})

export { router as transactionsRouter }
