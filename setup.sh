#!/bin/bash
docker exec demo-spicedb-1 spicedb migrate head --datastore-engine postgres --datastore-conn-uri postgresql://myuser:mypassword@172.21.0.7:5432/mydatabase
