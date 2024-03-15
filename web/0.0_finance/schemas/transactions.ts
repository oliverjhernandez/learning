import Joi from 'Joi'

export const transactionSchema = Joi.object({
  type: Joi.string().required(),
  amount: Joi.number().required().min(0),
  description: Joi.string().required(),
})
