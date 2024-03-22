import express, { NextFunction, Request, Response } from 'express'
import { Campground } from '../models/Campground'
import { errorCatcher } from '../utils/ErrorCatcher'
import { IReview, Review } from '../models/Reviews'
import { reviewSchema } from '../schemas/reviews'
import { ExpressError } from '../utils/ExpressError'

const router = express.Router({ mergeParams: true })

const validateReview = (req: Request, _: Response, next: NextFunction) => {
  const { error } = reviewSchema.validate(req.body)
  if (error) {
    const msg = error.details.map((el) => el.message).join(',')
    throw new ExpressError(msg, 400)
  } else {
    next()
  }
}

router.post(
  '/',
  validateReview,
  errorCatcher(async (req, res) => {
    const { id } = req.params
    const camp = await Campground.findById(id)
    if (camp) {
      const review: IReview = new Review(req.body.review)
      camp.reviews.push(review._id)
      await review.save()
      await camp.save()
      req.flash('success', 'Successful!')
      res.redirect(`/campgrounds/${camp._id}`)
    } else {
      throw new ExpressError('Unable to find campground', 404)
    }
  })
)

router.delete(
  '/:revId',
  errorCatcher(async (req, res) => {
    const { id, revId } = req.params
    const camp = await Campground.findById(id)
    const rev = await Review.findById(revId)
    if (camp && rev) {
      await Campground.findByIdAndUpdate(id, { $pull: { reviews: revId } })
      await Review.findByIdAndDelete(revId)
      req.flash('success', 'Successful!')
      res.redirect(`/campgrounds/${camp._id}`)
    } else {
      throw new ExpressError('Unable to delete review', 400)
    }
  })
)

export { router as reviewsRouter }
