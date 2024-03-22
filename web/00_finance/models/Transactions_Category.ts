import { Document, Schema, model } from 'mongoose'

interface ITransactionCategory extends Document {
  name: string
  color: string
  internalCategory: string
  accountId: Schema.Types.ObjectId
}
