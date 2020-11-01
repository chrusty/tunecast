# tunecast
Music server with Chromecast


## Features
- Runs as a single Docker container
- Scans for media on startup
- Listens for FS events
- Media server REST API
- Play out to Chromecast
- Web UI on top of REST (browse / play folders, bar to stop / ff / rr / pause)


## Config
- `CHROMECAST_ADDRESS` ("chromecast"): Address of the Chromecast
- `LIBRARY_PATH` ("/media"): Where to find your media files
- `DB_DISABLED` ("true"): Disable the DB (run in-memory)
- `HTTP_LISTEN_ADDRESS` (":8080"): Where to host the API & web services
- `LOGGING_LEVEL` ("INFO"): Log verbosity ["DEBUG", "INFO", "WARNING", "ERROR"]


## Todo
- [x] Config (music directory, port, chromecast device)
- [x] API (OAPI-Codegen)
    - [x] List
    - [ ] Play
    - [ ] Queue
    - [ ] Stop
    - [ ] Scan
- [x] Server
    - [x] main
        - [x] Config
        - [x] Logger
        - [x] HTTP router
        - [x] Listen
    - [x] Handler
    - [ ] Browser
    - [x] MediaShare (/media)
    - [x] Library (scan, retrieve etc)
        - [x] Storage
            - [x] SQLite
- [ ] Web UI (from API)
- [ ] Chromecast integration
    - [ ] Default player at first
    - [ ] Media links are served from a static endpoint
- [x] Makefile
    - [x] Local
    - [ ] Docker
    - [x] Test
