import { Schema, model } from 'mongoose'
import passportLocalMongoose from 'passport-local-mongoose'

// interface IUser extends Document {
//   first_name: string
//   last_name: string
//   username: string
//   password: string
//   email: string
//   accountId: Schema.Types.ObjectId
// }

const UserSchema = new Schema(
  {
    // first_name: { type: String, required: true },
    // last_name: { type: String, required: true },
    username: { type: String, required: true },
    password: { type: String, required: true },
    email: { type: String, required: true, unique: true },
    accountId: { type: Schema.Types.ObjectId, required: true },
  },
  { timestamps: true }
)

UserSchema.plugin(passportLocalMongoose)

const UserModel = model('User', UserSchema)

export { UserModel as User }
