import express, { NextFunction, Request, Response } from 'express'
import methodOverride from 'method-override'
import * as path from 'path'
import * as mg from 'mongoose'
import { CampGround } from './models/Campground'
import { ExpressError } from './utils/ExpressError'
import { errorCatcher } from './utils/ErrorCatcher'
import { campgroundSchema } from './schemas/campground'

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

const validateSchema = (req: Request, _: Response, next: NextFunction) => {
  const { error } = campgroundSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

app.get('/campgrounds', async (_, res) => {
  const campgrounds = await CampGround.find({})
  res.render('campgrounds/index', { campgrounds })
})

app.get('/campgrounds/new', async (_, res) => {
  res.render('campgrounds/new')
})

app.post(
  '/campgrounds',
  validateSchema,
  errorCatcher(async (req: Request, res: Response, _: NextFunction) => {
    const camp = new CampGround(req.body.campground)
    await camp.save()
    res.redirect(`/campgrounds/${camp._id}`)
  })
)

app.get(
  '/campgrounds/:id',
  errorCatcher(async (req, res) => {
    const campground = await CampGround.findById(req.params.id)
    res.render('campgrounds/show', { campground })
  })
)

app.get(
  '/campgrounds/:id/edit',
  errorCatcher(async (req, res, _) => {
    const campground = await CampGround.findById(req.params.id)
    res.render('campgrounds/edit', { campground })
  })
)

app.patch(
  '/campgrounds/:id',
  validateSchema,
  errorCatcher(async (req, res) => {
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
)

app.delete(
  '/campgrounds/:id',
  errorCatcher(async (req, res) => {
    const { id } = req.params
    const camp = await CampGround.findByIdAndDelete(id)
    if (camp !== null) {
      res.redirect(`/campgrounds`)
    } else {
      throw new Error('Could not delete campground')
    }
  })
)

app.all('*', (req, res, next) => {
  next(new ExpressError('Page Not Found', 404))
})

app.use(
  (err: ExpressError, req: Request, res: Response, next: NextFunction) => {
    const { statusCode = 500 } = err
    if (!err.message) err.message = 'Internal Server Error'
    res.status(statusCode).render('error', { err })
  }
)

app.listen(WEB_PORT, WEB_HOST, 3, () => {
  console.log(`Listening on port http://${WEB_HOST}:${WEB_PORT}`)
})
