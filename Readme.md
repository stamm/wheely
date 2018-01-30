Clone repo to $GOPATH/src/github.com/stamm/wheely: `mkdir -p $GOPATH/src/github.com/stamm/wheely; git clone https://github.com/stamm/wheely.git $GOPATH/src/github.com/stamm/wheely`

Copy file `.env` from template `.env.example` and set up tokens.

Run `docker-compose up`.

Now you can do request to api: `curl -XGET -d'{"start_lat": 55.751503, "start_long": 37.623783, "end_lat": 55.753912, "end_long": 37.572263}' "http://localhost:8091/calculate"`
