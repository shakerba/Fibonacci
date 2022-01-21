# fibonacci

## Run locally

1. Go to the directory `fibonacci` and run `make build` (Note: Binary created is only executable on Mac, change line 44 in Makefile from `GOOS=darwin` to `GOOS=linux` to make it executable on linux)
2. Run the server `fibonacci/build/fibonacci --httpAddr :<port>`
3. `curl https://localhost:<port>/current` returns `200` response with the current number
   `curl https://localhost:<port>/next` returns `200` response with the next number
   `curl https://localhost:<port>/previous` returns `200` response with the previous number
4. Use `ps` to get the process id of the server, and perform `kill -9 <process id>` to end the server (Other codes will cause the server to restart)
