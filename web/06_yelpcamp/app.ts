import express, { NextFunction, Request, Response } from 'express'
import methodOverride from 'method-override'
import * as path from 'path'
import * as mg from 'mongoose'
import { Campground } from './models/Campground'
import { ExpressError } from './utils/ExpressError'
import { errorCatcher } from './utils/ErrorCatcher'
import { campgroundSchema } from './schemas/campground'
import { IReview, Review } from './models/Reviews'

// @ts-ignore
import engine from 'ejs-mate'
import { reviewSchema } from './schemas/reviews'

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

const validateCampground = (req: Request, _: Response, next: NextFunction) => {
  const { error } = campgroundSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

const validateReview = (req: Request, _: Response, next: NextFunction) => {
  const { error } = reviewSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

app.get('/campgrounds', async (_, res) => {
  const campgrounds = await Campground.find({})
  res.render('campgrounds/index', { campgrounds })
})

app.get('/campgrounds/new', async (_, res) => {
  res.render('campgrounds/new')
})

app.post(
  '/campgrounds',
  validateCampground,
  errorCatcher(async (req: Request, res: Response, _: NextFunction) => {
    const camp = new Campground(req.body.campground)
    await camp.save()
    res.redirect(`/campgrounds/${camp._id}`)
  })
)

app.get(
  '/campgrounds/:id',
  errorCatcher(async (req, res) => {
    const camp = await Campground.findById(req.params.id).populate('reviews')
    res.render('campgrounds/show', { campground: camp })
  })
)

app.get(
  '/campgrounds/:id/edit',
  errorCatcher(async (req, res, _) => {
    const camp = await Campground.findById(req.params.id)
    res.render('campgrounds/edit', { campground: camp })
  })
)

app.patch(
  '/campgrounds/:id',
  validateCampground,
  errorCatcher(async (req, res) => {
    const { id } = req.params
    const camp = await Campground.findByIdAndUpdate(id, {
      ...req.body.campground,
    })
    if (camp !== null) {
      res.redirect(`/campgrounds/${camp._id}`)
    } else {
      throw new ExpressError('Could not update campground', 400)
    }
  })
)

app.delete(
  '/campgrounds/:id',
  errorCatcher(async (req, res) => {
    const { id } = req.params
    const camp = await Campground.findByIdAndDelete(id)
    if (camp !== null) {
      res.redirect(`/campgrounds`)
    } else {
      throw new ExpressError('Could not delete campground', 400)
    }
  })
)

app.post(
  '/campgrounds/:id/reviews',
  validateReview,
  errorCatcher(async (req, res) => {
    const { id } = req.params
    const camp = await Campground.findById(id)
    if (camp) {
      const review: IReview = new Review(req.body.review)
      camp.reviews.push(review._id)
      await review.save()
      await camp.save()
      res.redirect(`/campgrounds/${camp._id}`)
    } else {
      throw new ExpressError('Unable to find campground', 404)
    }
  })
)

app.delete(
  '/campgrounds/:campId/reviews/:revId',
  errorCatcher(async (req, res) => {
    const { campId, revId } = req.params
    const camp = await Campground.findById(campId)
    const rev = await Review.findById(revId)
    if (camp && rev) {
      await Campground.findByIdAndUpdate(campId, { $pull: { reviews: revId } })
      await Review.findByIdAndDelete(revId)
      res.redirect(`/campgrounds/${camp._id}`)
    }
  })
)

app.all('*', (_, __, next) => {
  next(new ExpressError('Page Not Found', 404))
})

app.use((err: ExpressError, _: Request, res: Response, __: NextFunction) => {
  const { statusCode = 500 } = err
  if (!err.message) err.message = 'Internal Server Error'
  res.status(statusCode).render('error', { err })
})

app.listen(WEB_PORT, WEB_HOST, 3, () => {
  console.log(`Listening on port http://${WEB_HOST}:${WEB_PORT}`)
})
