#!/bin/bash

SERVER="do"
SERVER_INBOX="/sw/_inbox/"
CUR_PATH=$(pwd)

function deploy () {
  # Triggers Deployment Script for both Frontend and Backend apps
  ssh "$SERVER" 'bash -s' < "$CUR_PATH/scripts/install.sh" "$TARGET"
  echo "==> Deployed."
}

function build_assets () {
  # Builds the frontend application for staging for production.
  echo "==> Creating build artifacts"
  cd frontend || exit 1
  npm run build || exit 1
  cd .. || exit 1
  cd api || exit 1
  go build -o release . || exit 1

  # Create Inbox
  if ssh "$SERVER" "[ ! -d $SERVER_INBOX ]"; then
    echo "==> Creating inbox on server..."
    ssh "$SERVER" "mkdir -p $SERVER_INBOX"
  fi

  echo "==> Uploading build artifacts"
  scp release "$SERVER:$SERVER_INBOX""release" || exit 1
  rm release
  cd "$CUR_PATH" || exit 1;
}

build_assets
deploy

