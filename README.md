# Artillery2k6

This is a CLI tool designed to translate Artillery load test scripts into equivalent k6 scripts, with the goal of simplifying the process of migrating from Artillery to k6.

This tool works to make a 'best effort' approach to migrating to k6, and may not support all features of Artillery. It is recommended to review the generated k6 script to ensure that it meets your requirements before use.

## Features
- Parses Artillery YAML test scripts
- Converts Artillery test scripts to k6 test scripts
- Supports the following features
  - Single scenarios
  - ArrivalRate Phases
  - Variable declarations
  - Multi-step flows
  - HTTP methods: GET, POST & PUT Requests (with support for headers and JSON payloads)
  - Response status code expectations
  - Response JSON captures
  - 'Think' actions
  - 'Log' actions
  - 'Function' actions
  - 'BeforeRequest' and 'AfterResponse' hooks
  - Importing of processors, including nested imports
    - Conversion of context variables to `globalThis`
    - Cleanup of any `export` statements

## Roadmap
- Support for additional Artillery features:
  - Multiple weighted scenarios
  - Loop actions
  - Additional phase types and k6 executors
  - Environment support
  - `BeforeScenario` and `AfterScenario` processor hooks
  - `before` and `after` sections
  - CLI Flags to enable/disable functionality

## Installation
`go install github.com/cjsaurusrex/artillery2k6@latest`

## Usage
`artillery2k6 convert <path-to-artillery-script> [-o output-file]`

## Example
Input file:
```yaml
config:
  phases:
    - duration: 60
      arrivalRate: 5
scenarios:
  - flow:
      - get:
          name: "Get Donkey"
          url: "https://en.wikipedia.org/wiki/Donkey"
```

Output:
```javascript
import http from "k6/http"

export const options = {
  stages: [
    {"duration":"60","target":"5"},
  ]
}

export default function() {
  let getDonkey = http.get("https://en.wikipedia.org/wiki/Donkey", {});
}
```