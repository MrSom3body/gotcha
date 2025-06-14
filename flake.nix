{
  description = "gotcha: a linux fetch written in go";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    systems = ["x86_64-linux" "aarch64-linux"];
    forEachSystem = nixpkgs.lib.genAttrs systems;
    pkgsForEach = nixpkgs.legacyPackages;
  in {
    packages = forEachSystem (system: {
      default = self.packages.${system}.gotcha;
      gotcha = pkgsForEach.${system}.callPackage ./nix/package.nix {};
      gotcha-update = pkgsForEach.${system}.callPackage ./nix/package.nix {
        enableUpdateCmd = true;
      };
    });

    devShells = forEachSystem (system: {
      default = pkgsForEach.${system}.callPackage ./nix/shell.nix {};
    });
  };
}
