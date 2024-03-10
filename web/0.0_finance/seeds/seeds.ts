import * as mg from 'mongoose'
import { Transaction } from '../models/Transactions'
import { transactions } from './data'

const MONGO_PORT = 27017
const HOST = '0.0.0.0'

mg.connect(`mongodb://${HOST}:${MONGO_PORT}/finance`, {
  autoIndex: true,
  autoCreate: true,
  serverSelectionTimeoutMS: 2000,
})

const seedDB = async () => {
  await Transaction.deleteMany({})

  for (const tr of transactions) {
    const transaction = new Transaction({
      description: tr.description,
      date: new Date(tr.date),
      amount: tr.amount,
      type: tr.type,
    })
    console.log(`saving doc`)
    await transaction.save()
  }
}

seedDB().then(() => {
  console.log('DB seeded. Closing now.')
  mg.connection.close()
})
