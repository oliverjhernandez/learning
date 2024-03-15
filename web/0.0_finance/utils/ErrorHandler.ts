import { Request, Response, NextFunction, Handler } from 'express'

export const errorHandler = (func: Handler) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    try {
      await func(req, res, next)
    } catch (error) {
      next(error)
    }
  }
}
