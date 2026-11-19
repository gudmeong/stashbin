# StashBin

StashBin is a simple, lightweight, free and open-source pastebin for your text file written with Go hosted on [stashbin.xyz](https://stashbin.xyz). Save and share your text to anyone for free.

## Database

MongoDB is the only backend. `DB_URI` is required (`mongodb://` or `mongodb+srv://` for Atlas); `DB_NAME` is optional and defaults to `stashbin`. There are **no SQL migrations** — the `documents` collection is created on first insert, with `slug` used as `_id` for uniqueness.

Local development with Podman:

```bash
cp .env.example .env          # set DB_URI to your local Mongo
podman compose up -d --build  # app at http://127.0.0.1:5055
```

Production: set `DB_URI` (and optionally `DB_NAME`) in the hosting platform's environment. No schema step is needed.
