import Joi from 'Joi'

export const transactionsSchema = Joi.object({
  name: Joi.string().required(),
  date: Joi.date().required(),
  accountId: Joi.string().required(),
  amount: Joi.number().required().min(0),
  currency: Joi.string().required(),
  type: Joi.string().required(),
  categoryId: Joi.string(),
  notes: Joi.string(),
})
