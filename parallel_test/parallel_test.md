Example use cases for parallel tests:
1. Simulating a real-world scenario
2. Verify that a type is truly thread-safe

(1) A web app with many users
(2) Verify that your in-memory cache can handle multiple concurrent web requests using it

Parallelism could also mean more work:
- Tests can't use as many hard-coded values; eg.: unique email constraints
- Tests might try to use shared resources incorrectly eg.:
  image manipulation on the same image or sharing a DB that doesn't suppor multiple
  concurrent connections