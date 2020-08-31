const tasks = arr => arr.join(' && ')

module.exports = {
  hooks: {
    'pre-commit': tasks(['lint-staged']),
    // "pre-commit": "lerna run --concurrency 1 --stream precommit",
    'commit-msg': tasks(['commitlint -E HUSKY_GIT_PARAMS']),
  },
}
