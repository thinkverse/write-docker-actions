# Welcome to write-docker-actions 🎉

This is the repo **I** used for the [GitHub Actions: Write Docker container actions] course on GitHub's [Learning Lab].<br>
Where **I** learned to create Docker based GitHub Actions, all actions can be found under [.github/actions] folder.

## How the Actions work

###### The hello-world Action

The `hello-world` action takes two required users to greet and a third optional one, it uses `golang` to write out `Hello {user}`

###### The cat-facts Action

The `cat-facts` action uses `Python 3` to fetch cat facts from the [catfact] API, then returns and outputs a random fact for another action to use.

###### The issue-maker Action

The `issue-maker` action uses `node` and the [actions/toolkit] to accept a `title`, `fact` and a `token` to create a new issue with that title and fact.

[gitHub actions: write docker container actions]: https://lab.github.com/githubtraining/github-actions:-write-docker-container-actions
[learning lab]: https://lab.github.com
[.github/actions]: https://github.com/thinkverse/write-docker-actions/tree/HEAD/.github/actions
[catfact]: https://catfact.ninja/
[actions/toolkit]: https://github.com/actions/toolkit
