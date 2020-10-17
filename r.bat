cd C:\Users\pushking\Go\Go\webpushUse
set DB_DSN="user=postgres password=12345 dbname=pushover sslmode=disable"
go build
webpush -r
pause