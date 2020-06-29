

token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpZXV1WVI2MTJKYWg0YjR2VVFpZHN6Vk9Mc0Q2R0YzS1hhZk0wb09DVDVvIn0.eyJleHAiOjE1OTMzMjc0MDYsImlhdCI6MTU5MzMwOTQwNiwianRpIjoiYTMyZTFiZjAtNTI3Mi00MGMyLWJkYmItNWMzZmY4MDRlZjdiIiwiaXNzIjoiaHR0cHM6Ly9hdXRoLmFpcm1hcC5jb20vcmVhbG1zL2Fpcm1hcCIsImF1ZCI6ImFpcm1hcC1hcGlzIiwic3ViIjoia2N8ZWY2ODYyMGYtNjY0NC00ZGY1LTllMDktODRhZTMxZjVlZmE1IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiMDU4NDIyMDAtMWEwZS00NzQ0LTk2ZjctNTcyYTBlZGM3MThmIiwic2Vzc2lvbl9zdGF0ZSI6IjFhYzhlMjk5LTM0ZGItNDY2OC1iMjRlLWJlNjkxNGYzYWFlNCIsImFjciI6IjEiLCJzY29wZSI6ImVtYWlsIGFtLWFwaSIsImtjX2lkIjoiZWY2ODYyMGYtNjY0NC00ZGY1LTllMDktODRhZTMxZjVlZmE1IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImVtYWlsIjoiYWFncmF3YTJAbmQuZWR1In0.Xfi4BYAf0lxnD6NA5szIyn7Odfy7jU57lg0dRZL3aoJcXgJkXJ7BHO2YPqP5yqkDPsnrE5Wr8zRuKUU_Oe5FMKiDp6XA_SuRMx3Mg-jK_9-Qw5P_1QutsOBStWYH41PrxHGHzKK7wPMNrYbY9D13aEYc57ZJyBhs6cb3SH2b0dDUAb0nssWWJoMOnOIjXeIEf3kJ9GGcgrRSjWDmUHCI7EEjVLXzPC19Sgoa4PCAoyiLPIOmLv2KcNuaJG80ad-k6hF3Mt6EzL6qqSjm9YiUpO3HtqRB5LoVLWD7kz4tytENxqdkk3MI_bI30GviY8iyYeMJmd0-h_xrgK4KN1CF3Q"
//flight_id = "flight_plan|Oz42w7pT0qW46RswBDl4JUW4AaE6"
flight_id=""
api_key = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxCT2drZFBRUzQzazZxa3NncUJwSk1peTBSMzAwIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnxQb0w4bFk5aFA1YjNwYWN3cFIwSjVDbXBua2U4Iiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfGdvR2F5UWFGbkJkTWQ1ZjJPWWs4eEZrMkJad2wiLCJpYXQiOjE1OTMxMTU5MTF9.chgVxepo7RJlaD329Fdp5xP5w7xw1M6L5QjdQGSW5IY.eyJjcmVkZW50aWFsX2lkIjoiY3JlZGVudGlhbHxrM3Z4eERKVEU3dmtPeEhlNFkybzlJT1IwRWdrIiwiYXBwbGljYXRpb25faWQiOiJhcHBsaWNhdGlvbnw2eTR5bWx6Y2FFRE9tQlN5UTBNUGRjM215YndRIiwib3JnYW5pemF0aW9uX2lkIjoiZGV2ZWxvcGVyfGdvR2F5UWFGbkJkTWQ1ZjJPWWs4eEZrMkJad2wiLCJpYXQiOjE1OTIyODA3MjR9.xwoHNzvK1woE9Ur-7WQM_Pn0ipZl8k180W5FWJ5Dt4w'
var lat = 34.0248
var lon = -118.44267
var buffer = 1000

var data = {
    latitude: lat,
    longitude: lon,
    buffer: buffer
  }



function sendapi(){
  $.ajax({
    url: "https://api.airmap.com/flight/v2/point",
    headers: {
      "X-API-KEY": api_key,
      "Authorization": "Bearer " + token,
      "Content-Type": "application/json; charset=utf-8"
    },
    type: "post",
    data: JSON.stringify(data),
    dataType: "json",
    success: function(res) {
     

    flight_id = res.data.id

      console.log("Flight created: " + flight_id)

      listenTraffic()

    },
    error: function(err) {
      alert(err.name + ":" + err.message)
    }
  })
}

// console.log(token)
function listenTraffic() {

    // SOME MQTT CLIENTS require a "/mqtt" at the end of the URL
    var client = new Paho.MQTT.Client('mqtt.airmap.com', 8884, "c-mqtt-cli")
    
    client.onMessageArrived = onMessageArrived
    
    function onMessageArrived(message) {
      console.log("onMessageArrived:" + message.payloadString);
    }
    
    client.connect({
      onSuccess: onConnect,
      userName: flight_id,
      password: token,
      useSSL: true
    })
    
    function onConnect() {
    
      console.log("Connected")
    
      var topic_alert = 'uav/traffic/alert/' + flight_id
    
      //var topic_sa = 'uav/traffic/sa/' + flight_id
    
      client.subscribe(topic_alert)
    
    }
    
    }




function auth(){
    console.log("sending authentication reqyest")
        $.ajax({
        uri: "https://auth.airmap.com/realms/airmap/protocol/openid-connect/token",
        type: "post",
        crossDomain: true,
        data: {
            grant_type: "password",
            client_id:'05842200-1a0e-4744-96f7-572a0edc718f',
            username: 'aagrawa2@nd.edu',
            password:'hUeDJq8_84#FJ-8',
         
        },
        dataType: "json",
        complete: function(result){
            alert(result)
        },
        
        
        success:function(res){
            console.log(res)

        },
        error:function(error){
            console.log(error)
        }
    })
}


// authenticate()