package constanta

import "time"

const DefaultDateFormat = `2006-01-02 15:04:05`

const RequestPOST = "POST"
const RequestGET = "GET"
const RequestDELETE = "DELETE"
const RequestPUT = "PUT"

const RequestSuccess = "Request Success"
const RequestFailed = "Request Failed"

const DescIncorrectFormat = "incorrect format"
const DescLoginFailed = "login failed"
const DescRedisDeleteDataFailed = "failed to delete data in redis by key"
const DescUnauthorized = "access forbidden"
const DescUpdateFailed = "failed to update"
const DescEmptyField = "field is empty"
const DescDataNotFound = "data not found"
const DescActivationFailed = "failed to activate your account"
const DescInvalidPassword = "password invalid"

const Time8Hour = 8 * time.Hour
