const request = require('request-promise')

const AUTH_ENDPOINT = `https://auth.airmap.com/realms/airmap/protocol/openid-connect/token`
const auth_endpoint_test = `https://api.airmap.com/auth/v1/anonymous/token`

async function authenticateServiceAccount() {
    const options = {
        uri: AUTH_ENDPOINT,
        method: 'POST',
        form: {
            grant_type: "password",
            client_id:'05842200-1a0e-4744-96f7-572a0edc718f',
            username: 'aagrawa2@nd.edu',
            password:'hUeDJq8_84#FJ-8',
            
        }
    }
    try {
        const response = await request(options)
        const body = JSON.parse(response)
        console.log(body.access_token)
        return body.access_token
    } catch (error) {
        console.log("throwing error from the auth module")
        throw error
    }
}






exports.authenticateServiceAccount = authenticateServiceAccount






