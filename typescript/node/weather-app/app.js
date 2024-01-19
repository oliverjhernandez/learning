const geoCode = require('./utils/geocode');
const forecast = require('./utils/forecast');
const process = require('process');


const query = process.argv[2];

if (!query) {
    console.log('Please provide an address')

} else {

    geoCode(query, (error, {latitude, longitude, place_name} = {}) => {
        if (error) {
            return console.log(error)
        }

        forecast(latitude, longitude, (error, forecastData) => {
            if (error) {
                return console.log(error)
            }

            console.log(place_name)
            console.log(forecastData)
        })
    })
}

