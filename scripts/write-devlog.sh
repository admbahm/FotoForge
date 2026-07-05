#!/usr/bin/env bash
set -euo pipefail

usage() {
  printf 'Usage: %s short-slug\n' "${0##*/}" >&2
}

if [[ $# -ne 1 ]]; then
  usage
  exit 2
fi

slug=$1
if [[ ! $slug =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]]; then
  printf 'error: slug must contain lowercase letters, numbers, and single hyphens only\n' >&2
  exit 2
fi

script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
repo_root=$(cd -- "$script_dir/.." && pwd)
template="$repo_root/docs/devlog/TEMPLATE.md"
today=$(date +%F)
relative_path="docs/devlog/${today}-${slug}.md"
target="$repo_root/$relative_path"
temporary="${target}.tmp.$$"

if [[ ! -f $template ]]; then
  printf 'error: template not found: %s\n' "$template" >&2
  exit 1
fi

if [[ -e $target ]]; then
  printf 'error: devlog already exists: %s\n' "$relative_path" >&2
  exit 1
fi

cleanup() {
  rm -f -- "$temporary"
}
trap cleanup EXIT

sed "s/{{DATE}}/${today}/g" "$template" >"$temporary"
mv -- "$temporary" "$target"
trap - EXIT

printf '%s\n' "$relative_path"
