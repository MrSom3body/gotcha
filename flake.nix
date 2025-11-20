{
  description = "gotcha: a linux fetch written in go";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default-linux";
    git-hooks-nix = {
      url = "github:cachix/git-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      ...
    }@inputs:
    let
      inherit (nixpkgs) lib;

      forEachSystem = f: lib.genAttrs (import inputs.systems) (system: f pkgsFor.${system});
      pkgsFor = lib.genAttrs (import inputs.systems) (
        system:
        import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        }
      );
    in
    {
      formatter = forEachSystem (pkgs: pkgs.nixfmt-tree);

      packages = forEachSystem (pkgs: {
        gotcha = pkgs.callPackage ./nix/package.nix { };
        gotcha-update = pkgs.callPackage ./nix/package.nix { enableUpdateCmd = true; };
        default = self.packages.${pkgs.stdenv.hostPlatform.system}.gotcha;
      });

      devShells = forEachSystem (
        pkgs:
        import ./nix/shell.nix {
          inherit self;
          inherit pkgs;
        }
      );

      checks = forEachSystem (
        pkgs:
        import ./nix/checks.nix {
          inherit inputs;
          inherit pkgs;
        }
      );
    };
}
