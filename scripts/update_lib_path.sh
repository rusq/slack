#!/bin/zsh

OLD="github.com/slack-go/slack"
NEW="github.com/rusq/slack"

find . -type f -name "*.go" -print0 | while IFS= read -r -d '' file; do
  sed -i '' "s|$OLD|$NEW|g" "$file"
done
