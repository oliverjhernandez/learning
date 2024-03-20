import Joi from 'Joi'

export const accountsSchema = Joi.object({
  subtype: Joi.string().required(),
  familyId: Joi.string().required(),
  name: Joi.string().required(),
  balance: Joi.number().required(),
  currency: Joi.string().required(),
  classification: Joi.string().required(),
})
