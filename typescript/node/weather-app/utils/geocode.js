const request = require('postman-request');

const geoCode = (address, callback) => {
   const url = `https://api.mapbox.com/geocoding/v5/mapbox.places/${encodeURIComponent(address)}.json?access_token=pk.eyJ1Ijoib2xpdmVyamhlcm5hbmRleiIsImEiOiJja2o5OWRoczEwZ3BiMnVwNzF5azE4Y3oyIn0.pl6o-cyqFOEDWeEAKexW9A&limit=1`

    request({url, json: true}, (error, { body }) => {
        if (error) {
            callback('Unable to connect to location services.', undefined);
        } else if (body.message) {
            callback(body.message, undefined)
        } else {
            callback(undefined, {
                latitude: body.features[0].center[0],
                longitude: body.features[0].center[1],
                place_name: body.features[0].place_name
            });
        }
    })
}

module.exports = geoCode;
