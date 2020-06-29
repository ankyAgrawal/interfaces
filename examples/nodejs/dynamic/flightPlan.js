const auth = require('./auth')
const request = require('request-promise')
//var GeoJSON = require('geojson')


//polygon_json = GeoJSON.parse(data,{'Polygon':'polygon'});
api_key = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxYWHZ5OFBuRnhNN1FacXVFcVFkbXBJejJiV2JLIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnxZZ3ZQQU1LSFhhZFdSUXU2MG9lb25UNEtKNk9tIiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfE1XSkJiWjJoWTdnWHkydE9xeGJFbVU5TVc2SmciLCJpYXQiOjE1OTMxMjgzMDl9.nRJ244ApZENSZrhzcVw_QdCwuo7P2IPv77R5d_a_r60"

pilot_id=""
flight_id=""
async function getPilot(token){


    const options = {
        uri: "https://api.airmap.com/pilot/v2/profile",
        method: "GET",
        headers: {
            "Authorization": "Bearer"+ " "+token,
            "X-API-Key": this.api_key
        }
        
    }
    try {
        const response = await request(options)
        const body = JSON.parse(response)
        
        pilot_id = body.data.id.toString()
        console.log(body.data.id)
        return pilot_id
    } catch (error) {
        console.log("throwing error from the flight module")
        throw error
    }

}

async function  createFligh(token){
    
    pilot_id = await getPilot(token)
    start_time = new Date('29 June 2020 14:00 UTC');
    end_time = new Date('30 June 2020 14:30 UTC')
    search_area_geoJson = 
    
    {
        // "type": "FeatureCollection",
        // "features": [
          // {
          //   "type": "Feature",
          //   "properties": {},
            // "geometry": {
              type: "Polygon",
              coordinates: [
                [
                  [
                    -87.93354034423828,
                    41.99024642991162
                  ],
                  [
                    -87.88770675659178,
                    41.995987667682414
                  ],
                  [
                    -87.89834976196289,
                    41.96051129429777
                  ],
                  [
                    -87.93439865112303,
                    41.98475987441191
                  ],
                  [
                    -87.93354034423828,
                    41.99024642991162
                  ]
                  
                ]
              ]
            //}
          //}
        // ]
      }
    
    const flightdata= {
        "rulesets": [
          "usa_part_107"
        ],
        "start_time": start_time.toISOString(),
        "end_time": end_time.toISOString(),
        "max_altitude_agl": "60.96",
        "takeoff_latitude": "33.85505651142062",
        "takeoff_longitude": "-118.37099075317383",
        "buffer": "1",
        "geometry": search_area_geoJson,
        "pilot_id": String(pilot_id)
      }

    const options = {
        uri: 'https://api.airmap.com/flight/v2/plan',
        method: 'POST',
        headers: {
            'content-type': 'application/json',
            'Authorization': "Bearer" + " "+token,
            'X-API-Key': this.api_key
        },
        body:flightdata,
        json:true
    }
    try {
        console.log(JSON.stringify(flightdata))
        
        const response = await request(options)
        //const body = JSON.parse(response)
        flight_id = response.data.id
        console.log("Flight_id is",response.data.id)
        return flight_id
    } catch (error) {
        console.log("throwing error from the flight module")
        throw error
    }

}


async function submitFlight(token){

  generate_flight_id = await createFligh(token)
  const options = {
    uri: 'https://api.airmap.com/flight/v2/plan/'+generate_flight_id+'/submit',
    method: 'POST',
    headers: {
        'content-type': 'application/json',
        'Authorization': "Bearer" + " "+token,
        'X-API-Key': this.api_key
    },
    body:{"public":true},
    json:true
}

try{
  const response = await request(options)
  console.log(response)
  return response.data.flight_id
}
catch(error){
  console.log("throwing error from the flight module")
        throw error
}




}


async function deleteflight(token){
  flight_Plan_id="flight_plan|DnX4okmUzKqWPzC4oXpEnHxzzdNe"
  var options = {
    method: 'DELETE',
    url: 'https://api.airmap.com/flight/v2/plan/'+flight_Plan_id,
    headers: {
      'content-type': 'application/json',
      'x-api-key': api_key,
      'Authorization': "Bearer" + " "+token,
    }
  };

  try{
    const response = await request(options)
    console.log(response)
    //return response.data.flight_id


  }
  catch(error){
    console.log(" error while deleting a flight plan")
    throw error;
  }
}

async function getFlight(){

  var options = {
    method: 'GET',
    url: 'https://api.airmap.com/flight/v2/',
    headers: {accept: 'application/json', 'x-api-key': api_key}
  };

  try{
    const response = await request(options)
    console.log(response)
    //return response.data.flight_id


  }
  catch(error){
    console.log(" error while deleting a flight plan")
    throw error;
  }

  

}


//getFlight().catch(console.log)

//auth.authenticateServiceAccount().then(submitFlight).catch(console.log)

auth.authenticateServiceAccount().then(deleteflight).catch(console.log)
