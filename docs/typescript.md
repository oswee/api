# TypeScript

Run `protoc-gen.sh` from the root of this project.

Make sure to create proper `package.json` file. Specify correct package name and `publishConfig` properties.

Remember, that GitHub does not allow to delete packages and all package names become "registered" even for deleted private packages.
NPM allows to delete package within 72h.

Make sure to create Github Personal Access Token.
Add PAT to the environment variables and make sure `${GITHUB_NPM_PAT}` is in `../.npmrc` and it is accessible.

Run `git add .` to stage all changes

Run `git commit -m feat(modulex): added new field` to commit the changes.

At this point all changes should be committed. Lerna will fail otherwise.

Run `yarn pub`. It will execute `lerna publish` and will publish all change and packages to the registry.
