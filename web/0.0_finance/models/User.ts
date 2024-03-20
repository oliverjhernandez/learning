import { Document, Schema, model } from 'mongoose'

interface IUser extends Document {
  first_name: string
  last_name: string
  email: string
  password_digest: string
  accountId: Schema.Types.ObjectId
}

const UserSchema = new Schema<IUser>(
  {
    first_name: { type: String, required: true },
    last_name: { type: String, required: true },
    email: { type: String, required: true },
    password_digest: { type: String, required: true },
    accountId: { type: Schema.Types.ObjectId, required: true },
  },
  { timestamps: true }
)

export const User = model<IUser>('User', UserSchema)
