
const {error} = require('console')
const http = require('http')

const url = 'http://api.weatherstack.com/current?access_key=660f8617abd6d9f3788352d4b2766805&query=Bogota'

const request = http.request(url, (response) => {
    let data = ''
    
    response.on('data', (chunk) => {
        data = data + chunk.toString()
        // console.log(chunk)
    })

    response.on('end', () => {
        const body = JSON.parse(data)
        console.log(body)
    
    })
})

request.on('error', () => {
    console.log(error)
})

request.end()
