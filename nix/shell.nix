{
  mkShell,
  go,
  gopls,
  gotools,
  go-tools,
  hyperfine,
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
    hyperfine
    just
  ];
}
