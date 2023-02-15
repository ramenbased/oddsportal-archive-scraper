#Oddsportal History Scraper

Barebone tool to get moneyline odds from oddsportal dot com. Basically just a [chromedp](https://github.com/chromedp/chromedp) copy paste. Does not scrape bookmaker specific data. Works on my machine. Might work on windows. 

I've seen most repos got rekt and are abandoned so i'm publishing this... 

Get [Go](https://go.dev).

```bash
git clone https://github.com/ramenbased/oddsportal-archive-scraper
```

```bash
go build
./oddsportalhistory -h

##prints help
```
Use flags to instruct it.

Example.. scraping the entire 2022 season of MLB:

```bash
./oddsportalhistory -u "https://www.oddsportal.com/baseball/usa/mlb-2022/results/#/page/" -p 55 -s "MLB2022-"
```

Additional tool for parsing "per page" data into a single.json of one year containing all games (for now only baseball because it based ball).

```bash
cd combinejsons
go build
./combinejsons -h

##prints help
```
Example of data:

```bash
...
                "date-start-base": 1667433780,
                "date-start-timestamp": 1667433780,
                "result": "0:5",
                "homeResult": "0",
                "awayResult": "5",
                "home-winner": "lost",
                "away-winner": "win",
                "info": [
                    {
                        "name": "4th leg.",
                        "sort_key": 0
                    },
                    {
                        "name": "Series tied 2-2.",
                        "sort_key": 1
                    }
                ],
                "partialresult": "0:0, 0:0, 0:0, 0:0, 0:5, 0:0, 0:0, 0:0, 0:0",

...

    "odds": [
                    {
                        "avgOdds": 1.87,
                        "bettingTypeId": 3,
                        "eventId": 6314078,
                        "maxOdds": 1.92,
                        "outcomeResultId": 2,
                        "scopeId": 1,
                        "outcomeId": "60m2ux2rrhkx0xg2bdj",
                        "maxOddsProviderId": 5,
                        "active": true
                    },
                    {
                        "avgOdds": 1.99,
                        "bettingTypeId": 3,
                        "eventId": 6314078,
                        "maxOdds": 2.09,
                        "outcomeResultId": 1,
                        "scopeId": 1,
                        "outcomeId": "60m2ux2rrhkx0xg2bdi",
                        "maxOddsProviderId": 381,
                        "active": true
                    }
                ],
                "name": "Philadelphia Phillies - Houston Astros",
...

```

Example of using scraped data:

![Basedball](https://i.imgur.com/dKwvJAj.png)
