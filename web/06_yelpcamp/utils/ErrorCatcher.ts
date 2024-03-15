import { NextFunction, Response, Request, Handler } from 'express'

// ts-ignore
export const errorCatcher = (func: Handler) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    try {
      // TODO: Improve this warning
      await func(req, res, next)
    } catch (error) {
      next(error)
    }
  }
}
