# TooGo

TooGo is a small collection of useful code snippets that save a bit of typing and copy-paste. 

## Functions
### Getenv
Read enviroment variables and use a default value if it's missing. 
`port := Getenv("PORT", ":3000")`

### FaviconHandler
Use in routing to serve the favicon. 
`router.HandleFunc("/favicon.ico", faviconHandler)`

### RoundToTwo
Takes a float64 and returns a string, rounded to two decimal points.

### TimeToDay
Takes a unix time in int format and returns the day, abbrevated to three letters.

### TimeToHour
Takes a unix timestamp in int format and returns the time in 24 hour format.