import express from 'express'
import methodOverride from 'method-override'
import * as path from 'path'
import * as mg from 'mongoose'
import { CampGround } from './models/Campground'
// @ts-ignore
import engine from 'ejs-mate'

const WEB_PORT = 8080
const WEB_HOST = '0.0.0.0'
const MG_PORT = 27017
const MG_HOST = 'localhost'

mg.connect(`mongodb://${MG_HOST}:${MG_PORT}/yelp-camp`, {
  autoIndex: true,
  autoCreate: true,
  serverSelectionTimeoutMS: 2000,
})

const app = express()

app.engine('ejs', engine)
app.set('view engine', 'ejs')
app.set('views', path.join(__dirname, 'views'))

app.use(express.urlencoded({ extended: true }))
app.use(express.json())
app.use(methodOverride('_method'))

app.get('/campgrounds', async (_, res) => {
  const campgrounds = await CampGround.find({})
  res.render('campgrounds/index', { campgrounds })
})

app.get('/campgrounds/new', async (_, res) => {
  res.render('campgrounds/new')
})

app.post('/campgrounds', async (req, res) => {
  const camp = new CampGround(req.body.campground)
  await camp.save()
  res.redirect(`/campgrounds/${camp._id}`)
})

app.get('/campgrounds/:id', async (req, res) => {
  const campground = await CampGround.findById(req.params.id)
  res.render('campgrounds/show', { campground })
})

app.get('/campgrounds/:id/edit', async (req, res) => {
  const campground = await CampGround.findById(req.params.id)
  res.render('campgrounds/edit', { campground })
})

app.patch('/campgrounds/:id', async (req, res) => {
  const { id } = req.params
  const camp = await CampGround.findByIdAndUpdate(id, {
    ...req.body.campground,
  })
  if (camp !== null) {
    res.redirect(`/campgrounds/${camp._id}`)
  } else {
    throw new Error('Could not update campground')
  }
})

app.delete('/campgrounds/:id', async (req, res) => {
  const { id } = req.params
  console.log('Hello')
  const camp = await CampGround.findByIdAndDelete(id)
  if (camp !== null) {
    res.redirect(`/campgrounds`)
  } else {
    throw new Error('Could not delete campground')
  }
})

app.use((_, res) => {
  res.status(404).send('NOT FOUND')
})

app.listen(WEB_PORT, WEB_HOST, 3, () => {
  console.log(`Listening on port http://${WEB_HOST}:${WEB_PORT}`)
})
