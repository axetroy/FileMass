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

const root = path.join(__dirname, "..");

console.log(`prepare publish to npm for: ${version} in ${root}`);

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
      "filemass-darwin-arm64": "FileMass_darwin_arm64_v8.0",
      "filemass-darwin-amd64": "FileMass_darwin_amd64_v1",
      "filemass-linux-amd64": "FileMass_linux_amd64_v1",
      "filemass-linux-arm64": "FileMass_linux_arm64_v8.0",
      "filemass-freebsd-arm64": "FileMass_freebsd_arm64_v8.0",
      "filemass-freebsd-amd64": "FileMass_freebsd_amd64_v1",
      "filemass-openbsd-arm64": "FileMass_openbsd_arm64_v8.0",
      "filemass-openbsd-amd64": "FileMass_openbsd_amd64_v1",
      "filemass-windows-amd64": "FileMass_windows_amd64_v1",
      "filemass-windows-arm64": "FileMass_windows_arm64_v8.0",
    };

    if (pkgName in fileMap === false)
      throw new Error(`Can not found prebuild file for package '${pkgName}'`);

    const distFolder = fileMap[pkgName];

    const executableFileName =
      "filemass" + (pkgName.indexOf("windows") > -1 ? ".exe" : "");

    const source = path.join(
      __dirname,
      "..",
      "dist",
      distFolder,
      executableFileName
    );

    const dest = path.join(__dirname, pkgName, executableFileName);

    if (!fs.existsSync(source)) {
      throw new Error(`Can not found prebuild file: ${source}`);
    }

    fs.copyFileSync(source, dest);
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
