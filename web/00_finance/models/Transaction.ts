import { Document, Schema, model } from 'mongoose'

interface ITransaction extends Document {
  description: string
  date: Date
  accountId: Schema.Types.ObjectId // TODO: maybe a better type?
  amount: number
  currency: number // TODO: better typing
  type: string // TODO: better typing
  categoryId: Schema.Types.ObjectId // TODO: maybe a better type?
  notes: string
  // excluded: boolean
}

const TransactionSchema = new Schema(
  {
    description: { type: String, required: true },
    date: { type: Date, required: true },
    accountId: { type: Schema.Types.ObjectId, required: true },
    amount: { type: Number, required: true },
    currency: { type: String, required: true },
    type: { type: String, required: true },
    categoryId: { type: Schema.Types.ObjectId, required: true },
    notes: { type: String, required: true },
    // excluded: { type: Boolean, required: true },
  },
  { timestamps: true }
)

export const Transaction = model<ITransaction>('Transaction', TransactionSchema)
