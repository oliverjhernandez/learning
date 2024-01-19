
const request = require('postman-request');

const forecast = (lon, lat, callback) => {
    const url = `http://api.weatherstack.com/current?access_key=660f8617abd6d9f3788352d4b2766805&query=${lat},${lon}`;
    request({url, json: true}, (error, { body }) => {
        if (error) {
            callback('Unable to connect to forecast services', undefined)
        } else if (body.error) {
            callback(body.error.info, undefined)
        } else {
            const description = body.current.weather_descriptions[0]
            const temperature = body.current.temperature
            const feel = body.current.feelslike
            callback(undefined, `${description}: It is currently ${temperature} C and it feels like ${feel} C`)
        }
    })
}


module.exports = forecast
