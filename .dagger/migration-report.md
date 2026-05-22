# Migration Report

## Root module requires explicit loading

The root `dagger.json` is still a valid module, but it must be loaded explicitly.

- **This works**: `dagger -m . call --help`
- **This no longer works**: `dagger call --help`

ACTION: If your scripts rely on implicit loading of the root module, change them to use explicit loading.

## mod-test requires explicit loading

`mod-test` is still a valid module, but it must be loaded explicitly.

- **This works**: `dagger -m mod-test call --help`
- **This no longer works**: `cd mod-test; dagger call --help`

ACTION: If your scripts rely on implicit loading of `mod-test`, change them to use explicit loading.

## polyfill requires explicit loading

`polyfill` is still a valid module, but it must be loaded explicitly.

- **This works**: `dagger -m polyfill call --help`
- **This no longer works**: `cd polyfill; dagger call --help`

ACTION: If your scripts rely on implicit loading of `polyfill`, change them to use explicit loading.
