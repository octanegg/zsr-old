FORMAT: 1A
HOST: https://zsr.octane.gg

# ZSR: octane.gg API

Rocket League Esports API by Octane™

# Group Events

## Events [/events]

### List All Events [GET]

+ Parameters
  + name: RLCS X North Americal Fall Major (optional, string) - filter by name
  + tier: S (optional, string) - filter by tier
  + region: NA (optional, string) - filter by region
  + mode: 3 (optional, int) - filter by mode
  + before: 2016-12-03T00:00:00Z (optional, date) - filter before this date
  + after: 2016-12-03T00:00:00Z (optional, date) - filter after this date
  + sort: start_date (optional, string) - sort by field
  + order: asc,desc (optional, string) - order ascending or descending
  + page: 1 (optional, int) - page number
  + per_page: 20 (option, int) - results per page

- Response 200 (application/json)
    [{"_id":"5f35882d53fbbb5894b43040","name":"RLCS Season 2 World Championship","start_date":"2016-12-03T00:00:00Z","end_date":"2016-12-04T00:00:00Z","region":"INT","mode":3,"prize":{"amount":125000,"currency":"USD"},"tier":"S","stages":[{"name":"Main Event","format":"bracket-8de","region":"INT","start_date":"2016-12-03T00:00:00Z","end_date":"2016-12-04T00:00:00Z","prize":{"amount":125000,"currency":"USD"},"liquipedia":"https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/Season_2"}]}]

## Event [/events/{id}]

+ Parameters
  + id: 5f35882d53fbbb5894b43040 (required, string) - the event id

### Retrieve Event [GET]

- Response 200 (application/json)
    {"_id":"5f35882d53fbbb5894b43040","name":"RLCS Season 2 World Championship","start_date":"2016-12-03T00:00:00Z","end_date":"2016-12-04T00:00:00Z","region":"INT","mode":3,"prize":{"amount":125000,"currency":"USD"},"tier":"S","stages":[{"name":"Main Event","format":"bracket-8de","region":"INT","start_date":"2016-12-03T00:00:00Z","end_date":"2016-12-04T00:00:00Z","prize":{"amount":125000,"currency":"USD"},"liquipedia":"https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/Season_2"}]}


# Group Matches

## Matches [/matches]

### List All Matches [GET]

+ Parameters
  + event: 5f35882d53fbbb5894b43040 (optional, string) - filter by event
  + stage: 0 (optional, int) - filter by stage
  + substage: 9 (optional, int) - filter by substage
  + before: 2016-12-03T00:00:00Z (optional, date) - filter before this date
  + after: 2016-12-03T00:00:00Z (optional, date) - filter after this date
  + sort: start_date (optional, string) - sort by field
  + order: asc,desc (optional, string) - order ascending or descending
  + page: 1 (optional, int) - page number
  + per_page: 20 (option, int) - results per page

- Response 200 (application/json)
    [{"_id":"5f3ab7c96c7120e153625bc9","octane_id":"0350107","event":"5f35882d53fbbb5894b43040","stage":0,"date":"2016-12-03T00:00:00Z","format":"5","blue":{"score":3,"winner":true,"team":{"_id":"5f3d8fdd95f40596eae24583","name":"NRG Esports"}},"orange":{"score":0,"winner":false,"team":{"_id":"5f3d8fdd95f40596eae2457e","name":"Genesis"}},"mode":3,"number":7}]

## Match [/matches/{id}]

+ Parameters
  + id: 5f3ab7c96c7120e153625bc9 (required, string) - the match id

### Retrieve Match [GET]

- Response 200 (application/json)
    {"_id":"5f3ab7c96c7120e153625bc9","octane_id":"0350107","event":"5f35882d53fbbb5894b43040","stage":0,"date":"2016-12-03T00:00:00Z","format":"5","blue":{"score":3,"winner":true,"team":{"_id":"5f3d8fdd95f40596eae24583","name":"NRG Esports"}},"orange":{"score":0,"winner":false,"team":{"_id":"5f3d8fdd95f40596eae2457e","name":"Genesis"}},"mode":3,"number":7}


# Group Games

## Games [/games]

### List All Games [GET]

+ Parameters
  + event: 5f35882d53fbbb5894b43040 (optional, string) - filter by event
  + match: 5f3ab7c96c7120e153625bc9 (optional, string) - filter by match
  + sort: start_date (optional, string) - sort by field
  + order: asc,desc (optional, string) - order ascending or descending
  + page: 1 (optional, int) - page number
  + per_page: 20 (option, int) - results per page

- Response 200 (application/json)
    [{"_id":"5f3c8b6f6e8f59c2f28f9b66","octane_id":"0350107","number":2,"match":"5f3ab7c96c7120e153625bc9","event":"5f35882d53fbbb5894b43040","map":"DFH Stadium","duration":300,"mode":3,"date":"2016-12-03T00:00:00Z","blue":{"goals":0,"winner":false,"team":{"_id":"5f3d8fdd95f40596eae2457e","name":"Genesis"},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d74","tag":"Espeon"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":0,"score":155,"mvp":false,"shooting_percentage":0,"goal_participation":0,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d76","tag":"Pluto"},"stats":{"core":{"shots":1,"goals":0,"saves":0,"assists":0,"score":90,"mvp":false,"shooting_percentage":0,"goal_participation":0,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d6b","tag":"Klassux"},"stats":{"core":{"shots":0,"goals":0,"saves":1,"assists":0,"score":80,"mvp":false,"shooting_percentage":0,"goal_participation":0,"rating":0}}}]},"orange":{"goals":3,"winner":true,"team":{"_id":"5f3d8fdd95f40596eae24583","name":"NRG Esports"},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7c","tag":"Sadjunior"},"stats":{"core":{"shots":2,"goals":2,"saves":0,"assists":0,"score":320,"mvp":true,"shooting_percentage":1,"goal_participation":0.6666666865348816,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d7b","tag":"Jacob"},"stats":{"core":{"shots":2,"goals":1,"saves":0,"assists":1,"score":180,"mvp":false,"shooting_percentage":0.5,"goal_participation":0.6666666865348816,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d7a","tag":"Fireburner"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":1,"score":170,"mvp":false,"shooting_percentage":0,"goal_participation":0.3333333432674408,"rating":0}}}]}}]

## Game [/games/{id}]

+ Parameters
  + id: 5f3c8b6f6e8f59c2f28f9b66 (required, string) - the game id

### Retrieve Game [GET]

- Response 200 (application/json)
    {"_id":"5f3c8b6f6e8f59c2f28f9b66","octane_id":"0350107","number":2,"match":"5f3ab7c96c7120e153625bc9","event":"5f35882d53fbbb5894b43040","map":"DFH Stadium","duration":300,"mode":3,"date":"2016-12-03T00:00:00Z","blue":{"goals":0,"winner":false,"team":{"_id":"5f3d8fdd95f40596eae2457e","name":"Genesis"},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d74","tag":"Espeon"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":0,"score":155,"mvp":false,"shooting_percentage":0,"goal_participation":0,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d76","tag":"Pluto"},"stats":{"core":{"shots":1,"goals":0,"saves":0,"assists":0,"score":90,"mvp":false,"shooting_percentage":0,"goal_participation":0,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d6b","tag":"Klassux"},"stats":{"core":{"shots":0,"goals":0,"saves":1,"assists":0,"score":80,"mvp":false,"shooting_percentage":0,"goal_participation":0,"rating":0}}}]},"orange":{"goals":3,"winner":true,"team":{"_id":"5f3d8fdd95f40596eae24583","name":"NRG Esports"},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7c","tag":"Sadjunior"},"stats":{"core":{"shots":2,"goals":2,"saves":0,"assists":0,"score":320,"mvp":true,"shooting_percentage":1,"goal_participation":0.6666666865348816,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d7b","tag":"Jacob"},"stats":{"core":{"shots":2,"goals":1,"saves":0,"assists":1,"score":180,"mvp":false,"shooting_percentage":0.5,"goal_participation":0.6666666865348816,"rating":0}}},{"player":{"_id":"5f3d8fdd95f40596eae23d7a","tag":"Fireburner"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":1,"score":170,"mvp":false,"shooting_percentage":0,"goal_participation":0.3333333432674408,"rating":0}}}]}}    

# Group Players

## Players [/players]

### List All Players [GET]

+ Parameters
  + country: us (optional, string) - filter by country
  + tag: GarrettG (optional, string) - filter by tag
  + team: 5f3d8fdd95f40596eae24583 (optional, string) - filter by team
  + sort: start_date (optional, string) - sort by field
  + order: asc,desc (optional, string) - order ascending or descending
  + page: 1 (optional, int) - page number
  + per_page: 20 (option, int) - results per page

- Response 200 (application/json)
    [{"data":[{"_id":"5f3d8fdd95f40596eae23d6f","tag":"GarrettG","name":"Garrett Gordon","country":"us","team":"5f3d8fdd95f40596eae24583","account":{"platform":"steam","id":"76561198136523266"}}]}]

## Player [/players/{id}]

+ Parameters
  + id: 5f3d8fdd95f40596eae23d74 (required, string) - the player id

### Retrieve Player [GET]

- Response 200 (application/json)
    {"data":[{"_id":"5f3d8fdd95f40596eae23d6f","tag":"GarrettG","name":"Garrett Gordon","country":"us","team":"5f3d8fdd95f40596eae24583","account":{"platform":"steam","id":"76561198136523266"}}]}


# Group Teams

## Teams [/teams]

### List All Teams [GET]

+ Parameters
  + name: NRG Esports (optional, string) - filter by name
  + sort: start_date (optional, string) - sort by field
  + order: asc,desc (optional, string) - order ascending or descending
  + page: 1 (optional, int) - page number
  + per_page: 20 (option, int) - results per page

- Response 200 (application/json)
    [{"_id":"5f3d8fdd95f40596eae24583","name":"NRG Esports"}]

## Team [/teams/{id}]

+ Parameters
  + id: 5f3d8fdd95f40596eae23d74 (required, string) - the team id

### Retrieve Team [GET]

- Response 200 (application/json)
    {"_id":"5f3d8fdd95f40596eae24583","name":"NRG Esports"}