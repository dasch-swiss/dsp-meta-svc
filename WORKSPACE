workspace(
    name = "svelte_bazel_example",
    managed_directories = {"@npm": ["node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

#
# bazel-skylib 1.0.2 released 2019.10.09 (https://github.com/bazelbuild/bazel-skylib/releases/tag/1.0.2)
#
skylib_version = "1.0.2"
http_archive(
    name = "bazel_skylib",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/{}/bazel-skylib-{}.tar.gz".format (skylib_version, skylib_version),
    sha256 = "97e70364e9249702246c0e9444bccdc4b847bed1eb03c5a3ece4f83dfe6abc44",
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
bazel_skylib_workspace()


#
# Download the rules_nodejs repository
#
rules_nodejs_version="3.0.0"
rules_nodejs_version_sha256="6142e9586162b179fdd570a55e50d1332e7d9c030efd853453438d607569721d"
http_archive(
    name = "build_bazel_rules_nodejs",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/%s/rules_nodejs-%s.tar.gz" % (rules_nodejs_version, rules_nodejs_version)],
    sha256 = rules_nodejs_version_sha256,
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories", "yarn_install")

# NOTE: this rule installs nodejs, npm, and yarn, but does NOT install
# your npm dependencies into your node_modules folder.
# You must still run the package manager to do this.
node_repositories(
    node_version = "14.15.4",
    node_repositories = {
        "14.15.4-darwin_amd64": ("node-v14.15.4-darwin-x64.tar.gz", "node-v14.15.4-darwin-x64", "6b0e19e5c2601ef97510f7eb4f52cc8ee261ba14cb05f31eb1a41a5043b0304e"),
        "14.15.4-linux_amd64": ("node-v14.15.4-linux-x64.tar.xz", "node-v14.15.4-linux-x64", "ed01043751f86bb534d8c70b16ab64c956af88fd35a9506b7e4a68f5b8243d8a"),
        "14.15.4-windows_amd64": ("node-v14.15.4-win-x64.zip", "node-v114.15.4-win-x64", "b2a0765240f8fbd3ba90a050b8c87069d81db36c9f3745aff7516e833e4d2ed6"),
    },
    node_urls = ["https://nodejs.org/dist/v{version}/{filename}"],
    package_json = ["//:package.json"]
)

load("@build_bazel_rules_nodejs//:package.bzl", "rules_nodejs_dev_dependencies")
rules_nodejs_dev_dependencies()

yarn_install(
  name = "npm",
  package_json = "//:package.json",
  yarn_lock = "//:yarn.lock",
)

##################################
# Support creating Docker images #
##################################

rules_docker_version="0.14.4"
rules_docker_version_sha256="4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036"
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = rules_docker_version_sha256,
    strip_prefix = "rules_docker-%s" % rules_docker_version,
    url = "https://github.com/bazelbuild/rules_docker/releases/download/v%s/rules_docker-v%s.tar.gz" % (rules_docker_version, rules_docker_version),
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")

pip_deps()

load("@io_bazel_rules_docker//nodejs:image.bzl", nodejs_image_repos = "repositories")

nodejs_image_repos()
