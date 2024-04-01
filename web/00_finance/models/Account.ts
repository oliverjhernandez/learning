import { Document, Schema, model } from 'mongoose'
import {
  TAccountTypes,
  TBankAccountTypes,
  TClassification,
  TCreditCardTypes,
  TCurrency,
} from '../definitions/accounts'

interface IAccount extends Document {
  subtype: TBankAccountTypes | TCreditCardTypes
  familyId: Schema.Types.ObjectId
  name: string
  type: TAccountTypes
  balance: number
  currency: TCurrency
  classification: TClassification
}

const AccountSchema = new Schema<IAccount>(
  {
    subtype: { type: String, required: true },
    familyId: { type: Schema.Types.ObjectId, required: true },
    name: { type: String, required: true },
    type: { type: String, required: true },
    balance: { type: Number, required: true },
    currency: { type: String, required: true },
    classification: { type: String, required: true },
  },
  { timestamps: true }
)

export const Account = model<IAccount>('Account', AccountSchema)
