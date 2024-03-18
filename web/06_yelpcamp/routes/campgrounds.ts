import express, { NextFunction, Request, Response } from 'express'
import { Campground } from '../models/Campground'
import { ExpressError } from '../utils/ExpressError'
import { errorCatcher } from '../utils/ErrorCatcher'
import { campgroundSchema } from '../schemas/campground'

const router = express.Router({ mergeParams: true })

const validateCampground = (req: Request, _: Response, next: NextFunction) => {
  const { error } = campgroundSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

router.get('/', async (_, res) => {
  const campgrounds = await Campground.find({})
  res.render('campgrounds/index', { campgrounds })
})

router.get('/new', async (_, res) => {
  res.render('campgrounds/new')
})

router.post(
  '/',
  validateCampground,
  errorCatcher(async (req: Request, res: Response, _: NextFunction) => {
    const camp = new Campground(req.body.campground)
    await camp.save()
    res.redirect(`//${camp._id}`)
  })
)

router.get(
  '/:id',
  errorCatcher(async (req, res) => {
    const camp = await Campground.findById(req.params.id).populate('reviews')
    res.render('campgrounds/show', { campground: camp })
  })
)

router.get(
  '/:id/edit',
  errorCatcher(async (req, res, _) => {
    const camp = await Campground.findById(req.params.id)
    res.render('campgrounds/edit', { campground: camp })
  })
)

router.patch(
  '/:id',
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

router.delete(
  '/:id',
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

export { router as campgroundsRouter }
