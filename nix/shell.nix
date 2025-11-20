{
  self,
  pkgs,
}:
{
  default = pkgs.mkShell {
    packages = with pkgs; [
      go
      gopls
      gotools
      go-tools
    ];

    buildInputs = with pkgs; [
      hyperfine
      just
    ];

    shellHook = ''
      ${self.checks.${pkgs.stdenv.hostPlatform.system}.pre-commit-check.shellHook}
    '';
  };
}
