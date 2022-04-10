# Repository of https://new-life-climbing.org

## Deploying
dist branch doesn't ignore changes in the dist/ folder  
simply:

``` bash
  git c dist
  go run .
  git add .
  git commit -m "whatever"
  git push origin dist
```

netlify ci should take care of the rest (it automatically deploys the **dist/** folder from branch **dist**)

We could automate this by a git hook or simply using the netlify build settings, but since the project is a simple templated static site i don't think its worth the effort. Maybe i'll get to it in the future.


