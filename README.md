# Search all repositories in an organizations on github

This is simple program that fetches and searches all Optimizely's Github repositories for a specific string.

## Install

``` 
$ go get github.com/optimizely/fetch-all-github-repos
```

## Usage

```
$fetch-all-github-repos 0X123ThisIsAGithubAAccessKey
```

Or

```
$ time ./fetch-all-github-repos 0X123ThisIsAGithubAAccessKey | tee output.txt

repo number: 0 - git@github.com:optimizely/optimizely.git
git clone --depth 1 git@github.com:optimizely/optimizely.git tmpdir
ag eslint-scope tmpdir
tmpdir/src/www/frontend/yarn.lock:860:    eslint-scope "~3.7.1"
tmpdir/src/www/frontend/yarn.lock:3870:eslint-scope@^3.7.1, eslint-scope@~3.7.1:
tmpdir/src/www/frontend/yarn.lock:3872:  resolved "https://optimizely.jfrog.io/optimizely/api/npm/npm/eslint-scope/-/eslint-scope-3.7.1.tgz?dl=https://registry.yarnpkg.com/eslint-scope/-/eslint-scope-3.7.1.tgz#3d63c3edfda02e06e01a452ad88caacc7cdcb6e8"
tmpdir/src/www/frontend/yarn.lock:3892:    eslint-scope "^3.7.1"
tmpdir/src/www/frontend/yarn.lock:11571:    eslint-scope "^3.7.1"
ag eslint-config-eslint tmpdir
rm -rf tmpdir
repo number: 1 - git@github.com:optimizely/blog.optimizely.git
...
rm -rf tmpdir
repo number: 821 - git@github.com:optimizely/optimizely-ktlint.git
git clone --depth 1 git@github.com:optimizely/optimizely-ktlint.git tmpdir
ag eslint-scope tmpdir
ag eslint-config-eslint tmpdir
rm -rf tmpdir
./fetch-all-github-repos  135.25s user 239.15s system 31% cpu 19:55.51 total
tee output.txt  0.01s user 0.16s system 0% cpu 19:55.51 total
```

## Credits

- Ola

## License

The MIT License (MIT).
