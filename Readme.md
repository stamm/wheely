Clone repo to $GOPATH/src/github.com/stamm/wheely: `mkdir -p $GOPATH/src/github.com/stamm/wheely; git clone https://github.com/stamm/wheely.git $GOPATH/src/github.com/stamm/wheely`

Copy file `.env` from template `.env.example` and set up tokens.

Run `docker-compose up`.
In case you are using MacOs, you need to add file sharing in Docker settings: ~/code/go/src/github.com/stamm/wheely/cfg/

Now you can do request to api: `curl -XGET -d'{"start_lat": 55.751503, "start_long": 37.623783, "end_lat": 55.753912, "end_long": 37.572263}' "http://localhost:8091/calculate"`

I use consul for service discovery.
For api I use toolkit go-kit. Because it is not a framework, but have pretty architecture.
For transport I use json via http. But I can easily change it to grpc between api's (what I prefer, go kit really helps with it).
For a quick running I use docker-compose.
