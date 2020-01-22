# Leaderboarder

Build a leaderboard for any application you can think of in under 5 minutes

## Features

- Add and retrieve leaderboard
- Support for local file saving to keep application completely self-contained!

## Endpoints
### [GET] `{url}/` 
Returns the list of scores in JSON format in descending order by score

### [POST] `{url}/submit` 

Creates a new leaderboard entry

**Params**
- name - Name for entry
- score - Score of entry
- metadata - Metadata for entry, plug anything into this though JSON is recommended for later consumption

## TODO
- [ ] TLS support for encrypting connections
- [ ] Filtering leaderboard retrieval via basic query language
- [ ] Optional leaderboard front-end support!  Just toggle to serve your leaderboard via a built-in web server.
- [ ] Support for multiple back-ends for saving data (it's extensible today)
- [ ] Schema definitions for metadata acceptance
