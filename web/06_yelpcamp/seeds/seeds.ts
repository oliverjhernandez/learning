import * as mongoose from 'mongoose'
import { CampGround } from '../models/campground'
import { cities, adjectives, nouns } from './data'

const MONGO_PORT = 27017
const HOST = '0.0.0.0'

mongoose.connect(`mongodb://${HOST}:${MONGO_PORT}/yelp-camp`, {
  autoIndex: true,
  autoCreate: true,
  serverSelectionTimeoutMS: 2000,
})

const randomName = () => {
  const first = adjectives[Math.floor(Math.random() * adjectives.length)]
  const second = nouns[Math.floor(Math.random() * nouns.length)]
  return `${first}_${second}`
}

const seedDB = async () => {
  await CampGround.deleteMany({})
  for (let c = 0; c < 50; c++) {
    const random1000 = Math.floor(Math.random() * 1000)
    const camp = new CampGround({
      location: `${cities[random1000].city}, ${cities[random1000].state}`,
      title: randomName(),
      price: Math.floor(Math.random() * 200) + 10,
      image: 'https://source.unsplash.com/collection/483251',
      description: `
        Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit 
        enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation 
        amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud 
        nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor 
        minim nulla est proident.
      `,
    })
    await camp.save()
  }
}

seedDB().then(() => {
  console.log('DB seeded. Closing now.')
  mongoose.connection.close()
})
