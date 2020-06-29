// Bootstrap our dependencies.
const grpc = require('grpc')
const auth = require('../auth')
const proto = require('../proto')

const util = require('util')

api_key = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxCT2drZFBRUzQzazZxa3NncUJwSk1peTBSMzAwIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnxQb0w4bFk5aFA1YjNwYWN3cFIwSjVDbXBua2U4Iiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfGdvR2F5UWFGbkJkTWQ1ZjJPWWs4eEZrMkJad2wiLCJpYXQiOjE1OTMxMTU5MTF9.chgVxepo7RJlaD329Fdp5xP5w7xw1M6L5QjdQGSW5IY.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxrM3Z4eERKVEU3dmtPeEhlNFkybzlJT1IwRWdrIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnw2eTR5bWx6Y2FFRE9tQlN5UTBNUGRjM215YndRIiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfGdvR2F5UWFGbkJkTWQ1ZjJPWWs4eEZrMkJad2wiLCJpYXQiOjE1OTIyODA3MjR9.xwoHNzvK1woE9Ur-7WQM_Pn0ipZl8k180W5FWJ5Dt4w'
client_id = '9c819ed1-45ce-4631-8934-c42e90ae0d12'
// Configure your env with your client id and client secret
const API_ENDPOINT = process.env.API_ENDPOINT || 'api.airmap.com:443'
const CLIENT_ID =  client_id
const CLIENT_SECRET =  api_key

// The starting coordinate for the simulation. Defaults to Baja California
const COORDINATE_LATITUDE = process.env.COORDINATE_LATITUDE | -87.92753219604492
const COORDINATE_LONGITUDE = process.env.COORDINATE_LONGITUDE || 41.99152230528149

function connectProvider(token) {

  const telemetry = proto.descriptors.telemetry
  const collector = new telemetry.Collector(API_ENDPOINT, grpc.credentials.createSsl())

  const metadata = new grpc.Metadata()
  metadata.add("authorization", "Bearer " + token)

  const client = collector.ConnectProvider(metadata)

  client.on('data', function (response) {
    console.log(response)
    switch (response.details) {
      case "ack":
        console.log('received acknowledgement (ACK) from collector: ', response.ack)
        break
      case "nack":
        console.log('WARNING: received negative acknowledgement (NACK) from collector: ', response.nack)
        break
      case "status":
        console.log('received status from collector: ', response.status.level, response.status.message)
        break
    }
  })

  client.on('end', function () {
    console.log('client connection to collector ended')
  })

  client.on('error', function (error) {
    throw error
  })

  return client
}

function runSimulation(client) {

  var lat = COORDINATE_LATITUDE
  var lng = COORDINATE_LONGITUDE

  setInterval(function () {
    // Advance the coordinate position at each tick
    lng += 0.001
    lat += 0.001

    const report = buildReport(lat, lng)
    console.log(util.inspect(report, {showHidden: false, depth: null}))
    client.write({report: report})

  }, 1000)
}

function buildReport(lat, lng) {

  // the current time in milliseconds
  const now = new Date()
  const seconds = now / 1000
  const nanos = (now % 1000) * 1e6

  const observed = {
    seconds: seconds,
    nanos: nanos
  }

  const position = {
    absolute: {
      coordinate: {
        latitude: {
          value: lat
        },
        longitude: {
          value: lng
        }
      },
      altitude: {
        height: {
          value: 50
        },
        reference: "REFERENCE_ELLIPSOID"
      }
    }
  }

  const report = {
    observed: observed,
    identities: [
      {
        imei: {
          as_string: "994085654109130"
        },
      }
    ],
    spatial: {
      position: position
    }
  }

  return report
}

// Login using the client credentials,
// connect to the collector with a token,
// and run the simulation
auth.authenticateServiceAccount()
  .then(connectProvider)
  .then(runSimulation)
  .catch(console.log)


