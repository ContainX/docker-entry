# docker-entry

[![Build Status](https://travis-ci.org/ContainX/docker-entry.svg)](https://travis-ci.org/ContainX/docker-entry)

Secrets, Configuration and Docker Exec - Ridiculously Simple!

This is currently in active development and will be released and documented soon!

**Pending Tasks**

- [x] Init System - forward signals to child process
- [ ] Local and Remote configuration processing into context
- [ ] Secret loading, Applying to context
- [ ] Template processing
- [ ] Customization flags / environment variables
- [ ] Tests

## Quick Start

To get started `docker-entry` should be the main entry-point to your Dockerfile. The CMD or at runtime
command will be executed as the child process.

```
FROM alpine

ADD https://github.com/ContainX/docker-entry/releases/download/0.1/docker-entry /docker-entry
RUN chmod +x /docker-entry

ENTRYPOINT ["/docker-entry"]
```