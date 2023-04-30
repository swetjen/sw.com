#!/bin/bash

TARGET="production"
SERVER_INBOX="/sw/_inbox"

SERVICE="sw-production.service"

BASE_FOLDER="/sw"
RELEASE_FOLDER="$BASE_FOLDER/releases"

EPOC=$(date +%s)
# Slice four digits near the end of epoc for our sub-version number.
VERSION=${EPOC:5:4}
RELEASE_NAME="release_$TARGET.$(date '+%Y.%m.%d').$VERSION"

echo "==> Starting deployment..."

# If first deployment, bootstrap the requirements
if [[ ! -d "$RELEASE_FOLDER" ]]; then
  mkdir "$RELEASE_FOLDER" -p
  cd "$RELEASE_FOLDER" || exit 1
fi

  cd "$RELEASE_FOLDER" || exit 1
  # Rename build assets
  mv "$SERVER_INBOX/release" "$RELEASE_FOLDER/$RELEASE_NAME" || exit 1

  # Make Symlink
  cd $RELEASE_FOLDER || exit 1
  if [ -f "$TARGET" ]; then
    rm "$TARGET"
  fi

  ln -s "$RELEASE_FOLDER/$RELEASE_NAME" $TARGET || exit 1
  sudo systemctl restart "$SERVICE"

  exit 0

fi
