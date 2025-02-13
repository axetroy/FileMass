/**
 * Usage:
 *
 * GIT_REF=refs/tags/v1.0.0 node npm/prepare.js
 */
const fs = require("fs");
const path = require("path");

const ref = process.env.GIT_REF; // refs/tags/v1.0.0

const arr = ref.split("/");
const version = arr[arr.length - 1].replace(/^v/, "");

console.log(`prepare publish to npm for: ${version}`);

const packages = fs
  .readdirSync(__dirname)
  .filter((v) => v.startsWith("filemass-"))
  .concat(["filemass"]);

for (const pkgName of packages) {
  const pkgPath = path.join(__dirname, pkgName, "package.json");

  const pkg = require(pkgPath);

  pkg.version = version;

  if (pkg.optionalDependencies) {
    for (const subDeps in pkg.optionalDependencies) {
      if (subDeps.startsWith("@axetroy/filemass-")) {
        pkg.optionalDependencies[subDeps] = version;
      }
    }
  }

  fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2));

  if (pkgName.startsWith("filemass-")) {
    const fileMap = {
      "filemass-darwin-arm64": "filemass_darwin_arm64",
      "filemass-darwin-amd64": "filemass_darwin_amd64_v1",
      "filemass-linux-amd64": "filemass_linux_amd64_v1",
      "filemass-linux-arm64": "filemass_linux_arm64",
      "filemass-freebsd-arm64": "filemass_freebsd_arm64",
      "filemass-freebsd-amd64": "filemass_freebsd_amd64_v1",
      "filemass-openbsd-arm64": "filemass_openbsd_arm64",
      "filemass-openbsd-amd64": "filemass_openbsd_amd64_v1",
      "filemass-windows-amd64": "filemass_windows_amd64_v1",
      "filemass-windows-arm64": "filemass_windows_arm64",
    };

    if (pkgName in fileMap === false)
      throw new Error(`Can not found prebuild file for package '${pkgName}'`);

    const distFolder = fileMap[pkgName];

    const executableFileName =
      "filemass" + (pkgName.indexOf("windows") > -1 ? ".exe" : "");

    const executableFilePath = path.join(
      __dirname,
      "..",
      "dist",
      distFolder,
      executableFileName
    );

    fs.copyFileSync(
      executableFilePath,
      path.join(__dirname, pkgName, executableFileName)
    );
  } else {
    fs.copyFileSync(
      path.join(__dirname, "..", "README.md"),
      path.join(__dirname, pkgName, "README.md")
    );

    fs.copyFileSync(
      path.join(__dirname, "..", "LICENSE"),
      path.join(__dirname, pkgName, "LICENSE")
    );
  }
}
