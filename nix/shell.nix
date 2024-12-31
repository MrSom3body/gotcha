{
  mkShell,
  go,
  gopls,
  gotools,
  go-tools,
  just,
}:
mkShell {
  buildInputs = [
    go
    gopls
    gotools
    go-tools
  ];

  packages = [
    just
  ];
}
