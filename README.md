# count-jobs-api

Sample API based on a scrapper that gets the number of jobs from indeed.

## Usage

```
make serve
```

Visit [localhost:5001](http://localhost:5001)

## Options

There are currently 4 countries handled by this API:

- fr (France 🇫🇷)
- uk (United Kingdom 🇬🇧)
- usa (United States 🇺🇸)
- pt (Portugal 🇵🇹)

## Example Requests

```
http://localhost:5001/api?term=Javascript&location=Paris&country=fr
```
