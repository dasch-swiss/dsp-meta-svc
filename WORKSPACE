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
http_archive(
    name = "build_bazel_rules_nodejs",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/3.0.0/rules_nodejs-3.0.0.tar.gz"],
    sha256 = "6142e9586162b179fdd570a55e50d1332e7d9c030efd853453438d607569721d",
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories")

# NOTE: this rule installs nodejs, npm, and yarn, but does NOT install
# your npm dependencies into your node_modules folder.
# You must still run the package manager to do this.
node_repositories(package_json = ["//:package.json"])

load("@build_bazel_rules_nodejs//:package.bzl", "rules_nodejs_dev_dependencies")
rules_nodejs_dev_dependencies()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
  name = "npm",
  package_json = "//:package.json",
  yarn_lock = "//:yarn.lock",
)

#
# Download the rules_svelte repository
#
http_archive(
    name = "build_bazel_rules_svelte",
    url = "https://github.com/thelgevold/rules_svelte/archive/0.15.zip",
    strip_prefix = "rules_svelte-0.15",
    sha256 = "1b04eb08ef80636929d152bb2f2733e36d9e0b8ad10aca7b435c82bd638336f5"
) 

load("@build_bazel_rules_svelte//:defs.bzl", "rules_svelte_dependencies")
rules_svelte_dependencies()


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
