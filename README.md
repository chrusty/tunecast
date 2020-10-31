# tunecast
Music server with Chromecast


## Features
- Runs as a single Docker container
- Scans for media on startup
- Listens for FS events
- Media server REST API
- Play out to Chromecast
- Web UI on top of REST (browse / play folders, bar to stop / ff / rr / pause)


## Todo
- [ ] Config (music directory, port, chromecast device)
- [ ] API (OAPI-Codegen)
    - [ ] List
    - [ ] Play
    - [ ] Queue
    - [ ] Stop
    - [ ] Scan
- [ ] Server
    - main
    - handler
    - library (scan, retrieve etc)
    - storage (SQLite etc for persisting the library)
- [ ] Web UI (from API)
- [ ] Chromecast integration
    - [ ] Default player at first
    - [ ] Media links are served from a static endpoint
- [ ] Build
    - [ ] Docker
    - [ ] Test
