const core = require("@actions/core");
const github = require("@actions/github");

async function run() {
  const title = core.getInput("title");
  const fact = core.getInput("fact");

  const token = core.getInput("token");

  try {
    const octokit = new github.GitHub(token);

    const newIssue = await octokit.issues.create({
      repo: github.context.repo.repo,
      owner: github.context.repo.owner,
      title,
      body: fact
    });
  } catch (error) {
    core.setFailed(error.message);
  }
}

run();
