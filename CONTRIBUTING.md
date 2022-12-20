# Contributing

If you have an idea or feature request please open an [issue][1], even if you
don't have time to contribute!

## Making Changes

> **Note**: This guide assumes you have a working Go 1.17 or later installation.

To get started, [fork][2] this repository on GitHub and clone a working copy for
development:

    $ git clone git@github.com:YOUR-USERNAME/go-hid.git

Once you are finished, be sure to test changes locally by issuing:

    $ go test ./...

Finally, commit your changes and create a [pull request][3] against the `master`
branch for review.

## Making New Releases

Making new releases is automated by GitHub Actions. Releases should only be
created from the `master` branch.

To make a new release, follow these steps:

1. Verify the latest results of the [CI][4] workflow on the `master` branch.

2. Create a release tag by issuing:

       $ git tag -a -m 'Release v<version>' v<version>

3. Push the release tag to the remote repository and verify the results of the
   [Release][5] workflow:

       $ git push origin --tags

## License

By contributing to this repository, you agree that your contributions will be
licensed under its Simplified BSD License.

[1]: https://github.com/sstallion/go-hid/issues
[2]: https://docs.github.com/en/github/getting-started-with-github/fork-a-repo
[3]: https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request
[4]: https://github.com/sstallion/go-hid/actions/workflows/ci.yml
[5]: https://github.com/sstallion/go-hid/actions/workflows/release.yml
[6]: https://docs.freebsd.org/en/books/porters-handbook/upgrading/
