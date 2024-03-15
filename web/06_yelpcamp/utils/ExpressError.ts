export class ExpressError extends Error {
  statusCode: number
  message: string
  constructor(message: string, status_code: number) {
    super()
    this.message = message
    this.statusCode = status_code
  }
}
