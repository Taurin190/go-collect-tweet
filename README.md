# go-collect-tweet
Simple command line tool to collect data from twitter api.
This tool insert twitter datas to mongodb.

This command line use mongodb with docker-compose so please execute following command with shell in this root directory.

    docker-compose up -d

To access twitter api, you need to register twitter developer and get auth keys.
Please add 'twitter.conf' like 'twitter.conf.dummy'.

To setup mongo configure, please modify 'mongo.conf' in the conf directory.