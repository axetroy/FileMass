const spawn = require("child_process").spawn;
const os = require("os");
const path = require("path");

const platform = os.platform();
const arch = os.arch();

const ERR_NOT_SUPPORT = new Error("filemass does not support your platform");

const platformMap = {
  win32: {
    arm64: "filemass-windows-arm64",
    x64: "filemass-windows-amd64",
  },
  darwin: {
    arm64: "filemass-darwin-arm64",
    x64: "filemass-darwin-amd64",
  },
  linux: {
    arm64: "filemass-linux-arm64",
    x64: "filemass-linux-amd64",
  },
  freebsd: {
    arm64: "filemass-freebsd-arm64",
    x64: "filemass-freebsd-amd64",
  },
  openbsd: {
    arm64: "filemass-openbsd-arm64",
    x64: "filemass-openbsd-amd64",
  },
};

const archMap = platformMap[platform];

if (!archMap) throw ERR_NOT_SUPPORT;

const prebuildPackageName = archMap[arch];

if (!prebuildPackageName) throw ERR_NOT_SUPPORT;

const binaryPackageDir = path.dirname(
  require.resolve(`@axetroy/${prebuildPackageName}/package.json`)
);

const executableFileName = "filemass" + (platform === "win32" ? ".exe" : "");

const executableFilePath = path.join(binaryPackageDir, executableFileName);

/**
 *
 * @param {Array<string>} argv
 * @param {SpawnOptionsWithoutStdio} [spawnOptions]
 * @returns
 */
function exec(argv, spawnOptions = {}) {
  const ps = spawn(executableFilePath, argv, {
    ...spawnOptions,
    stdout: "piped",
  });

  return ps;
}

/**
 * @param {Object} params0
 * @param {string} params0.config The config file path
 * @param {number} [params0.maxError] The max error
 * @returns {Promise<any>}
 */
function filemass({ config, maxError }) {
  const args = ["--json", "--no-color", "--config", config];

  if (maxError) {
    args.push("--max-error");
    args.push(maxError);
  }

  const ps = exec(args, {
    stdout: "pipe",
    stderr: "pipe",
  });

  let stdout = Buffer.from("");
  let stderr = Buffer.from("");

  ps.stdout.on("data", (/** @type {Buffer} */ buf) => {
    stdout = Buffer.concat(stdout, buf);
  });

  ps.stderr.on("data", (/** @type {Buffer} */ buf) => {
    stderr = Buffer.concat(stderr, buf);
  });

  return new Promise((resolve, reject) => {
    ps.on("exit", (code) => {
      if (code === 0) {
        const output = stdout.toString("utf-8").trim();

        try {
          resolve(JSON.parse(output));
        } catch (err) {
          reject(err);
        }
      } else {
        reject(new Error(`filemass error: \n${stderr.toString("utf-8").trim()}`));
      }
    });
  });
}

module.exports = filemass;
module.exports.filemass = filemass;
module.exports.exec = exec;
