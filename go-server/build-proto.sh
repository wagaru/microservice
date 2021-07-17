#! /bin/bash
docker run -v `pwd`:/defs namely/protoc-all -f schema.proto -l go