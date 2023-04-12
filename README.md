# How-To

1. `npm install` and then `npm start` within the web directory
1. Within another terminal `go run main.go`
1. Login only works with email=`test@test.com` and pw=`Password1`

# Assumptions

- Actual authentication was not needed for this exercise according to Andrew Owen. I would likely use a identity service instead of somethin in-house, hence why I didn't bother implementing a local db with hashed pw's etc.
- I did not return auth tokens or cookie to FE since the instructions only said present success or failure based on the user/pw

# Additional Notes

- NB: This is my first time trying out golang

## Follow up tasks

- prod - setup env vars
- prod - improved logging
- prod - hide metrics page
- store user information in context or reducer upon login success
- use react routes instead of different views when logged in
- call a real authentication service
- setup devcontainer
- perhaps E2E test for good credentials and bad credentials
- setup project to run FE and BE with a single command
- I would likely not use Create React App given that it is going out of fashion
