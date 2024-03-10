import { Document, Schema, model } from 'mongoose'

interface ITransaction extends Document {
  description: string
  date: string
  amount: number
  type: string
}

const TransactionSchema = new Schema({
  description: String,
  date: Date,
  amount: Number,
  type: String,
})

export const Transaction = model<ITransaction>('Transaction', TransactionSchema)
