#!/bin/bash
revive -config ./revive.toml -formatter friendly -exclude ./vendor/... ./...
