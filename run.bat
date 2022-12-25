set GeneralConfiguration=D:/test-jubelio/config/
set HOST=http://localhost
set PORT=7001
set VERSION=1.0.0
set RESOURCE_ID=chat
set PREFIX_PATH=jubelio
set DB_CONNECTION=user=postgres password=bg1603 dbname=jubelio sslmode=disable host=localhost port=5432 TimeZone=Asia/Jakarta
set DB_SCHEMA=jubelio
set DB_VIEW_CONNECTION=user=postgres password=bg1603 dbname=jubelio sslmode=disable host=localhost port=5432 TimeZone=Asia/Jakarta
set DB_VIEW_SCHEMA=jubelio
set JWT_KEY=e49aaec6-068f-4dc2-a0ce-dad1bc4f595f

go run main.go development