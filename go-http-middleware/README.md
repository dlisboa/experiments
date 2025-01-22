Server runs on :8080.

Make requests:

* `GET :8080/` for request logger middleware
* `POST :8080/limited @postbody` for request size middleware
* `GET :8080/panic` for panic/recover middleware
* `GET :8080/timeout` for timeout middleware
