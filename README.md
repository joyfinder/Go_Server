<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Go_Server_Api](#go_server_api)
  - [Project description](#project-description)
  - [Status](#status)
  - [API debugging issue](#api-debugging-issue)
  - [See also](#see-also)
  - [Getting started](#getting-started)
  - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Go_Server_Api

## Project description

This project is designed for Golang side project of building server API from scratch.

## Status

The project is currently under construction. 

## API debugging issue

        Error: write EPROTO 5636819960:error:100000f7:SSL routines:OPENSSL_internal:WRONG_VERSION_NUMBER:../../../src/third_party/boringssl/src/ssl/tls_record.cc:242:
    
    Solution: Root cause could be the use of https instead of http. Changing url into http, and re-run it.

## See also

## Getting started

## Notes
