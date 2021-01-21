# count-jobs-api

Sample API based on a crawler that gets the number of jobs from indeed.

## Usage

```
git clone https://github.com/lucasnevespereira/count-jobs-api
```

```
go run main.go
```

Visit [localhost:3000](http://localhost:3000)

## Options

There are currently 4 countries handled by this API:

- fr (France 🇫🇷)
- uk (United Kingdom 🇬🇧)
- usa (United States 🇺🇸)
- pt (Portugal 🇵🇹)

## Example Request

```
http://localhost:3000/api?term=Javascript&location=Paris&country=fr
```
