#!/bin/bash

export DATABASE_LOCALIZATION="$(pwd)/database"
export DATABASE_NAME="documents.db"

[[ -e "${DATABASE_LOCALIZATION}" ]] || mkdir "${DATABASE_LOCALIZATION}"

[[ -e "document-api" ]] || go build

./document-api

