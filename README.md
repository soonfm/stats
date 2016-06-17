# SOON_ FM 2.0 Stats

Our SOON_ FM Stats Service. Gathers usage metrics from realtime events and provides a HTTP REST API.

## Development

1. Make your root `GOPATH`: `mkdir -p ~/stats/{src/stats,bin,pkg}`
2. Clone the repository into `$GOPATH/src/stats`: `git clone git@github.com:soonfm/stats.git $GOPATH/src/stats`
3. Install `govendor` for dependency management: https://github.com/kardianos/govendor/releases
4. Install dependencies: `cd $GOPATH/src/stats && make install_dependencies` or `govendor sync`
5. Have a 🍺

### Adding a new dependency

`govendor` handles version pinning at commit refs, branches or tags. To add a new dependency
to the repositoru simply run `govendor fethch github.com/foo/bar@1.1` for example. The
dependency and any others will be added to the `vendor.json` file in the `vendor` directory.

You can ofcourse edit the `vendor.json` file directory and then run `govendor sync`.
