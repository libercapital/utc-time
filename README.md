# Welcome to utctime module 👋

Module for handling datetime in JSON messages in RFC3339 format

## Install

```go
go get gitlab.com/bavatech/architecture/software/libs/go-modules/utctime.git
```

## Usage

Inside the module there is a new type UTCTime

UTCTime must be used in place of datetime fields as they support direct JSON conversion

### Explicit Statement

```go
type Request struct {
	EndTime   utctime.UTCTime `json:"end_time"`
	StartTime utctime.UTCTime `json:"start_time"`
	EndDate   utctime.UTCDate `json:"end_date"`
	StartDate utctime.UTCDate `json:"start_date"`
}
```

### Locate an utc time to GMT-03:00 America/São Paulo

```go
utclocalized = utctime.ToSPTimeZone()
```

### Conversions and Casts samples

```go
// utctime to string formated yyyy-mm-ddThh-mm-ssZ
formatedstr := utctime.String()

// utctime to string formated yyyy-mm-ddThh-mm-ss.000Z
formatedstr := utctime.StringMiliseconds()

// string to utctime
utctime = utctime.ParseToUTCTime(datetimeString)

// time.Time to UTCTime
utctime = UTCTime(time)

// UTCTime to time.Time
timet = time.Time(utctime)

// **** DATE ****
date = UTCDate(time.Now())
// utctime to string formated yyyy-mm-dd
formatedstr := date.String()

// string to UTCDate
utcdate = utctime.ParseToUTCDate(time.Now().String())

// time.Time to UTCDate
utcdate = UTCDate(time)

// UTCDate to time.Time
timet = time.Time(utcdate)
```

## Author

👤 **Eduardo Mello**

- Gitlab: [@eduardo.mello@bavabank.com](https://gitlab.com/eduardo.mello)

## Contributors

👤 **Vinícius Deuner**

- Gitlab: [@vinicius.deuner@bavabank.com](https://gitlab.com/vinicius.deuner)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
