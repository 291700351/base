#!/usr/bin/env bash

tag_version=$1
git tag codes/"$tag_version"
git tag database/"$tag_version"
git tag tools/"$tag_version"

git push --tags