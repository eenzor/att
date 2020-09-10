# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.3] - 2020-09-11

### Added
- The request logs now mimic apache's Combined Log Format

### Changed
- The request log output is now configured using the `-log` command line flag

## [0.0.2] - 2020-09-10

### Changed
- Refactor version handler function to be more testable
- Read the description from the DESCRIPTION environment variable (or use default)

## [0.0.1] - 2020-09-09

### Added
- Basic functionality for /version endpoint
