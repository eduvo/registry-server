#!/bin/sh

rethinkdb import -f dbseed/applications.json --table registry.applications --pkey public_key --force
rethinkdb import -f dbseed/domains.json --table registry.domains --pkey identifier --force
rethinkdb import -f dbseed/identities.json --table registry.identities --pkey email --force