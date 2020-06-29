// Bootstrap our dependencies.
const grpc = require('grpc')
const auth = require('../auth')
const proto = require('../proto')
api_key = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxLYk5OQmcyRlBMWE5nTGhuTk44YW9GQTBZRHZrIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnw2eTR5bWx6Y2FFRE9tQlN5UTBNUGRjM215YndRIiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfGdvR2F5UWFGbkJkTWQ1ZjJPWWs4eEZrMkJad2wiLCJpYXQiOjE1OTMxMTU2ODB9.adJZ6R-D73zPyHsu_tZYwUJ9LuNQdxbdO1MYvkQKcxM.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxrM3Z4eERKVEU3dmtPeEhlNFkybzlJT1IwRWdrIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnw2eTR5bWx6Y2FFRE9tQlN5UTBNUGRjM215YndRIiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfGdvR2F5UWFGbkJkTWQ1ZjJPWWs4eEZrMkJad2wiLCJpYXQiOjE1OTIyODA3MjR9.xwoHNzvK1woE9Ur-7WQM_Pn0ipZl8k180W5FWJ5Dt4w'
client_id = '406c4c14-34d6-4be7-8624-54562174962b'
// Configure your env with your client id and client secret
const API_ENDPOINT = process.env.API_ENDPOINT || 'api.airmap.com:443'
const CLIENT_ID = client_id
const CLIENT_SECRET = api_key

function connectProcessor(token) {

  const tracking = proto.descriptors.tracking
  const collector = new tracking.Collector(API_ENDPOINT, grpc.credentials.createSsl())

  const metadata = new grpc.Metadata()
  metadata.add("authorization", "Bearer " + token)
  
  const client = collector.ConnectProcessor(metadata)

  client.on('data', function (response) {
    switch (response.details) {
      case "batch":
        for (let i in response.batch.tracks) {
          let track = response.batch.tracks[i]
          console.log(track.identities)
          console.log(track.position.absolute)
        }
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
}

// Login using the client credentials,
// and connect to the collector with a token
auth.authenticateServiceAccount()
  .then(connectProcessor)
  .catch(console.log)
