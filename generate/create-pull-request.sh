#!/bin/bash

# This script runs on a regular cron within Travis-CI.
# 
# It checks to see if any AWS CloudFormation resources have been updated, 
# and if so, creates a new pull request to apply them back to the awslabs repo.

# Bomb out on errors
set -e

# The repo/branch to create the PR against
REPO="awslabs/goformation"
SRC_BRANCH="master"
DST_BRANCH="master"

# Git details (for the commit)
COMMIT_NAME="AWS GoFormation"
COMMIT_EMAIL="goformation@amazon.com"
COMMIT_MSG="AWS CloudFormation Update ($(date +%F))"

# Pull requests and commits to other branches shouldn't try to deploy, just build to verify
if [ "$TRAVIS_EVENT_TYPE" != "cron" ]; then
    echo "Skipping auto generation of AWS CloudFormation resources; just doing a build."
    exit 0
fi

# Pull requests and commits to other branches shouldn't try to deploy, just build to verify
if [ "$TRAVIS_PULL_REQUEST" != "false" -o "$TRAVIS_BRANCH" != "$SRC_BRANCH" ]; then
    echo "Skipping deploy; just doing a build."
    exit 0
fi

# Configure the Git user for any commits
git config --global user.name "${COMMIT_NAME}"
git config --global user.email "${COMMIT_EMAIL}"

echo "Checking out github.com/${REPO}..."
UPSTREAM_DIR=/tmp/upstream/src/github.com/awslabs/goformation
mkdir -p ${UPSTREAM_DIR}
git clone https://github.com/${REPO}.git ${UPSTREAM_DIR} > /dev/null 2>&1
GOPATH=/tmp/upstream
cd ${UPSTREAM_DIR}

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
REQUEST_BRANCH="aws-goformation-updates"
echo "Creating new branch: ${REQUEST_BRANCH} tracking origin/${SRC_BRANCH}"
git checkout -b ${REQUEST_BRANCH} origin/${SRC_BRANCH}
git pull --rebase origin ${REQUEST_BRANCH} || true
 
echo "Auto-generating AWS CloudFormation resources..."
go generate

# Don't proceed if there were no new resources generated/updated
if [[ -z $(git status -s | grep "cloudformation/") ]]; then
    echo "No changes - no pull request necessary."
    exit 0;
fi
  
echo "Changes found:"
git status -s

echo "Running tests..."
go test -v ./...

echo "Committing changes..."
git add cloudformation/*
git commit -m "${COMMIT_MSG}" 

echo "Pushing changes..."
git remote add origin-push https://${GITHUB_TOKEN}@github.com/${REPO}.git > /dev/null 2>&1
git push --quiet --set-upstream origin-push ${REQUEST_BRANCH}

echo "Installing GitHub Hub"
git clone https://github.com/github/hub.git /tmp/hub
cd /tmp/hub
./script/build 

echo "Generating Pull Request for merging ${REPO}/${REQUEST_BRANCH} to ${REPO}/${DST_BRANCH}..."
cd ${UPSTREAM_DIR}
echo "${COMMIT_MSG}" | /tmp/hub/bin/hub pull-request -F - /dev/null 2>&1 || true