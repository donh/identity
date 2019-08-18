const KeyManager = artifacts.require("KeyManager");

module.exports = function(deployer) {
  deployer.deploy(KeyManager);
};
