#!/bin/sh
set -euxo

generate_spec() {
  (
    cd server
    # TODO: Enable
    #composer update
    #composer install --no-dev

    composer exec merge-specs -- \
      --first-status-code \
      --core ./core/openapi.json \
      --merged ../openapi.json \
      ./apps/comments/openapi.json \
      ./apps/dashboard/openapi.json \
      ./apps/dav/openapi.json \
      ./apps/files/openapi.json \
      ./apps/files_external/openapi.json \
      ./apps/files_reminders/openapi.json \
      ./apps/files_sharing/openapi.json \
      ./apps/files_trashbin/openapi.json \
      ./apps/files_versions/openapi.json \
      ./apps/provisioning_api/openapi.json \
      ./apps/settings/openapi.json \
      ./apps/sharebymail/openapi.json \
      ./apps/theming/openapi.json \
      ./apps/updatenotification/openapi.json \
      ./apps/user_status/openapi.json \
      ./apps/weather_status/openapi.json
  )
}

prep_dir() {
  dir="$1"
  rm -rf "$dir"
  mkdir -p "$dir"
}

generate_client() {
  openapi-generator-cli generate \
    -i openapi.json \
    --skip-validate-spec \
    --git-user-id nextcloud \
    --git-repo-id client-sdks \
    --git-host github.com \
    "$@"
}

clean_dir() {
    dir="$1"
    rm -f "$dir/git_push.sh" "$dir/.travis.yml"
}

generate_python() {
  dir="clients/python"

  prep_dir "$dir"
  generate_client -o "$dir" -g python -c configs/python.yml
  rm -r "$dir/.github" "$dir/.gitlab-ci.yml" "$dir/test" "$dir/test-requirements.txt"
  clean_dir "$dir"
}

generate_typescript() {
  dir="clients/typescript"

  prep_dir "$dir"
  generate_client -o "$dir" -g typescript-axios -c configs/typescript.yml
  clean_dir "$dir"

  (
    cd "$dir"
    npm install
  )
}

generate_go() {
  dir="clients/go"

  prep_dir "$dir"
  generate_client -o "$dir" -g go -c configs/go.yml
  rm -r "$dir/test"
  clean_dir "$dir"
}

generate_rust() {
  dir="clients/rust"

  prep_dir "$dir"
  generate_client -o "$dir" -g rust -c configs/rust.yml
  clean_dir "$dir"
}

generate_spec

generate_python
generate_typescript
generate_go
generate_rust
