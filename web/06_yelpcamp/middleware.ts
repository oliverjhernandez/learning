import { NextFunction, Response, Request } from 'express'

declare module 'express-session' {
  interface SessionData {
    returnTo?: string
  }
}

export const isLoggedId = (req: Request, res: Response, next: NextFunction) => {
  if (!req.isAuthenticated()) {
    req.session.returnTo = req.originalUrl
    req.flash('error', 'You need to login')
    return res.redirect('/login')
  }
  next()
}

export const storeReturnTo = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  if (req.session.returnTo) {
    res.locals.returnTo = req.session.returnTo
  }
  next()
}
